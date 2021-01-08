package databases

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"go-template/internal/app/models"
	"go-template/internal/pkg/common"
	"go-template/internal/pkg/utils"
)

func NewDemoDatabase(db *gorm.DB) *DemoDatabase {
	return &DemoDatabase{
		db: db,
	}
}

var (
	// compile time type assertion
	_ models.BaseDatabase = &DemoDatabase{}
)

type DemoDatabase struct {
	db *gorm.DB
}

func (d DemoDatabase) Create(ctx context.Context, req interface{}) error {
	return d.db.Create(req).Error
}

func (d DemoDatabase) Update(ctx context.Context, req interface{}) error {
	return d.db.Save(req).Error
}

func (d DemoDatabase) GetOneByID(ctx context.Context, id string) (interface{}, error) {
	demo := new(models.Demo)
	if err := d.db.Where("id = ?", id).First(demo).Error; err != nil {
		return nil, err
	}
	return demo, nil
}

func (d DemoDatabase) GetAll(ctx beego.Controller) (interface{}, error) {
	var demo []*models.Demo
	if dbErr, chainErr := common.With(ctx).
		WhereUUID("id", []string{"id", "=", "?"}).
		WhereUUID("creator_id", []string{"creator_id", "=", "?"}).
		WhereUUID("updater_id", []string{"updater_id", "=", "?"}).
		WhereTime("created_at", []string{"created_at", "=", "?"}).
		WhereStringAdv("name", []string{"name", "ilike", "?"}, "%", "%").
		WhereInt("age", []string{"age", "=", "?"}).
		Order("sort").
		Pagination(&models.Demo{}).
		Find(&demo); dbErr != nil || chainErr != nil {
		if dbErr != nil {
			return nil, utils.Gen(dbErr.Error()).BadRequest
		}
		if chainErr != nil {
			return nil, utils.Gen(chainErr.Error()).BadRequest
		}
	}
	return demo, nil
}

func (d DemoDatabase) Delete(ctx context.Context, req interface{}) error {
	return d.db.Delete(req).Error
}
