function pack(){
    # refreshData函数,对接数据库
    # pageSize
    # 整理图标
    # 去掉默认账号密码的填写
    echo pack begin
    echo pack end
}

cd `dirname $0`

NOGEN=0
OS=""
while getopts 'no:' OPT; do
    case $OPT in
        n)
        NOGEN=1;;
        o)
        OS=$OPTARG
        ;;
    esac
done
shift $((OPTIND-1))

if [ "$NOGEN" == "0" ]; then
    # go run -tags generate gen.go
    go generate
fi

if [ "$OS" == "win" ] || [ "$OS" == "windows" ];then
    # 交叉编译 https://blog.csdn.net/tsxylhs/article/details/91852841
    # 首先 brew install mingw-w64 
    CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -H windowsgui -X main.TRIAL_DAY=21 -X 'main.BUILD_TIME=`date +%Y%m%d`'  "

    # CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -x -v -ldflags "-s -w -H windowsgui -X main.TRIAL_DAY=2 -X 'main.BUILD_TIME=`date +%Y%m%d`'  " # 失败, gcc_libinit_windows.c:7:10: fatal error: 'windows.h' file not found
    # GOOS=windows GOARCH=amd64 go build -x -v -ldflags "-s -w -H windowsgui -X main.TRIAL_DAY=2 -X 'main.BUILD_TIME=`date +%Y%m%d`' "  #失败
    # xgo -x -v  --targets='windows/*' -ldflags "-s -w -H windowsgui -X main.TRIAL_DAY=21 -X 'main.BUILD_TIME=20191101' " .    # 成功, 以后考虑编译android ios
else
    go build -ldflags "-X main.TRIAL_DAY=2 -X 'main.BUILD_TIME=`date +%Y%m%d`'  "
    ./lorcamgr
fi
