package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/rs/xid"
	"lai.com/GraphQL_Server/service"

	"lai.com/GraphQL_Server/graph/generated"
	"lai.com/GraphQL_Server/model"
)

func (r *mutationResolver) CreateEnterpriseInfo(ctx context.Context, input model.NewEnterpriseInfo) (*model.EnterpriseInfo, error) {
	enterpriseInfo := &model.EnterpriseInfo{
		EnterpriseName: input.EnterpriseName,
		EnterpriseID:   xid.New().String(),
	}
	err := service.AddEnterpriseInfo(enterpriseInfo)
	if err != nil {
		fmt.Printf("service.AddEnterpriseInfo() failed,err:%v\n", err)
	}
	return enterpriseInfo, err
}

func (r *queryResolver) EnterpriseInfoList(ctx context.Context) ([]*model.EnterpriseInfo, error) {
	enterpriseList, err := service.GetEnterpriseList()
	return enterpriseList, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
