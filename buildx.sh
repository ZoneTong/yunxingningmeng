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
    
    rm -rf yunxingWinMac/macYunxing`date +%Y%m%d`.app
    mv -f macYunxing.app yunxingWinMac/macYunxing`date +%Y%m%d`.app
    mv winYunxing`date +%Y%m%d`.zip  yunxingWinMac/
    # mv -f yunxingWinMac`date +%Y%m%d`.zip yunxingWinMac.bak.zip
    # zip -rq yunxingWinMac.zip  yunxingWinMac
    # rm -rf yunxingWinMac
    echo pack end
}

function buildwin(){
     # 交叉编译 https://blog.csdn.net/tsxylhs/article/details/91852841
    # 首先 brew install mingw-w64 
    CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -H windowsgui -X main.trialday=36500 -X 'main.builddate=`date +%Y%m%d`'  " -o yunxing.exe

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

function nostaticGenAssets(){
    echo gen assets.go
    # 此时不生成静态库
    mkdir -p static
    mv -f www/css www/fonts www/js www/less www/ningmeng.ico www/excel.png static/

    # go run -tags generate gen.go
    go generate

    mv -f static/* www/
}

cd `dirname $0`

GEN=""
OS=""
while getopts 'g:s:' OPT; do
    case $OPT in
        g)
        GEN=$OPTARG
        # echo $OPTARG
        ;;
        s)
        OS=$OPTARG
        ;;
    esac
done
shift $((OPTIND-1))

# 平常调试界面: ./buildx.sh 或 ./buildx.sh -g 1
# 添加依赖库: ./buildx.sh -g 2
# 不编译界面: ./buildx.sh -g 0
if [ "$GEN" == "0" ]; then
    echo no gen assets.go
elif [ "$GEN" == "1" ] || [ "$GEN" == "" ]; then
    nostaticGenAssets
else   # 重新生成依赖库资源文件
    echo gen static.go
    go generate
    echo "package main"     > static.go
    echo                    >> static.go
    echo "func init() {"    >> static.go
    grep ".ico\|.png\|/js/\|/css/" assets.go >> static.go
    echo "}"                >> static.go
    nostaticGenAssets
fi

# ./build.sh -g 0 -s win
if [ "$OS" == "win" ] || [ "$OS" == "windows" ];then
    buildwin
elif [ "$OS" == "pack" ]; then
    pack
else
    go build -ldflags "-X main.trialday=2 -X 'main.builddate=`date +%Y%m%d`' " -o yunxingningmeng
    ./yunxingningmeng -v
    ./yunxingningmeng
fi