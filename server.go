package main

import (
	"os"
	"log"
	"net/http"
	"context"
	"encoding/json"

	dapr "github.com/dapr/go-sdk/client"
	// pb "github.com/dapr/dapr/pkg/proto/runtime/v1"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/zeimbeekor/go-dapr-gql-service/graph"
	"github.com/zeimbeekor/go-dapr-gql-service/graph/generated"
)

const (
	defaultPort = "7000"
	daprGRPCPort = "3501"
    bindingName = "graphql"
    bindingOperation = "query"
    commandQuery = "query"
)

type BodyRequest struct {
    // Name is name of binding to invoke
    Name 		string
    // Operation is the name of the operation type for the binding to invoke
    Operation	string
    // Metadata is the input binding metadata
    Metadata 	map[string]string
}

type Post struct {
  	Id			int
  	Title		string
	Body		string
}

type GraphQLResponse struct {
	Post		Post
}

// API Graphql service alongside a Dapr sidecar
func apiHandler(w http.ResponseWriter, r *http.Request) {
	port := os.Getenv("DARP_GRPC_PORT")
	if port == "" {
		port = daprGRPCPort
	}

	// Client instance
	client, err := dapr.NewClientWithPort(port)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	body := new(BodyRequest)
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Fatal(err)
	}

	// Using Dapr SDK to invoke output binding (graphql)
	in := &dapr.InvokeBindingRequest{
		Name:      body.Name,
		Operation: body.Operation,
		Metadata:  body.Metadata,
	}
	resp, err := client.InvokeBinding(context.Background(), in)

	var qglResponse map[string]interface{}
	if err := json.Unmarshal([]byte(resp.Data), &qglResponse); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&qglResponse)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	http.HandleFunc("/dapr-gql", apiHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
