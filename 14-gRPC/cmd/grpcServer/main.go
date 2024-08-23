package main

import (
	"database/sql"
	"net"

	"github.com/pedrogutierresbr/pos-goexpert/14-grpc/internal/database"
	"github.com/pedrogutierresbr/pos-goexpert/14-grpc/internal/pb"
	"github.com/pedrogutierresbr/pos-goexpert/14-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

// client gRPC Evans --> programa feito em go, criado para interagir com gRPC
// go install github.com/ktr0731/evans@latest
// interagindo com o evans

// evans -r repl
// package pb
// service CategoryService
// call CreateCategory

// agora insere o name, exemplo: cat 3
// insere a description, exemplo: cat 3 desc

//  ou

// call ListCategories
