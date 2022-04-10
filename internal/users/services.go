package users

import (
	"context"
)

type Service struct {
	storage Storage
}

func Create(ctx context.Context, user_dto UserDTO) (User, error) {
	panic("pass")
}
