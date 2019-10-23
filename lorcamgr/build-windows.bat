@echo off

echo pack begin
xcopy icons\ningmeng.ico www\ 	/Y 
go generate
rm www\ningmeng.ico
go build -ldflags "-H windowsgui -X main.TRIAL_DAY=21 -X 'main.BUILD_TIME=%date:~0,4%%date:~5,2%%date:~8,2%'"  -o yunxing.exe
:: go build -ldflags "-H windowsgui -X main.TRIAL_DAY=2 -X 'main.BUILD_TIME=20191023'"  -o yunxing.exe

mkdir yunxingmgr
xcopy  yunxing.exe yunxingmgr 	/Y 
xcopy  icons\* yunxingmgr 		/Y 
zip -q -r ‘À–Àƒ˚√ .zip  yunxingmgr
rmdir /s /q yunxingmgr

echo pack end