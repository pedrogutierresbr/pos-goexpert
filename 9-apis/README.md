## Comandos interessantes


```
- Vai entrar no banco de dados, depois retorna o os dados

sqlite3 cmd/server/test.db

select * from products; 
```

```
- Criando docs do swaggo

swag init -g cmd/server/main.go

- tem que ser assim, apontando pro arquivo quandos e cria do root do repositório
- tem que apontar pro arquivo de bootstrap da aplicação

para acessar a doc, abra no navegador: http://localhost:8000/docs/index.html
```
