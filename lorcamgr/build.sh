cd `dirname $0`
if [ "$1" == "" ]; then
    go run -tags generate gen.go
fi
go build
./lorcamgr