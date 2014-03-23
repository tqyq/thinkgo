set MPATH=D:\git\cms_go
set GOPATH=%GOPATH%;%MPATH%
cd %MPATH%
go build -o bin\main.exe src\main.go
bin\main.exe
pause