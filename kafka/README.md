[docker-kafka] is being sponsored by the following tool; please help to support us by taking a look and signing up to a free trial
<a href="https://tracking.gitads.io/?repo=docker-kafka"> <img src="https://images.gitads.io/docker-kafka" alt="GitAds"/> </a>

## 项目说明
一站式Kafka环境部署，可用于本地开发环境或者测试环境，如果用于生产环境，请先确保服务器安全。

> 特性
* 支持通过修改 docker-compose.yml 文件中 kafka 的环境变量 KAFKA_VERSION, SCALA_VERSION 来自定义 Kafka 版本
* Kafka Server 启动配置放置在容器外部 conf/kafka-server.properties 文件中，可完全自定义配置（需要重新执行 build）
* Kafka 安装目录可以通过修改 docker-compose.yml 文件中 kafka 的环境变量 KAFKA_INSTALL_PATH 来设置，同时在内部创建了一个 /opt/kafka 的软连接来帮助快速定位 kafka 目录
* 独立 zookeeper 组件，不仅可为 Kafka 服务，同时也可作为其它服务的 zk 组件
* 集成最新版 KafkaManager，且支持通过修改 docker-compose.yml 文件中 kafka-manager 的环境变量设置是否开启 KafkaManager 认证，以及设置认证用户信息

> 使用手册
- build kafka image
```bash
docker-compose build --force-rm --no-cache
```
该命令通常情况下只需要执行一次，但是如果你之前已经 build 过，然后你想重新 build 或者修改了 kafka 的配置之后要重新执行 build，都可以直接使用该命令执行 build.

- 启动服务
```bash
docker-compose up --force-recreate -d
```

- 停止服务
```bash
docker-compose rm -svf
```
- 访问KafkaManager
```
http://localhost:9000
```

## 最后
欢迎提PR，不定时更新
