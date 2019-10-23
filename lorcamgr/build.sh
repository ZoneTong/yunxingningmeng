function pack(){
    # refreshData函数,对接数据库
    # pageSize
    # 整理图标
    # 去掉默认账号密码的填写
    echo pack begin
    echo pack end
}

cd `dirname $0`
if [ "$1" == "" ]; then
    # go run -tags generate gen.go
    go generate
fi
go build -ldflags "-X main.TRIAL_DAY=8 -X 'main.BUILD_TIME=`date +%Y%m%d`'  "
./lorcamgr