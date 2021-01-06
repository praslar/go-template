package models

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/google/uuid"
)

type BaseDatabase interface {
	Create(ctx context.Context, req interface{}) error
	Update(ctx context.Context, req interface{}) error
	GetOneByID(ctx context.Context, id string) (interface{}, error)
	GetAll(ctx beego.Controller) ([]interface{}, error)
	Delete(ctx context.Context, req interface{}) error
}

type BaseHelper interface {
	Create(ctx context.Context, req interface{}) (interface{}, error)
	Update(ctx context.Context, req interface{}) (interface{}, error)
	GetOneByID(ctx context.Context, id uuid.UUID) (interface{}, error)
	GetAll(ctx beego.Controller) ([]interface{}, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type BaseController interface {
	Post()
	Put()
	GetOne()
	GetAll()
	Delete()
}
