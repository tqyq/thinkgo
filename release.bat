set MPATH=D:\baidu_sync\git\push_server\mpush_server
set GOPATH=%GOPATH%;%MPATH%
go build -o bin\push_srv.exe src\main.go
echo "done"