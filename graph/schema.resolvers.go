package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"lai.com/GraphQL_Server/service"

	"lai.com/GraphQL_Server/graph/generated"
	"lai.com/GraphQL_Server/model"
)

func (r *mutationResolver) CreateEnterpriseInfo(ctx context.Context, input model.NewEnterpriseInfo) (*model.EnterpriseInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) EnterpriseInfoList(ctx context.Context) ([]*model.EnterpriseInfo, error) {
	//operationContext := graphql.GetOperationContext(ctx)
	//fmt.Println(operationContext.OperationName)
	//graphql.CollectFieldsCtx() 可以获取到查询的字段
	fieldsCtx := graphql.CollectFieldsCtx(ctx, nil)
	var fieldNames []string
	for _, v := range fieldsCtx {
		fieldNames = append(fieldNames, v.Name)
	}
	enterpriseList, err := service.GetEnterpriseList(fieldNames)
	return enterpriseList, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
