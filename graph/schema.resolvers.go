package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/subedibimal/go-mini-assignment/graph/generated"
	"github.com/subedibimal/go-mini-assignment/graph/model"
	"github.com/subedibimal/go-mini-assignment/service"
	"github.com/subedibimal/go-mini-assignment/middlewares"
)

func (r *authOpsResolver) Login(ctx context.Context, obj *model.AuthOps, email string, password string) (interface{}, error) {
	return service.UserLogin(ctx, email, password)
}

func (r *authOpsResolver) Register(ctx context.Context, obj *model.AuthOps, input model.NewUser) (interface{}, error) {
	return service.UserRegister(ctx, input)
}

func (r *mutationResolver) Auth(ctx context.Context) (*model.AuthOps, error) {
	return &model.AuthOps{}, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return service.UserGetByID(ctx, id)
}

func (r *queryResolver) Getuser(ctx context.Context) (*model.User, error) {
	tokenData := middlewares.CtxValue(ctx)
	id := tokenData.ID
	return service.UserGetByID(ctx, id)
}

// AuthOps returns generated.AuthOpsResolver implementation.
func (r *Resolver) AuthOps() generated.AuthOpsResolver { return &authOpsResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type authOpsResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
