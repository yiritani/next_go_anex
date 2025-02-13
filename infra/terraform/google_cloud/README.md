# Google Cloud

command
```
terraform apply -var="project_id=[project_id]"
```

## フォルダについて
当初はモノリスで建てる前提だったが、せっかくgRPCならマイクロサービスだろうということでAPI gatewayをかますことにした。
- development: frontendからそれぞれのバックエンドに繋がるモノリス構成
- microservices: API gatewayをかませるマイクロサービス構成