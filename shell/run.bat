taskkill /F /IM cms_go.exe /T
cd D:\git\cms_go
rm cms_go.exe
go fmt & go build
cms_go.exe
pause