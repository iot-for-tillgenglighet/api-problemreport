package handler

import (
	"compress/flate"
	"math"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	gql "github.com/iot-for-tillgenglighet/api-problemreport/internal/pkg/graphql"
	"github.com/iot-for-tillgenglighet/api-problemreport/pkg/database"
	"github.com/iot-for-tillgenglighet/api-problemreport/pkg/models"
	"github.com/iot-for-tillgenglighet/ngsi-ld-golang/pkg/datamodels/fiware"
	"github.com/iot-for-tillgenglighet/ngsi-ld-golang/pkg/ngsi-ld"
	"github.com/iot-for-tillgenglighet/ngsi-ld-golang/pkg/ngsi-ld/types"
	"github.com/rs/cors"

	log "github.com/sirupsen/logrus"
)

func Router() {

	mux := http.NewServeMux()

	port := os.Getenv("PROBLEMREPORT_API_PORT")
	if port == "" {
		port = "8880"
	}

	log.Printf("Starting api-problemreport on port %s.\n", port)

	mux.HandleFunc("/api/graphql/playground", handler.Playground("GraphQL playground", "/api/graphql"))
	mux.HandleFunc("/api/graphql", handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: &gql.Resolver{}})))

	c := cors.Default().Handler(mux)

	log.Fatal(http.ListenAndServe(":"+port, c))
}

//newRequestRouter creates and returns a new router wrapper
func newRequestRouter() *RequestRouter {
	router := &RequestRouter{impl: chi.NewRouter()}

	router.impl.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	// Enable gzip compression for ngsi-ld responses
	compressor := middleware.NewCompressor(flate.DefaultCompression, "application/json", "application/ld+json")
	router.impl.Use(compressor.Handler)
	router.impl.Use(middleware.Logger)

	return router
}

func createRequestRouter(contextRegistry ngsi.ContextRegistry, db database.Datastore) *RequestRouter {
	router := newRequestRouter()

	router.addGraphQLHandlers(db)
	router.addNGSIHandlers(contextRegistry)

	return router
}

//CreateRouterAndStartServing creates a request router, registers all handlers and starts serving requests
func CreateRouterAndStartServing(db database.Datastore) {

	contextRegistry := ngsi.NewContextRegistry()
	ctxSource := contextSource{db: db}
	contextRegistry.Register(ctxSource)

	router := createRequestRouter(contextRegistry, db)

	port := os.Getenv("PROBLEMREPORT_API_PORT")
	if port == "" {
		port = "8880"
	}

	log.Printf("Starting api-problemreport on port %s.\n", port)

	log.Fatal(http.ListenAndServe(":"+port, router.impl))
}

type contextSource struct {
	db database.Datastore
}

func convertDatabaseRecordToWeatherObserved(r *models.Temperature) *fiware.WeatherObserved {
	if r != nil {
		entity := fiware.NewWeatherObserved("temperature:"+r.Device, r.Latitude, r.Longitude, r.Timestamp)
		entity.Temperature = types.NewNumberProperty(math.Round(float64(r.Temp*10)) / 10)
		return entity
	}

	return nil
}

func (cs contextSource) GetEntities(query ngsi.Query, callback ngsi.QueryEntitiesCallback) error {

	var problemReport []models.ProblemReport
	var err error

	problemReport, err = cs.db.GetAll()

	if err == nil {
		for _, v := range problemReport {
			err = callback(convertDatabaseRecordToWeatherObserved(&v))
			if err != nil {
				break
			}
		}
	}

	return err
}

func (cs contextSource) ProvidesAttribute(attributeName string) bool {
	return attributeName == "temperature"
}

func (cs contextSource) ProvidesType(typeName string) bool {
	return typeName == "WeatherObserved"
}
