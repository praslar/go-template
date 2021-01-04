package helpers

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/google/uuid"
	"go-template/internal/app/models"
)

type tempDatabase interface {
	models.BaseDatabase
	// some func of Temp
}

var (
	// compile time type assertion
	_ models.BaseHelper = &TempHelper{}
)

type TempHelper struct {
	tempDB tempDatabase
}

func (t *TempHelper) Create(ctx context.Context, req interface{}) (interface{}, error) {
	panic("implement me")
	//logger := logs.WithContext(ctx)
	//validator := validation.Validation{RequiredFirst: true}
	//passed, err := validator.Valid(req)
	//if err != nil {
	//	return nil, err
	//}
	//if !passed {
	//	var err string
	//	for _, e := range validator.Errors {
	//		err += fmt.Sprintf("[%s: %s] ", e.Field, e.Message)
	//	}
	//	return nil, utils.Gen(err).BadRequest
	//}
	//variant := models.Temp{
	//	BaseModel: models.BaseModel{
	//		CreatorID: req.creatorID,
	//		UpdaterID: req.creatorID,
	//	},
	//}
	//
	//common.Sync(*req, &variant)
	//
	//if err := s.variantDB.Create(ctx, &variant); err != nil {
	//	logger.Errorf("Fail to insert variant to db: %s", err)
	//	return nil, utils.Gen("").Internal
	//}
	//
	//return &variant, nil
}

func (t *TempHelper) Update(ctx context.Context, req interface{}) (interface{}, error) {
	panic("implement me")
	//logger := logs.WithContext(ctx)
	//variant, err := s.variantDB.GetOneByID(ctx, req.id.String())
	//if err != nil && !utils.IsErrNotFound(err) {
	//	logger.Errorf("Fail to get project by key: %v", err)
	//	return nil, utils.Gen("").Internal
	//}
	//if utils.IsErrNotFound(err) {
	//	return nil, utils.Temp("").NotFound
	//}
	//variant.BaseModel.UpdaterID = req.updaterID
	//common.Sync(*req, variant)
	//
	//if err := s.variantDB.Update(ctx, variant); err != nil {
	//	logger.Errorf("Fail to update variant to db: %s", err)
	//	return nil, utils.Gen("").Database
	//}
	//
	//return variant, nil
}

func (t *TempHelper) GetOneByID(ctx context.Context, id uuid.UUID) (interface{}, error) {
	panic("implement me")
	//logger := logs.WithContext(ctx)
	//variant, err := s.variantDB.GetOneByID(ctx, id.String())
	//if err != nil && !utils.IsErrNotFound(err) {
	//	logger.Errorf("Fail to get variant by key: %v", err)
	//	return nil, utils.Gen("").Internal
	//}
	//if utils.IsErrNotFound(err) {
	//	return nil, utils.Temp("").NotFound
	//}
	//return variant, nil
}

func (t *TempHelper) GetAll(ctx beego.Controller) ([]interface{}, error) {
	panic("implement me")
	//variant, err := s.variantDB.GetAll(ctx)
	//if err != nil {
	//	return nil, err
	//}
	//return variant, nil
}

func (t *TempHelper) Delete(ctx context.Context, id uuid.UUID) error {
	panic("implement me")
	//logger := logs.WithContext(ctx)
	//productAttr, err := s.variantDB.GetOneByID(ctx, id.String())
	//if err != nil && !utils.IsErrNotFound(err) {
	//	logger.Errorf("Fail to get variant by key: %v", err)
	//	return utils.Gen("").Internal
	//}
	//if utils.IsErrNotFound(err) {
	//	return utils.Temp("").NotFound
	//}
	//
	//if err := s.variantDB.Delete(ctx, productAttr); err != nil {
	//	return utils.Gen("").Database
	//}
	//
	//return nil
}

func NewTempHelper(db tempDatabase) *TempHelper {
	return &TempHelper{
		db,
	}
}