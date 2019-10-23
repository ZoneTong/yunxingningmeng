@echo off

echo pack begin
move icons/ningmeng.ico www/

go generate
move www/ningmeng.ico icons/
go build -ldflags "-H windowsgui -X main.TRIAL_DAY=15 -X 'main.BUILD_TIME=`date +%Y%m%d`'"  -o yunxing.exe
mkdir yunxingmgr
copy yunxing.exe yunxingmgr
copy icons/* yunxingmgr
zip yunxingmgr

echo pack end