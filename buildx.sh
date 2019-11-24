function pack(){
    # refreshData函数,对接数据库
    # pageSize
    # 整理图标
    # 去掉默认账号密码的填写
    echo pack begin
    ./build-macos.sh
    buildwin
    mkdir -p yunxingWinMac
    cp docs/运兴合作社管理软件使用说明书.docx yunxingWinMac/
    mv -f macYunxing.app yunxingWinMac/
    mv winYunxing`date +%Y%m%d`.zip  yunxingWinMac/
    mv -f yunxingWinMac.zip yunxingWinMac.bak.zip
    # zip -m -q -r yunxingWinMac.zip  yunxingWinMac
    # rm -rf yunxingWinMac
    echo pack end
}

function buildwin(){
     # 交叉编译 https://blog.csdn.net/tsxylhs/article/details/91852841
    # 首先 brew install mingw-w64 
    CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -H windowsgui -X main.trialday=90 -X 'main.builddate=`date +%Y%m%d`'  " -o yunxing.exe

    # CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -x -v -ldflags "-s -w -H windowsgui -X main.trialday=2 -X 'main.builddate=`date +%Y%m%d`'  " # 失败, gcc_libinit_windows.c:7:10: fatal error: 'windows.h' file not found
    # GOOS=windows GOARCH=amd64 go build -x -v -ldflags "-s -w -H windowsgui -X main.trialday=2 -X 'main.builddate=`date +%Y%m%d`' "  #失败
    # xgo -x -v  --targets='windows/*' -ldflags "-s -w -H windowsgui -X main.trialday=21 -X 'main.builddate=20191101' " .    # 成功, 以后考虑编译android ios

    APP=winYunxing
    mkdir -p yunxingmgr
    mv yunxing.exe yunxingmgr/
    cp -rf windows/* yunxingmgr/
    zip  -m -j -q -r $APP`date +%Y%m%d`.zip yunxingmgr
    rm -rf yunxingmgr
    echo $APP over
}

cd `dirname $0`

NOGEN=0
OS=""
while getopts 'ns:' OPT; do
    case $OPT in
        n)
        NOGEN=1;;
        s)
        OS=$OPTARG
        ;;
    esac
done
shift $((OPTIND-1))

if [ "$NOGEN" == "0" ]; then
    echo go generate 
    # go run -tags generate gen.go
    cd ui && mv ../www/unpackage ..
    go generate 
    mv ../unpackage ../www/ && cd -
fi

# ./build.sh -n -s win
if [ "$OS" == "win" ] || [ "$OS" == "windows" ];then
    buildwin
elif [ "$OS" == "pack" ]; then
    pack
else
    echo go build
    go build -ldflags "-X main.trialday=2 -X 'main.builddate=`date +%Y%m%d`' " -o yunxingningmeng .
    ./yunxingningmeng -v
    ./yunxingningmeng
fi