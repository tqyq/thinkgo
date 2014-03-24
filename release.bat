set MPATH=D:\git\cms_go
cd %MPATH%
taskkill /F /IM cms_go.exe /T
rm cms_go.exe
go fmt && go build
cms_go.exe
pause