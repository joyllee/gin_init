# gin_init
搭建基于gin框架的基础组件

## 项目初始化
1.配置环境
  * 设置go module: `$ export GO111MODULE=on`
  * 设置代理: export GOPROXY=https://goproxy.cn
  * 拉取依赖包: go mod download

2.项目运行
  * 命令运行: go run main.go -c conf/dev/config.yaml
  * 脚本执行: 
       - cd cmd
       - sh build.sh
       - sh start.sh (为后台运行)
