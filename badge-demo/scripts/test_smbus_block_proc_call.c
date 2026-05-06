#include <errno.h>
#include <fcntl.h>
#include <linux/i2c-dev.h>
#include <linux/i2c.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/ioctl.h>
#include <unistd.h>

static void usage(const char *argv0)
{
	fprintf(stderr,
		"Usage: %s <bus> <addr> <command> <count> <value>\n"
		"\n"
		"Example:\n"
		"  %s 3 0x30 0x03 1 4\n",
		argv0, argv0);
}

static long parse_num(const char *s, const char *label)
{
	char *end = NULL;
	long v;

	errno = 0;
	v = strtol(s, &end, 0);
	if (errno || !end || *end != '\0') {
		fprintf(stderr, "invalid %s: %s\n", label, s);
		exit(2);
	}

	return v;
}

int main(int argc, char **argv)
{
	char dev[32];
	struct i2c_smbus_ioctl_data ioctl_data;
	union i2c_smbus_data data;
	unsigned long funcs = 0;
	long bus, addr, command, count, value;
	int fd;
	int ret;
	int i;

	if (argc != 6) {
		usage(argv[0]);
		return 2;
	}

	bus = parse_num(argv[1], "bus");
	addr = parse_num(argv[2], "addr");
	command = parse_num(argv[3], "command");
	count = parse_num(argv[4], "count");
	value = parse_num(argv[5], "value");

	if (bus < 0 || bus > 255 || addr < 0 || addr > 0x7f ||
	    command < 0 || command > 0xff || count < 0 ||
	    count > I2C_SMBUS_BLOCK_MAX || value < 0 || value > 0xff) {
		fprintf(stderr, "argument out of range\n");
		return 2;
	}

	snprintf(dev, sizeof(dev), "/dev/i2c-%ld", bus);
	fd = open(dev, O_RDWR);
	if (fd < 0) {
		perror(dev);
		return 1;
	}

	if (ioctl(fd, I2C_FUNCS, &funcs) < 0) {
		perror("ioctl(I2C_FUNCS)");
		close(fd);
		return 1;
	}

	if (!(funcs & I2C_FUNC_SMBUS_BLOCK_PROC_CALL)) {
		fprintf(stderr,
			"adapter %s does not advertise I2C_FUNC_SMBUS_BLOCK_PROC_CALL\n",
			dev);
		close(fd);
		return 1;
	}

	if (ioctl(fd, I2C_SLAVE_FORCE, addr) < 0) {
		perror("ioctl(I2C_SLAVE_FORCE)");
		close(fd);
		return 1;
	}

	memset(&data, 0, sizeof(data));
	data.block[0] = count;
	data.block[1] = value;

	memset(&ioctl_data, 0, sizeof(ioctl_data));
	ioctl_data.read_write = I2C_SMBUS_WRITE;
	ioctl_data.command = command;
	ioctl_data.size = I2C_SMBUS_BLOCK_PROC_CALL;
	ioctl_data.data = &data;

	ret = ioctl(fd, I2C_SMBUS, &ioctl_data);
	if (ret < 0) {
		perror("ioctl(I2C_SMBUS BLOCK_PROC_CALL)");
		close(fd);
		return 1;
	}

	printf("count=%u\n", data.block[0]);
	for (i = 1; i <= data.block[0] && i <= I2C_SMBUS_BLOCK_MAX; i++)
		printf("%s0x%02x", (i == 1) ? "data=" : " ", data.block[i]);
	printf("\n");

	close(fd);
	return 0;
}
