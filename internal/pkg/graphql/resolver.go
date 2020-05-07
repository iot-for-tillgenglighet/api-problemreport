package graphql

import (
	"context"
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
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetAll(ctx context.Context) ([]*ProblemReport, error) {
	panic("not implemented")
}
func (r *queryResolver) GetCategories(ctx context.Context) ([]*ProblemReportCategory, error) {
	panic("not implemented")
}
