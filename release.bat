cd D:\git\cms_go
taskkill /F /IM cms_go.exe /T
go fmt && go clean && go build
cms_go.exe
pause