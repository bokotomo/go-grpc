# grpc-go
GRPCの実行サンプル

## 動かし方
```
cd ./infrastructure/docker
docker-compose up -d
docker-compose exec go-grpc bash
```

## unaryr RPC
２つのターミナルで、serverとclientを実行
```
go run ./app/unaryrpc/server/
go run ./app/unaryrpc/client/
```

## Server Streaming RPC
２つのターミナルで、serverとclientを実行
```
go run ./app/serverside/server/
go run ./app/serverside/client/
```

## Client Streaming RPC
２つのターミナルで、serverとclientを実行
```
go run ./app/clientside/server/
go run ./app/clientside/client/
```

## Bidirectional Streaming RPC
２つのターミナルで、serverとclientを実行
```
go run ./app/bidirectional/server/
go run ./app/bidirectional/client/
```

## create pb
```
// unary rpc 足算
make protoc OUT=./pb/calc NAME=calc.proto

// server streaming rpc 通知
make protoc OUT=./pb/notification NAME=notification.proto

// client streaming rpc アップロード
make protoc OUT=./pb/upload NAME=upload.proto

// client streaming rpc チャット
make protoc OUT=./pb/chat NAME=chat.proto
```

