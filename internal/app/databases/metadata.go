package databases

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/app/types"
	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/pkg/common"
	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/pkg/utils"
)

type MetadataDatabase struct {
	db *gorm.DB
}

func NewMetadataDatabase(db *gorm.DB) *MetadataDatabase {
	return &MetadataDatabase{
		db: db,
	}
}

func (d *MetadataDatabase) GetListMetaFromType(ctx beego.Controller) (res []*types.MetaData, err error) {
	var metadata []*types.MetaData

	if dbErr, chainErr := common.With(ctx).
		WhereStringAdv("type", []string{"type", "ilike", "?"}, "%", "%").
		WhereStringAdv("platform_key", []string{"platform_key", "ilike", "?"}, "%", "%").
		WhereBool("is_active", []string{"is_active", "=", "?"}).
		Order("sort").
		Pagination(&types.Variant{}).
		Find(&metadata); dbErr != nil || chainErr != nil {
		if dbErr != nil {
			return nil, utils.Gen(dbErr.Error()).BadRequest
		}
		if chainErr != nil {
			return nil, utils.Gen(chainErr.Error()).BadRequest
		}
	}

	return metadata, nil
}
