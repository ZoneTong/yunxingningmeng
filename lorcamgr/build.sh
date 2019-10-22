function pack(){
    # refreshData函数,对接数据库
    # 整理图标
    # 去掉默认账号密码的填写
    echo pack begin
    echo pack end
}

cd `dirname $0`
if [ "$1" == "" ]; then
    go run -tags generate gen.go
fi
go build
./lorcamgr