#!/bin/sh

APP="macYunxing.app"
exe="yunxing"
mkdir -p $APP/Contents/{MacOS,Resources,Frameworks}
# go build -o $exe -ldflags '-linkmode "external" -extldflags "-static"'  .
go build -ldflags "-s -w -X main.trialday=3650 -X 'main.builddate=`date +%Y%m%d`'  "
mv yunxingningmeng $APP/Contents/MacOS/$exe

# install_name_tool -add_rpath @loader_path/../Frameworks $APP/Contents/MacOS/$exe
# # install_name_tool -change @rpath/libsqlite3.dylib @loader_path/libsqlite3.dylib $APP/Contents/MacOS/$exe
# install_name_tool -id @executable_path/../Frameworks/libsqlite3.dylib libsqlite3.dylib
# install_name_tool -change /usr/lib/libsqlite3.dylib @executable_path/../Frameworks/libsqlite3.dylib $exe

cp windows/data.db.bak $APP/Contents/MacOS/data.db
cp docs/ningmeng.icns $APP/Contents/Resources/ningmeng.icns
cat > $APP/Contents/Info.plist << EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleExecutable</key>
	<string>yunxing</string>
	<key>CFBundleIconFile</key>
	<string>ningmeng.icns</string>
	<key>CFBundleIdentifier</key>
	<string>com.zht.yunxing</string>
</dict>
</plist>
EOF
# find $APP
echo $APP over
