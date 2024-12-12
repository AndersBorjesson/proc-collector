
export LIBPCAP=${PWD}/libpcap_arm64
export CC="${NDK}/toolchains/llvm/prebuilt/linux-x86_64/bin/aarch64-linux-android21-clang"
export CC_FOR_TARGET="${NDK}/toolchains/llvm/prebuilt/linux-x86_64/bin/aarch64-linux-android21-clang"
export CGO_ENABLED=1
export GOOS=android
export GOARCH=arm64
export PATH=${NDK}/toolchains/llvm/prebuilt/linux-x86_64/bin:${PATH}
export SYSROOT_PATH=${NDK}/toolchains/llvm/prebuilt/linux-x86_64/sysroot
export CGO_CFLAGS="--sysroot=${SYSROOT_PATH}  -I${NDK}/toolchains/llvm/prebuilt/linux-x86_64/sysroot/usr/include/ -I${LIBPCAP}"
export CGO_LDFLAGS="--sysroot=${SYSROOT_PATH}  -L${NDK}/toolchains/llvm/prebuilt/linux-x86_64/sysroot/usr/lib/ -L${LIBPCAP} -lc -v"
go build  -buildvcs=false -o syscollector_android_arm64
