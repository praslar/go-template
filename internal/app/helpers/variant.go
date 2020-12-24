package helpers

import (
	"context"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/google/uuid"

	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/app/types"
	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/pkg/common"
	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/pkg/logs"
	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/pkg/utils"
)

type VariantDatabase interface {
	Create(ctx context.Context, variant *types.Variant) error
	Update(ctx context.Context, variant *types.Variant) error
	Delete(ctx context.Context, variant *types.Variant) error
	GetOneByID(ctx context.Context, id string) (*types.Variant, error)
	GetAll(ctx beego.Controller) ([]*types.Variant, error)
}

type VariantHelper struct {
	variantDB VariantDatabase
}

func NewVariantHelper(db VariantDatabase) *VariantHelper {
	return &VariantHelper{
		db,
	}
}

func (s *VariantHelper) Delete(ctx context.Context, id uuid.UUID) error {
	logger := logs.WithContext(ctx)
	productAttr, err := s.variantDB.GetOneByID(ctx, id.String())
	if err != nil && !utils.IsErrNotFound(err) {
		logger.Errorf("Fail to get variant by key: %v", err)
		return utils.Gen("").Internal
	}
	if utils.IsErrNotFound(err) {
		return utils.Variant("").NotFound
	}

	if err := s.variantDB.Delete(ctx, productAttr); err != nil {
		return utils.Gen("").Database
	}

	return nil
}

func (s *VariantHelper) Create(ctx context.Context, req *types.CreateVariant, creatorID uuid.UUID) (*types.Variant, error) {
	logger := logs.WithContext(ctx)
	validator := validation.Validation{RequiredFirst: true}
	passed, err := validator.Valid(req)
	if err != nil {
		return nil, err
	}
	if !passed {
		var err string
		for _, e := range validator.Errors {
			err += fmt.Sprintf("[%s: %s] ", e.Field, e.Message)
		}
		return nil, utils.Gen(err).BadRequest
	}
	variant := types.Variant{
		BaseModel: types.BaseModel{
			CreatorID: creatorID,
			UpdaterID: creatorID,
		},
	}

	common.Sync(*req, &variant)

	if err := s.variantDB.Create(ctx, &variant); err != nil {
		logger.Errorf("Fail to insert variant to db: %s", err)
		return nil, utils.Gen("").Internal
	}

	return &variant, nil
}

func (s *VariantHelper) Update(ctx context.Context, id uuid.UUID, updaterID uuid.UUID, req *types.UpdateVariant) (*types.Variant, error) {
	logger := logs.WithContext(ctx)
	variant, err := s.variantDB.GetOneByID(ctx, id.String())
	if err != nil && !utils.IsErrNotFound(err) {
		logger.Errorf("Fail to get project by key: %v", err)
		return nil, utils.Gen("").Internal
	}
	if utils.IsErrNotFound(err) {
		return nil, utils.Variant("").NotFound
	}
	variant.BaseModel.UpdaterID = updaterID
	common.Sync(*req, variant)

	if err := s.variantDB.Update(ctx, variant); err != nil {
		logger.Errorf("Fail to update variant to db: %s", err)
		return nil, utils.Gen("").Database
	}

	return variant, nil
}

func (s *VariantHelper) GetOneByID(ctx context.Context, id uuid.UUID) (*types.Variant, error) {
	logger := logs.WithContext(ctx)
	variant, err := s.variantDB.GetOneByID(ctx, id.String())
	if err != nil && !utils.IsErrNotFound(err) {
		logger.Errorf("Fail to get variant by key: %v", err)
		return nil, utils.Gen("").Internal
	}
	if utils.IsErrNotFound(err) {
		return nil, utils.Variant("").NotFound
	}
	return variant, nil
}

func (s *VariantHelper) GetAll(ctx beego.Controller) ([]*types.Variant, error) {
	variant, err := s.variantDB.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return variant, nil
}
