package main

import (
	"context"
	"flag"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tiquiet/mongoGraphql/graph"
	"github.com/tiquiet/mongoGraphql/graph/generated"
	"github.com/tiquiet/mongoGraphql/internal/repository/db"
	"github.com/tiquiet/mongoGraphql/pkg/mongo_client"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	certFile := flag.String("certfile", "cert.pem", "certificate PEM file")
	keyFile := flag.String("keyfile", "key.pem", "key PEM file")
	flag.Parse()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	Host := "dataBase"
	Port := "27017"
	Username := ""
	Password := ""
	Database := "books"
	//AuthDB := ""

	mongoClient, err := mongo_client.NewClient(context.Background(), Host, Port,
		Username, Password, Database)
	if err != nil {
		log.Fatal(err)
	}

	repository := db.NewStorage(mongoClient, "books")

	resolver := &graph.Resolver{
		Repository: repository,
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/graphql", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/graphql for GraphQL playground", port)
	log.Fatal(http.ListenAndServeTLS(":"+port, *certFile, *keyFile, nil))
}
