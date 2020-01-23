package graphql

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Entity() EntityResolver {
	return &entityResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type entityResolver struct{ *Resolver }

func (r *entityResolver) FindDeviceByID(ctx context.Context, id string) (*Device, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Snowdepths(ctx context.Context) ([]*Problemreport, error) {
	panic("not implemented")
}
