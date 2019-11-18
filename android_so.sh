
# function androidx(){
    mkdir -p android/jni
    export PATH=${PATH}:${ANDROID_NDK_ROOT}/toolchains/llvm/prebuilt/darwin-x86_64/bin
    export CGO_ENABLED=1
    export GOOS=android
    arch_dir=common
# }

function armeabi_v7a(){
    # source androidx
    export GOARCH=arm
    export CC=armv7a-linux-androideabi19-clang
    export CXX=armv7a-linux-androideabi19-clang++
    arch_dir=armeabi-v7a
    _libso
}

function arm64_v8a(){
    # source androidx
    export GOARCH=arm64
    export CC=aarch64-linux-android21-clang 
    export CXX=aarch64-linux-android21-clang++
    arch_dir=arm64-v8a
    _libso
}

function x86(){
    # source androidx
    export GOARCH=386
    export CC=i686-linux-android19-clang 
    export CXX=i686-linux-android19-clang++
    arch_dir=x86
    _libso
}


function x86_64(){
    # source androidx
    export GOARCH=amd64
    export CC=x86_64-linux-android21-clang 
    export CXX=x86_64-linux-android21-clang++
    arch_dir=x86_64
    _libso
}

function _libso(){
    echo $arch_dir
    go build -buildmode=c-shared -o android/jni/${arch_dir}/libyunxingningmeng.so
}

armeabi_v7a
arm64_v8a
x86
x86_64
# $@

