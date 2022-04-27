package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/subedibimal/go-mini-assignment/graph/generated"
	"github.com/subedibimal/go-mini-assignment/graph/model"
	"github.com/subedibimal/go-mini-assignment/service"
)

func (r *authOpsResolver) Register(ctx context.Context, obj *model.AuthOps, input model.NewUser) (interface{}, error) {
	return service.UserRegister(ctx, input)
}

func (r *mutationResolver) Auth(ctx context.Context) (*model.AuthOps, error) {
	return &model.AuthOps{}, nil
}

// AuthOps returns generated.AuthOpsResolver implementation.
func (r *Resolver) AuthOps() generated.AuthOpsResolver { return &authOpsResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type authOpsResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
