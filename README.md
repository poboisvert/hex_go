# hex_go

```
protoc --go-grpc_out=require_unimplemented_servers=false:internal/adapters/framework/left/grpc --proto_path=internal/adapters/framework/left/grpc/proto internal/adapters/framework/left/grpc/proto/arithmetic_svc.proto
```

```
protoc --go_out=internal/adapters/framework/left/grpc --proto_path=internal/adapters/framework/left/grpc/proto internal/adapters/framework/left/grpc/proto/number_msg.proto
```

Base on https://www.youtube.com/watch?v=MpFog2kZsHk
