package graphql

import (
	"context"

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

}

func (r *mutationResolver) Create(ctx context.Context, input ProblemReportCreateResource) (*ProblemReport, error) {
	panic("not implemented")

}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetAll(ctx context.Context) ([]*ProblemReport, error) {
	panic("not implemented")
}
