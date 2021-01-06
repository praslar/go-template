package databases

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"go-template/internal/app/models"
)

var (
	// compile time type assertion
	_ models.BaseDatabase = &TempDatabase{}
)

type TempDatabase struct {
	db *gorm.DB
}

func (t TempDatabase) Create(ctx context.Context, req interface{}) error {
	panic("implement me")
	//return d.db.Create(variant).Error
}

func (t TempDatabase) Update(ctx context.Context, req interface{}) error {
	panic("implement me")
	//return d.db.Save(variant).Error
}

func (t TempDatabase) GetOneByID(ctx context.Context, id string) (interface{}, error) {
	panic("implement me")
	//variant := new(models.Temp)
	//if err := d.db.Where("id = ?", id).First(variant).Error; err != nil {
	//	return nil, err
	//}
	//return variant, nil
}

func (t TempDatabase) GetAll(ctx beego.Controller) ([]interface{}, error) {
	panic("implement me")
	//var variant []*models.Temp
	//
	//if dbErr, chainErr := common.With(ctx).
	//	WhereUUID("id", []string{"id", "=", "?"}).
	//	WhereUUID("creator_id", []string{"creator_id", "=", "?"}).
	//	WhereUUID("updater_id", []string{"updater_id", "=", "?"}).
	//	WhereTime("created_at", []string{"created_at", "=", "?"}).
	//	WhereStringAdv("admin_name", []string{"admin_name", "ilike", "?"}, "%", "%").
	//	WhereStringAdv("platform_key", []string{"platform_key", "ilike", "?"}, "%", "%").
	//	WhereSearch("search").
	//	WhereBool("active", []string{"active", "=", "?"}).
	//	Order("sort").
	//	Pagination(&models.Temp{}).
	//	Find(&variant); dbErr != nil || chainErr != nil {
	//	if dbErr != nil {
	//		return nil, utils.Gen(dbErr.Error()).BadRequest
	//	}
	//	if chainErr != nil {
	//		return nil, utils.Gen(chainErr.Error()).BadRequest
	//	}
	//}
	//
	//return variant, nil
}

func (t TempDatabase) Delete(ctx context.Context, req interface{}) error {
	panic("implement me")
	//return d.db.Delete(variant).Error
}

func NewTempDatabase(db *gorm.DB) *TempDatabase {
	return &TempDatabase{
		db: db,
	}
}
