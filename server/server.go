package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/bygui86/go-graphql/server/graph"
	"github.com/bygui86/go-graphql/server/graph/generated"
	"github.com/bygui86/go-graphql/server/graph/model"
)

const (
	port = "8080"

	dialect    = "mysql"
	dbUsername = "root"
	dbPassword = "supersecret"
)

var db *gorm.DB

func main() {
	initDB()

	playHandler := playground.Handler("GraphQL playground", "/query")
	http.Handle("/", playHandler)

	resolver := &graph.Resolver{
		DB: db,
	}
	cfg := generated.Config{
		Resolvers: resolver,
	}
	execSchema := generated.NewExecutableSchema(cfg)
	sampleHandler := handler.NewDefaultServer(execSchema)
	http.Handle("/query", sampleHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func initDB() {
	var err error
	db, err = gorm.Open(
		dialect,
		fmt.Sprintf("%s:%s@tcp(localhost:3306)/?parseTime=True", dbUsername, dbPassword),
	)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.LogMode(true)

	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	db.Exec("CREATE DATABASE test_db")
	db.Exec("USE test_db")

	// Migration to create tables for Order and Item schema
	// /!\	AutoMigrate will ONLY create tables, missing columns and missing indexes, and WON’T change existing
	//		column’s type or delete unused columns to protect your data.
	db.AutoMigrate(&model.Order{}, &model.Item{})
}
