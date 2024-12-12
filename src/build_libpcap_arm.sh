git clone https://github.com/android/platform_external_libpcap
mv platform_external_libpcap libpcap_arm32
export CC="${NDK}/toolchains/llvm/prebuilt/linux-x86_64/bin/armv7a-linux-androideabi21-clang"
cd libpcap_arm32
./configure --host=armv7a-linux-androideabi --with-pcap=linux
make
