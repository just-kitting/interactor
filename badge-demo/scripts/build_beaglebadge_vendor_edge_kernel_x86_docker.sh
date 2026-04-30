#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd -- "${SCRIPT_DIR}/.." && pwd)"
ARMBIAN_DIR="${REPO_ROOT}/components/armbian-build"
KERNEL_PATCH="${ARMBIAN_DIR}/patch/kernel/archive/k3-6.12/0001-i2c-omap-add-slave-registration-support.patch"

if [[ "$(uname -s)" != "Linux" ]]; then
	echo "This script expects a Linux x86_64 host with Docker installed." >&2
	exit 1
fi

if [[ "$(uname -m)" != "x86_64" ]]; then
	echo "This script is intended for an x86_64 host. Current host: $(uname -m)" >&2
	exit 1
fi

if [[ ! -x "${ARMBIAN_DIR}/compile.sh" ]]; then
	echo "Armbian build framework not found at ${ARMBIAN_DIR}" >&2
	exit 1
fi

if [[ ! -f "${KERNEL_PATCH}" ]]; then
	echo "Expected BeagleBadge slave-mode kernel patch not found:" >&2
	echo "  ${KERNEL_PATCH}" >&2
	exit 1
fi

if ! command -v docker >/dev/null 2>&1; then
	echo "docker is required on the host." >&2
	exit 1
fi

if ! docker info >/dev/null 2>&1; then
	echo "docker is installed but not usable by the current user." >&2
	echo "Fix Docker access first, then rerun this script." >&2
	exit 1
fi

cd "${ARMBIAN_DIR}"

echo "Building BeagleBadge vendor-edge kernel packages through Armbian's Docker workflow..."
echo "Output packages will be written under:"
echo "  ${ARMBIAN_DIR}/output/debs/"
echo "Kernel patch expected in this build:"
echo "  ${KERNEL_PATCH}"

exec ./compile.sh \
	kernel \
	BOARD=beaglebadge \
	BRANCH=vendor-edge \
	RELEASE=trixie \
	BUILD_MINIMAL=yes \
	BUILD_DESKTOP=no \
	KERNEL_BTF=no \
	"$@"
