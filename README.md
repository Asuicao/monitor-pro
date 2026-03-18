# iFlow Monitor Pro

企业级 Kubernetes 集群监控平台

## 架构

- 前端: React + Nginx
- 后端: Go + Gin
- 数据库: MySQL 8.0
- 缓存: Redis
- 监控存储: VictoriaMetrics
- 日志聚合: Loki

## 快速开始

```bash
kubectl apply -k k8s/
```
