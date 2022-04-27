package service

import (
	"context"
	"github.com/subedibimal/go-mini-assignment/graph/model"

	"gorm.io/gorm"
)

func UserRegister(ctx context.Context, input model.NewUser) (interface{}, error) {
	_, err := UserGetByEmail(ctx, input.Email)
	if err == nil {
		// if err != record not found
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}

	createdUser, err := UserCreate(ctx, input)
	if err != nil {
		return nil, err
	}

	token, err := JwtGenerate(ctx, createdUser.ID)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token": token,
	}, nil
}