<!-- Comando para criar arquivos sql -->
<!-- migrate create -ext=sql -dir=sql/migrations -seq init -->

<!-- Rodar docker -->
<!-- docker-compose up -d -->

<!-- Rodar migration -->
<!-- migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up -->

<!-- Verficar se comando acima deu certo -->
<!-- abra novo bash -->
<!-- docker-compose exec mysql bash -->
<!-- mysql -uroot -p courses -->
<!-- digite a senha: root -->
<!-- show tables; -->
<!-- desc courses; -->
<!-- desc categories; -->
<!-- verificar o que foi criado -->

<!-- Derrubar migration -->
<!-- migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down -->



<!-- Criado makelfile para auxiliar -->
<!-- para utilizar, segue os comandos: -->

<!-- make migrate -->
<!-- make migratedown -->


<!-- Gerar as querys pelo SQLC -->
<!-- sqlc generate -->