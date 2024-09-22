Destinado a centralizar comandos para este projeto

Para criar arquivos do gRPC (pasta pb)
protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto  