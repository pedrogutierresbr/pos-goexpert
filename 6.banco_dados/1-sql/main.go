// Aulas
// preparando base do código
// inseridno dados no banco (create)
// alterando dados no banco (update)
// trabalhando com QueryRow (read)
// selecionando múltiplos registros
// removendo registros (delete)
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // driver do mysql, o _ é para manter esse pacote aqui quando não estiver usando, não remove na compilação
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	// estabelecendo conexão com o db (tipo do banco, user, password, @protocolo, endereço, /nome do banco)
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// criando um novo produto
	product := NewProduct("Notebook", 1899.90)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	product.Price = 1980.00
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	// p, err := selectProduct(db, product.ID)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Product: %v, possui o preço de %.2f", p.Name, p.Price)

	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("Product: %v, possui o preço de %.2f\n", p.Name, p.Price)
	}

	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	// maneira de proteger o código contra sql injection (Prepare statement)
	// vai inserir em products id, name e price, mas vai substituir os values por ?, para sanitizar (não deixar entrar código malicioso)
	// cada banco usa um identificado para sanitizar, mysql usa ?, sqlite usa $1 , $2 .... (ver doc quando for usar)
	stmt, err := db.Prepare("insert into products(id, name, price) values(?, ?, ?)") // linha que prepara a ação a ser executada
	if err != nil {
		return err
	}
	defer stmt.Close()

	// usei o blank identifier, pois não quero pegar mais dados da operação (o que aconteceu com a transação), apenas executar a operação de insert
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?") // linha que prepara a ação a ser executada
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var p Product
	// agora vamos executar uma busca com o id, pegar o resultado e jogar em cada propriedade de p
	// QueryRow busca apenas um resultado
	// Scan olha o valor de cada coluna, e atribui o valor de cada coluna com cada campo correspondente do p
	// no Scan passamos a referência, ou seja o local na memória onde o p esta armazenado
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	// se tiver tudo certo, retorna o ponteiro de p (apontamento na memória onde esta p, e nil referente ao error lá de cima)
	// não retorno uma cópia de p, estou retornando exatamente aonde o valor de p esta guardado
	return &p, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	// não é ncessário usar o Prepare pois não passamos parâmetros
	// Query vai buscar os registros (busca mais de uma row)
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() { // vai varrer todos os registros
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

// Query quando quer buscar algo
// Exec quando quer executar algo
