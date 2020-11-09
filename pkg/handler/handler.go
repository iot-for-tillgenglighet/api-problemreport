package handler

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	gql "github.com/iot-for-tillgenglighet/api-problemreport/internal/pkg/graphql"
	"github.com/iot-for-tillgenglighet/api-problemreport/pkg/database"

	"github.com/rs/cors"

	log "github.com/sirupsen/logrus"
)

//RequestRouter wraps the concrete router implementation
type RequestRouter struct {
	impl *chi.Mux
}

//Router sets up and serves http router.
func CreateRouterAndStartServing(db database.Datastore) *RequestRouter {

	router := &RequestRouter{impl: chi.NewRouter()}

	gqlServer := handler.New(gql.NewExecutableSchema(gql.Config{Resolvers: &gql.Resolver{}}))
	gqlServer.AddTransport(&transport.POST{})
	gqlServer.Use(extension.Introspection{})

	router.impl.Use(database.Middleware(db))

	router.impl.Handle("/api/graphql/playground", playground.Handler("GraphQL playground", "/api/graphql"))
	router.impl.Handle("/api/graphql", gqlServer)

	port := os.Getenv("PROBLEMREPORT_API_PORT")
	if port == "" {
		port = "8880"
	}

	log.Printf("Starting api-problemreport on port %s.\n", port)

	c := cors.Default().Handler(router.impl)

	log.Fatal(http.ListenAndServe(":"+port, c))

	return router
}
