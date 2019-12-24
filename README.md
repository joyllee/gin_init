## 框架支持
1.日志文件打印 (可根据日期)

2.具有MySQL,mongodb,redis连接池 (连接之后需要关闭)

3.常用方法(util): 图片落盘, httpClient
 
## 项目初始化
1.配置环境
  * 设置go module: `$ export GO111MODULE=on`
  * 设置代理: `$ export GOPROXY=https://goproxy.cn`
  * 拉取依赖包: `$ go mod download`

2.项目运行
  * 命令运行: `$ go run main.go -c conf/dev/config.yaml`
  * 脚本运行: 
       - `$ cd cmd`
       - `$ sh build.sh`
       - `$ sh start.sh` (为后台运行)
