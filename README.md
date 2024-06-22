```
.
├── cmd
│   └── api
│       └── main.go  # 服务入口文件，包含HTTP服务初始化等
├── configs
│   ├── config.yaml  # 默认配置文件
│   └── local.yaml   # 本地开发配置文件
├── internal
│   ├── database
│   │   ├── database.go  # 数据库模型定义
│   │   └── operation.go    # 数据库相关操作
│   ├── logger
│   │   └── logger.go    # 日志记录相关代码
│   ├── routes
│   │   └── routes.go    # HTTP路由定义
│   ├── services
│   │   ├── good.go     # 物品服务相关逻辑
│   │   └── graph.go     # 物流节点服务相关逻辑，包括Dijstra实现
|   |   |__ sorter.go    # 排序器
│   └── utils
│       ├── display.go  # 展示物品信息
│       └── utils.go     # 工具函数
├── pkg
│   └── apis
│       ├── items.go     # 物品API定义
│       └──  
├── scripts
│   └── start.sh         # 启动脚本
├── static                # 静态资源目录
│   └── ...
├── testdata             # 测试数据目录
│   └── ...
├── web
│   └── frontend         # Vue.js前端项目目录
│       ├── public
│       ├── src
│       └── package.json
├── .gitignore           # Git忽略文件
├── go.
```
