Destinado a centralizar comandos para este projeto

gRPC
Para criar arquivos (pasta pb)
$ protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto  

GraphQL
Para gerar arquivo tools.go
$ printf '//go:build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go

$ go mod tidy

$ go run github.com/99designs/gqlgen init

Depois vc altera seu schema.graphqls
Para refletir as alterações que vc fez no schema, use

$ go run github.com/99designs/gqlgen generate
