# proc-collector
Tools for collecting computer system statistics and transforming to analytical format

The intention of the project is to be able to store the content in the Linux /proc folder with a given frequency and store this as an easily parseable format.

This data can then be converted into a format, e.g., parquet, that is suitable for analytics.


## Build information

The collector requires libpcap built for the target architecture and operating system. 


A Makefile is available for building for Android 32 and 64 bit arm architectures. 

### Android Build

Android builds requires the Android NDK to be installed as this is used to build libpcap for target. 

Prior to build the path to the Android NDK needs to be assigned to variable NDK, eg. 
export NDK=~/Android/Sdk/ndk-bundle

The proc-collector can then be built in two steps

- First build libpcap for the relevant architecture
- Second build proc-collector for the same. 

That is 

```bash
make build-libpcap-arm
make build-android-arm
```

for 32 bit architecture or 

```bash
make build-libpcap-arm64
make build-android-arm64
```

for 64 bit architecture.