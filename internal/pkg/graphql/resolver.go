package graphql

import (
	"context"

	"github.com/iot-for-tillgenglighet/api-problemreport/pkg/database"
	"github.com/iot-for-tillgenglighet/api-problemreport/pkg/models"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Entity() EntityResolver {
	return &entityResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type entityResolver struct{ *Resolver }

func (r *entityResolver) FindProblemReportByID(ctx context.Context, id string) (*ProblemReport, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func convertEntityToGQL(entity *models.ProblemReport) *ProblemReport {
	if entity == nil {
		panic("Missing model")
	}

	resource := &ProblemReport{
		Pos: &WGS84Position{
			Lat: entity.Latitude,
			Lon: entity.Longitude,
		},
		Type: entity.Type,
	}

	return resource
}

func (r *mutationResolver) Create(ctx context.Context, input ProblemReportCreateResource) (*ProblemReport, error) {
	entity := &models.ProblemReport{
		Latitude:  input.Pos.Lat,
		Longitude: input.Pos.Lon,
		Type:      input.Type,
	}

	savedEntity, err := database.Create(entity)
	return convertEntityToGQL(savedEntity), err
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetAll(ctx context.Context) ([]*ProblemReport, error) {
	panic("not implemented")
}
