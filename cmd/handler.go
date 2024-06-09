package cmd

import (
	"log"
	"net/http"

	"github.com/anduckhmt146/graphql-api/internal/db"
	"github.com/anduckhmt146/graphql-api/internal/schema"
	"github.com/anduckhmt146/graphql-api/internal/services"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/spf13/viper"
)

func Start() {
	// Initialize the database connection (assuming db.SetupDB exists and configures a *gorm.DB)
	database, err := db.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	// Initialize the UserService with the database connection
	userService := services.NewUserService(database)

	// Create GraphQL schema configuration
	schemaConfig := graphql.SchemaConfig{
		Query:    schema.NewQueryType(userService),
		Mutation: schema.NewMutationType(userService),
	}

	// Create a new GraphQL schema with the provided configuration
	graphqlSchema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create GraphQL schema: %s", err)
	}

	// Set up a GraphQL HTTP handler with the configured schema
	h := handler.New(&handler.Config{
		Schema:   &graphqlSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	port := viper.GetString("service.port")

	// Set up the HTTP server to serve the GraphQL endpoint
	http.Handle("/graphql", h)
	log.Println("GraphQL API server running on http://localhost:8080/graphql")
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
