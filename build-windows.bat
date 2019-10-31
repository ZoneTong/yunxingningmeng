@echo off

echo pack begin
:: xcopy windows\ningmeng.ico www\ 	/Y 
:: go generate
:: rm www\ningmeng.ico
go build -ldflags "-H windowsgui -X main.trialday=30 -X 'main.builddate=%date:~0,4%%date:~5,2%%date:~8,2%'"  -o yunxing.exe
:: go build -ldflags "-H windowsgui -X main.trialday=2 -X 'main.builddate=20191023'"  -o yunxing.exe

mkdir yunxingmgr
xcopy  yunxing.exe yunxingmgr 	/Y 
xcopy  windows\* yunxingmgr 		/Y 
zip -q -r yunxing%date:~0,4%%date:~5,2%%date:~8,2%.zip  yunxingmgr
rmdir /s /q yunxingmgr

echo pack end