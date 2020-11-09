package context

import (
	"errors"
	"strings"

	"github.com/iot-for-tillgenglighet/api-problemreport/pkg/database"
	ngsi "github.com/iot-for-tillgenglighet/ngsi-ld-golang/pkg/ngsi-ld"
)

type contextSource struct {
	db database.Datastore
}

//CreateSource instantiates and returns a Fiware ContextSource that wraps the provided db interface
func CreateSource(db database.Datastore) ngsi.ContextSource {
	return &contextSource{db: db}
}

func (cs contextSource) CreateEntity(typeName, entityID string, req ngsi.Request) error {
	return errors.New("CreateEntity not supported for type " + typeName)
}

func (cs contextSource) GetEntities(query ngsi.Query, callback ngsi.QueryEntitiesCallback) error {
	var err error
	return err
}

func (cs contextSource) ProvidesAttribute(attributeName string) bool {
	return attributeName == "problemreport"
}

func (cs contextSource) ProvidesEntitiesWithMatchingID(entityID string) bool {
	return strings.HasPrefix(entityID, "urn:ngsi-ld:Open311ServiceRequest:")
}

func (cs contextSource) ProvidesType(typeName string) bool {
	return typeName == "ProblemReport"
}

func (cs contextSource) UpdateEntityAttributes(entityID string, req ngsi.Request) error {
	return errors.New("UpdateEntityAttributes is not supported by this service")
}

func queriedAttributesDoNotInclude(attributes []string, requiredAttribute string) bool {
	for _, attr := range attributes {
		if attr == requiredAttribute {
			return false
		}
	}

	return true
}
