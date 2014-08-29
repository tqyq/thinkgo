ThinkGo
======

ThinkGo是一个仿照ThinkPHP方式的go web开发示例，使用框架[beego](http://beego.me/)。
示例包含了常见的web开发应用场景，数据的操作目前只支持mongodb的CRUD操作，方式上仿照ThinkPHP的D()，I()，S()等函数，简化了go web开发。
代码片段：

	func (this *Action) UserList() {
		start, rows := this.PageParam("page", "rows")
		p := this.F2p().Rm("page", "rows").Like("name")
		total := D(User).Find(p).Count()
		ps := D(User).Find(p).Skip(start).Limit(rows).Sort("-name").All()
		this.EchoJson(&P{"total": total, "rows": ps})
	}
	
	func (this *Action) UserAdd() {
		p := this.F2p()
		D(User).Add(p)
		this.EchoJsonOk()
	}
	
	func (this *Action) UserUpdate() {
		p := this.F2p()
		D(User).Save(p)
		this.EchoJsonOk()
	}
	
	func (this *Action) UserDel() {
		ids := this.Is("ids[]")
		Debug(ids)
		for _, v := range ids {
			D(User).RemoveId(v)
		}
		this.EchoJsonOk()
	}

### 以下安装说明以windows环境下为例 ###

### 前提 ###

已经安装了go1.3以上版本的开发环境。

### 配置GOPATH ###

GOPATH git:

https://github.com/tqyq/gopath.git

假定将以上地址clone到d:\git\gopath，那么配置系统环境变量GOPATH=d:\git\gopath。

然后`运行GOPATH路径下面的install.bat`。

以上步骤如果运行无误，将%GOPATH%\bin加入到系统PATH里面，这样可以在任意位置运行bee.exe。

### 演示工程 ###

https://github.com/tqyq/cms_go.git

将上述工程clone到本地，`启动mongo数据库`，然后`运行工程目录下的run.bat`，默认监听3001端口，如果没有异常，在浏览器输入http://localhost:3001就可以看到页面。

![alt tag](https://github.com/tqyq/cms_go/blob/master/public/img/tg.jpg?raw=true)

运行成功后，之后修改go代码并保存，环境会自动编译和重新运行，不用手工重启，如果编译出现错误，修正后仍可自动运行。
