package databases

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"go-template/internal/app/models"
)

var (
	// compile time type assertion
	_ models.BaseDatabase = &GraphQlDatabase{}
)

type GraphQlDatabase struct {
	db *gorm.DB
}

func (g GraphQlDatabase) Create(ctx context.Context, req interface{}) error {
	panic("implement me")
}

func (g GraphQlDatabase) Update(ctx context.Context, req interface{}) error {
	panic("implement me")
}

func (g GraphQlDatabase) GetOneByID(ctx context.Context, id string) (interface{}, error) {
	panic("implement me")
}

func (g GraphQlDatabase) GetAll(ctx beego.Controller) (interface{}, error) {
	panic("implement me")
}

func (g GraphQlDatabase) Delete(ctx context.Context, req interface{}) error {
	panic("implement me")
}

func NewGraphQLDatabase(db *gorm.DB) *GraphQlDatabase {
	return &GraphQlDatabase{
		db: db,
	}
}
