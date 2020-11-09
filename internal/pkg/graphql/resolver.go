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
func (r *entityResolver) FindProblemReportCategoryByID(ctx context.Context, id string) (*ProblemReportCategory, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Create(ctx context.Context, input ProblemReportCreateResource) (*ProblemReport, error) {
	entity := &models.ProblemReport{
		Latitude:  input.Pos.Lat,
		Longitude: input.Pos.Lon,
		Type:      input.Type,
	}

	db, _ := database.ConnectToDB()
	savedEntity, err := db.Create(entity)
	return convertEntityToGQL(savedEntity), err
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetAll(ctx context.Context) ([]*ProblemReport, error) {
	db, _ := database.ConnectToDB()
	entities, err := db.GetAll()

	if err != nil {
		panic("Query failed: " + err.Error())
	}

	count := len(entities)

	if count == 0 {
		return []*ProblemReport{}, nil
	}

	resources := make([]*ProblemReport, 0, count)

	for _, v := range entities {
		resources = append(resources, convertEntityToGQL(&v))
	}

	return resources, nil
}

func (r *queryResolver) GetCategories(ctx context.Context) ([]*ProblemReportCategory, error) {
	db, _ := database.ConnectToDB()
	entities, err := db.GetCategories()

	if err != nil {
		panic("Query failed: " + err.Error())
	}

	count := len(entities)

	if count == 0 {
		return []*ProblemReportCategory{}, nil
	}

	resources := make([]*ProblemReportCategory, 0, count)

	for _, v := range entities {
		resources = append(resources, convertCategoryEntityToGQL(&v))
	}

	return resources, nil
}

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

func convertCategoryEntityToGQL(entity *models.ProblemReportCategory) *ProblemReportCategory {
	if entity == nil {
		panic("Missing model")
	}

	resource := &ProblemReportCategory{
		Label:      entity.Label,
		ReportType: entity.ReportType,
	}

	return resource
}
