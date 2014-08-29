ThinkGo
======

### 以下说明以windows环境下为例 ###

### 前提 ###

已经安装了go1.3以上版本的开发环境。

### 配置GOPATH ###

GOPATH git:

https://github.com/tqyq/gopath.git

假定将以上地址clone到d:\git\gopath，那么配置系统环境变量GOPATH=d:\git\gopath。

然后`运行GOPATH路径下面的install.bat`。

### 演示工程 ###

https://github.com/tqyq/cms_go.git

将上述工程clone到本地，`启动mongo数据库`，然后`运行工程目录下的run.bat`，默认监听3001端口，如果没有异常，在浏览器输入http://localhost:3001就可以看到页面。

运行成功后，之后修改go代码并保存，环境会自动编译和重新运行，不用手工重启，如果编译出现错误，修正后仍可自动运行。
