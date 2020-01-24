# grpc-go
GRPCの実行サンプル

## 動かし方
```
cd ./infrastructure/docker
docker-compose up -d
docker-compose exec go-grpc bash
```

## Simple RPC
２つのターミナルで、serverとclientを実行
```
go run ./simplerpc/server/
go run ./simplerpc/client/
```

## create pb
```
// simplerpc 足算
make protoc OUT=./pb/calc NAME=calc.proto
```

