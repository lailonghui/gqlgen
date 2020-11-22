package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"lai.com/GraphQL_Server/graph/generated"
	graph "lai.com/GraphQL_Server/graph/model"
)

func (r *mutationResolver) CreateVehicleInfo(ctx context.Context, input *graph.NewVehicleInfo) (*graph.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleInfo(ctx context.Context, input *graph.NewVehicleInfo) (*graph.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVehicleInfo(ctx context.Context, input *graph.NewVehicleInfo) (*graph.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateDriverInfo(ctx context.Context, input *graph.NewDriverInfo) (*graph.DriverInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDriverInfo(ctx context.Context, input *graph.NewDriverInfo) (*graph.DriverInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDriverInfo(ctx context.Context, input *graph.NewDriverInfo) (*graph.DriverInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateVehicleDriverBinding(ctx context.Context, input *graph.NewVehicleDriverBinding) (*graph.VehicleDriverBinding, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleDriverBinding(ctx context.Context, input *graph.NewVehicleDriverBinding) (*graph.VehicleDriverBinding, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVehicleDriverBinding(ctx context.Context, input *graph.NewVehicleDriverBinding) (*graph.VehicleDriverBinding, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetDistrictVehicleList(ctx context.Context, input *graph.DefaultInput) ([]*graph.DistrictCount, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetVehicleInfoList(ctx context.Context, input *graph.DefaultInput) ([]*graph.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetDistrictDriverList(ctx context.Context, input *graph.DefaultInput) ([]*graph.DistrictCount, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetDriverInfoList(ctx context.Context, input *graph.DefaultInput) ([]*graph.DriverInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetVehicleDriverBinding(ctx context.Context, input *graph.NewVehicleDriverBinding) ([]*graph.VehicleDriverBinding, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetVehicleAlarmDataList(ctx context.Context, input *graph.DefaultInput) ([]*graph.VehicleAlarmData, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetDistrictVehicleViolationList(ctx context.Context, input *graph.DefaultInput) ([]*graph.DistrictCount, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetVehicleViolatoinDetailList(ctx context.Context, input *graph.DefaultInput) ([]*graph.VehicleViolationDetailsOutput, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
