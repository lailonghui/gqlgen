package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"lai.com/GraphQL_Server/graph/generated"
	graph "lai.com/GraphQL_Server/graph/model"
)

func (r *mutationResolver) CreateVehicleInfo(ctx context.Context, input graph.NewVehicleInfo) (*graph.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DistrictVehicleList(ctx context.Context, input *graph.DefaultInput) ([]*graph.DistrictVehicle, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VehicleInfoList(ctx context.Context, input *graph.DefaultInput) ([]*graph.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
