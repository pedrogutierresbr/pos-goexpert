package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pedrogutierresbr/pos-goexpert/12-graphql/graph"
	"github.com/pedrogutierresbr/pos-goexpert/12-graphql/internal/database"
)

const defaultPort = "8080"

func main() {
	// criado conexão com o banco de dados
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	// criado arquivo categoria que manipula o banco de dados
	// criado arquivo course que manipula o banco de dados
	categoryDb := database.NewCategory(db)
	courseDb := database.NewCourse(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// injeto categoria no resolver
	// injeto course no resolver
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CategoryDB: categoryDb, CourseDb: courseDb,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// criando tables
// entrar no sqlite: sqlite3 data.db

// Categories: create table categories (id string, name string, description string);
// Courses: create table courses (id string, name string, description string, category_id string);
