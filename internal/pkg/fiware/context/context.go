package context

import (
	"errors"
	"strings"

	"github.com/iot-for-tillgenglighet/api-problemreport/pkg/database"
	"github.com/iot-for-tillgenglighet/ngsi-ld-golang/pkg/datamodels/fiware"
	ngsi "github.com/iot-for-tillgenglighet/ngsi-ld-golang/pkg/ngsi-ld"
)

type contextSource struct {
	db              database.Datastore
	serviceRequests []fiware.Open311ServiceRequest
}

//CreateSource instantiates and returns a Fiware ContextSource that wraps the provided db interface
func CreateSource(db database.Datastore) ngsi.ContextSource {
	return &contextSource{db: db}
}

func (cs contextSource) CreateEntity(typeName, entityID string, req ngsi.Request) error {

	serviceRequest := &fiware.Open311ServiceRequest{}
	err := req.DecodeBodyInto(serviceRequest)

	if err == nil {
		cs.serviceRequests = append(cs.serviceRequests, *serviceRequest)
	}

	return err
}

func (cs contextSource) GetEntities(query ngsi.Query, callback ngsi.QueryEntitiesCallback) error {

	var err error

	for _, serviceRequest := range cs.serviceRequests {
		err = callback(serviceRequest)
		if err != nil {
			break
		}
	}

	return err
}

func (cs contextSource) ProvidesAttribute(attributeName string) bool {
	return true
}

func (cs contextSource) ProvidesEntitiesWithMatchingID(entityID string) bool {
	return strings.HasPrefix(entityID, "urn:ngsi-ld:Open311ServiceRequest:")
}

func (cs contextSource) ProvidesType(typeName string) bool {
	return typeName == "Open311ServiceRequest"
}

func (cs contextSource) UpdateEntityAttributes(entityID string, req ngsi.Request) error {
	return errors.New("UpdateEntityAttributes is not supported by this service")
}
