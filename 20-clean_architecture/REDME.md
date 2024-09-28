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


Para subir o db e rabbitMQ
abre um terminal
$ docker-compose up -d
$ docker-compose exec mysql bash
$ mysql -uroot -p orders
$ password: root
$ CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id));
$ select * from orders; 

Para rodar o projeto
cd cmd/ordersystem/
go run main.go wire_gen.go


Para mandar uma REQ HTTP
utilizar o create_order.http


Para mandar uma REQ gRPC
abre um terminal
$ evans -r repl
Pode ser que vc entre e o package e service não estejam selecionados
ai precisa selecionar manualmente
$ package <escolhe pb>
$ service <escolhe OrderService>
$ call CreateOrder
$ coloca os inputs pedidos


Para mandar via GraphQL
abrir navegador na 8080
cria uma mutation:

mutation createOrder{
    createOrder(input: {id: 'dafsdasf', Price: xx.xx, Tax x.x}) {
        id
        Price
        Tax
        FinalPrice
    }
}


RabbitMQ
localhost:15672
user: guest
password: guest

Queues
Add a new queue
Name: orders --> add queue
entrar na fila, clicando no nome
Bindings
From exchange: amq.direct
clicar em bind

Ai faz mais uma inserção no db e verifica no RabbitMQ se foi disparado o evento
fica no Queues
Get Message(s)