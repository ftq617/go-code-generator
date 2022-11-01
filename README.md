## Golang 代码生成器

#### v1.0.0

### 简介
>
> 根据定制的模板生成对应数据库表的go代码
>
> 可以按照自己项目结构生成不同的目录结构
>
> 现在只支持MySQL

### 设置-项目配置
```yaml
database:
    url: localhost
    ip: localhost
    port: 3306
    database: disaster_backup
    username: root
    password: 123456
project:
    abbr: "internal"        # 代码在项目下的目录
    mod:  "disaster_backup" # 项目 mod 名
    router: "disaster_backup" # 模块路由名
    path: "C:\\lt_file\\lt_file\\project\\disaster_backup"  # 代码生成路径

```

### 项目使用
```go run main.go table list```   
查看数据库列表
```go run main.go table show <table_name>```
查看表结构
```go run main.go code create -n <table_name> ```
生成表代码，可多个





