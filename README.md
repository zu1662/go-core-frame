# 启动说明

## 服务端启动说明

```bash
# 进入 go-admin 后端项目
cd ./go-core-frame

# 编译项目
go build

# 修改配置 
# 文件路径  go-admin/config/settings.yml
vi ./config/setting.yml 

# 1. 配置文件中修改数据库信息 
# 注意: settings.database 下对应的配置数据
# 2. 确认log路径
```

## 初始化数据库，以及服务启动
```
# 首次配置需要初始化数据库资源信息
导入 `config/gocore.sql`, 生成数据库信息


# 启动项目，也可以用IDE进行调试
`go run main.go`

```

## 文档生成

```bash
swag init  

# 如果没有swag命令 go get安装一下即可
go get -u github.com/swaggo/swag/cmd/swag
```

## 交叉编译
```bash
env GOOS=windows GOARCH=amd64 go build main.go

# or

env GOOS=linux GOARCH=amd64 go build main.go
```