package users

import "context"

type Storage interface {
	Create(ctx context.Context, user User) (string, error)
	Retrieve(ctx context.Context, id string) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, user User) error
}
