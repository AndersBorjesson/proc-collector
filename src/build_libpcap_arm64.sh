git clone https://github.com/android/platform_external_libpcap
mv platform_external_libpcap libpcap_arm64
export CC="${NDK}/toolchains/llvm/prebuilt/linux-x86_64/bin/aarch64-linux-android21-clang"
cd libpcap_arm64
./configure --host=aarch64-linux-android --with-pcap=linux
make
