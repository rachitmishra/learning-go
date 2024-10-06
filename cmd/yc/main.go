package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"reflect"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	h "rachitmishra.com/yc/cmd/handler"
	gg "rachitmishra.com/yc/generated/graph"
	pb "rachitmishra.com/yc/generated/proto/todo"
	d "rachitmishra.com/yc/generated/sql"
	g "rachitmishra.com/yc/graph"
)

const _defaultPort = "8080"

func runGrpcServer() {
	lis, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTodoServer(s, &h.Handler{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func runGqlServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = _defaultPort
	}

	srv := handler.NewDefaultServer(gg.NewExecutableSchema(gg.Config{Resolvers: &g.Resolver{}}))
	srv.AddTransport(&transport.SSE{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func runHttpGateway() {
	conn, err := grpc.NewClient(
		":5050",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	gwmux := runtime.NewServeMux()
	err = pb.RegisterTodoHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalf("failed to register: %v", err)
	}
	gwServer := &http.Server{
		Addr:    ":5051",
		Handler: gwmux,
	}
	gwServer.ListenAndServe()
}

func main() {
	go runGrpcServer()
	runGqlServer()
	runHttpGateway()
}

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := d.New(conn)

	// list all authors
	authors, err := queries.ListUsers(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	insertedUser, err := queries.CreateUser(ctx, d.CreateUserParams{
		Username:     "rachitmishra",
		Email:        "g.dot.1@gmail.com",
		PasswordHash: "xyz123",
	})
	if err != nil {
		return err
	}
	log.Println(insertedUser)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetUser(ctx, insertedUser.ID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedUser, fetchedAuthor))
	return nil
}
