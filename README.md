# job

## 项目描述

job项目是一个中心化任务分发执行服务，基于CS架构

### jobserver

中心服务，负责接收用户分发的任务，任务/结果的存储

### joblet

运行在节点上，启动时自动注册节点信息到job server，定时从job server获取当前节点的任务并执行，执行完成后将结果推送回job server

定时获取任务成功后，server会判定该节点存活

## 架构

![image-20220616140106895](images/readme-1.png)

## 接口

### jobserver

#### 任务创建

URL: http://<ip:port>/api/task/create

method: post

content-type: application/json

| 字段    | 数据类型 | 描述               | 是否必须 |
| ------- | -------- | ------------------ | -------- |
| name    | string   | 任务名             | 必须     |
| target  | string   | 任务执行的服务器ip | 必须     |
| command | string   | 执行的命令         | 必须     |
| cron    | bool     | 是否定时           | 必须     |
| runtime | int64    | 定时运行的时间     | 必须     |

## 快速开始

### 获取源码

```shell
git clone https://github.com/fixJ/job
```



### 编译

```shell
go build -o jobserver cmd/jobserver/jobserver.go
go build -o joblet cmd/joblet/joblet.go
```



### 运行

#### jobserver

需要先提前安装MySQL，并创建库表，建表语句在init/task.sql中

```shell
go run .\cmd\jobserver\jobserver.go --serverHost=<监听的ip> --serverPort=<监听的端口> --dbHost=<数据库ip> --dbPort=<数据库端口> --dbUsername=<数据库用户名> --dbPassword=<数据库密码> --dbName=<db名>
```



#### joblet

```shell
go run .\cmd\joblet\joblet.go --server http://<jobserver ip+port> --ip=<joblet ip>
```

