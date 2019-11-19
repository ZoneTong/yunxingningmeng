# SHELL := /bin/csh
mac:
	./buildx.sh 
win:
	./buildx.sh -n -s win
pack:
	./buildx.sh -n -s pack

all: pack

aar:
	xgo --targets=android-16/*  .   
	unzip -d android/aar yunxingningmeng-android-16.aar
	# nm -D android/jni/x86/libyunxingningmeng.so | grep -v Java

so:
	# xgo --targets=android-16/arm -buildmode=c-shared  .
	xgo --targets=linux/arm64,linux/arm,linux/amd64,linux/386 -buildmode=c-shared  .
	cp yunxingningmeng-linux-arm-5.h 	android/jni/armeabi-v7a/yunxingningmeng.h
	cp yunxingningmeng-linux-arm-5.so 	android/jni/armeabi-v7a/libyunxingningmeng.so
	cp yunxingningmeng-linux-arm64.h 	android/jni/arm64-v8a/yunxingningmeng.h
	cp yunxingningmeng-linux-arm64.so 	android/jni/arm64-v8a/libyunxingningmeng.so
	cp yunxingningmeng-linux-386.h 		android/jni/x86/yunxingningmeng.h
	cp yunxingningmeng-linux-386.so 	android/jni/x86/libyunxingningmeng.so
	cp yunxingningmeng-linux-amd64.h 	android/jni/x86_64/yunxingningmeng.h
	cp yunxingningmeng-linux-amd64.so 	android/jni/x86_64/libyunxingningmeng.so

liba:
	xgo --targets=linux/386,linux/arm -buildmode=c-archive  .
	cp yunxingningmeng-linux-arm-5.h 	android/jni/armeabi-v7a/yunxingningmeng.h
	cp yunxingningmeng-linux-arm-5.a 	android/jni/armeabi-v7a/libyunxingningmeng.a
	cp yunxingningmeng-linux-386.h 		android/jni/x86/yunxingningmeng.h
	cp yunxingningmeng-linux-386.a 		android/jni/x86/libyunxingningmeng.a

android: # CGO_ENABLED=1 GOOS=android PATH=$(PATH):$(ANDROID_NDK_ROOT)/toolchains/llvm/prebuilt/darwin-x86_64/bin
	./android.sh
	# GOARCH=arm CC=armv7a-linux-androideabi19-clang CXX=armv7a-linux-androideabi19-clang++  go build -buildmode=c-shared -o android/jni/armeabi-v7a/libyunxingningmeng.a
	# cp -rf android/jni/* /opt/gopath/src/github.com/ZoneTong/MyApplication/app/libs/


clean:
	rm -f *.o *.a *.so yunxingningmeng*