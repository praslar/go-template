package databases

import (
	"context"

	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/app/types"
	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/pkg/common"
	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/pkg/utils"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

type VariantDatabase struct {
	db *gorm.DB
}

func NewVariantDatabase(db *gorm.DB) *VariantDatabase {
	return &VariantDatabase{
		db: db,
	}
}

func (d *VariantDatabase) Create(ctx context.Context, variant *types.Variant) error {
	return d.db.Create(variant).Error
}

func (d *VariantDatabase) Update(ctx context.Context, variant *types.Variant) error {
	return d.db.Save(variant).Error
}

func (d *VariantDatabase) GetOneByID(ctx context.Context, id string) (*types.Variant, error) {
	variant := new(types.Variant)
	if err := d.db.Where("id = ?", id).First(variant).Error; err != nil {
		return nil, err
	}
	return variant, nil
}

func (d *VariantDatabase) GetAll(ctx beego.Controller) ([]*types.Variant, error) {
	var variant []*types.Variant

	if dbErr, chainErr := common.With(ctx).
		WhereUUID("id", []string{"id", "=", "?"}).
		WhereUUID("creator_id", []string{"creator_id", "=", "?"}).
		WhereUUID("updater_id", []string{"updater_id", "=", "?"}).
		WhereTime("created_at", []string{"created_at", "=", "?"}).
		WhereStringAdv("admin_name", []string{"admin_name", "ilike", "?"}, "%", "%").
		WhereStringAdv("platform_key", []string{"platform_key", "ilike", "?"}, "%", "%").
		WhereSearch("search").
		WhereBool("active", []string{"active", "=", "?"}).
		Order("sort").
		Pagination(&types.Variant{}).
		Find(&variant); dbErr != nil || chainErr != nil {
		if dbErr != nil {
			return nil, utils.Gen(dbErr.Error()).BadRequest
		}
		if chainErr != nil {
			return nil, utils.Gen(chainErr.Error()).BadRequest
		}
	}

	return variant, nil
}

func (d *VariantDatabase) Delete(ctx context.Context, variant *types.Variant) error {
	return d.db.Delete(variant).Error
}
