from os.path import join

from SCons.Script import Builder, DefaultEnvironment, Default

env = DefaultEnvironment()

platform = env.PioPlatform()
board = env.BoardConfig()
toolchain_dir = platform.get_package_dir("toolchain-gccarmnoneeabi")

if not toolchain_dir:
    raise RuntimeError("PlatformIO ARM GCC toolchain package is missing")

toolchain_bin = join(toolchain_dir, "bin")
project_dir = env.subst("$PROJECT_DIR")
ldscript_path = join(project_dir, board.get("build.ldscript"))

common_flags = [
    "-mcpu=%s" % board.get("build.cpu"),
    "-mthumb",
    "-Os",
    "-g3",
    "-ffreestanding",
    "-fdata-sections",
    "-ffunction-sections",
]

env.Replace(
    AR=join(toolchain_bin, "arm-none-eabi-ar"),
    AS=join(toolchain_bin, "arm-none-eabi-gcc"),
    CC=join(toolchain_bin, "arm-none-eabi-gcc"),
    CXX=join(toolchain_bin, "arm-none-eabi-g++"),
    GDB=join(toolchain_bin, "arm-none-eabi-gdb"),
    OBJCOPY=join(toolchain_bin, "arm-none-eabi-objcopy"),
    RANLIB=join(toolchain_bin, "arm-none-eabi-ranlib"),
    SIZETOOL=join(toolchain_bin, "arm-none-eabi-size"),
    PROGNAME="firmware",
    LDSCRIPT_PATH=ldscript_path,
)

env.Append(
    ASFLAGS=common_flags,
    CCFLAGS=common_flags + ["-std=c11"],
    LINKFLAGS=common_flags
    + [
        "-nostartfiles",
        "-nostdlib",
        "-Wl,--gc-sections",
        "-Wl,-Map,%s" % join("$BUILD_DIR", "${PROGNAME}.map"),
    ],
    CPPPATH=[join(project_dir, "include")],
    LIBS=[],
)

env.Append(
    BUILDERS={
        "ElfToBin": Builder(
            action="$OBJCOPY -O binary $SOURCE $TARGET",
            suffix=".bin",
        )
    }
)

program = env.BuildProgram()
binary = env.ElfToBin(join("$BUILD_DIR", "${PROGNAME}"), program)

upload_cmd = "%s %s" % (
    join(project_dir, "scripts", "flash_zepto_bsl.sh"),
    join(env.subst("$BUILD_DIR"), "${PROGNAME}.bin"),
)

env.AddPostAction(
    program,
    env.VerboseAction(
        "$SIZETOOL -A $TARGET",
        "Calculating size $TARGET",
    ),
)

env.AlwaysBuild(
    env.Alias(
        "upload",
        binary,
        env.VerboseAction(upload_cmd, "Uploading ${PROGNAME}.bin"),
    )
)

Default(program, binary)
