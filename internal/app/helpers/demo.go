package helpers

import (
	"context"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/google/uuid"
	"go-template/internal/app/models"
	"go-template/internal/pkg/common"
	"go-template/internal/pkg/logs"
	"go-template/internal/pkg/utils"
)

func NewDemoHelper(db demoDatabase) *DemoHelper {
	return &DemoHelper{
		db,
	}
}

type demoDatabase interface {
	models.BaseDatabase
	// some func of Temp
}

var (
	// compile time type assertion
	_ models.BaseHelper = &DemoHelper{}
)

type DemoHelper struct {
	demoDB demoDatabase
}

func (d DemoHelper) Create(ctx context.Context, req interface{}, creatorID uuid.UUID) (interface{}, error) {
	reqDemo := req.(*models.CreateDemoRequest)
	logger := logs.WithContext(ctx)
	validator := validation.Validation{RequiredFirst: true}
	passed, err := validator.Valid(reqDemo)
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
	demo := models.Demo{
		BaseModel: models.BaseModel{
			CreatorID: creatorID,
		},
	}

	common.Sync(*reqDemo, &demo)

	if err := d.demoDB.Create(ctx, &demo); err != nil {
		logger.Errorf("Fail to insert demo to db: %s", err)
		return nil, utils.Gen("").Internal
	}

	return &demo, nil
}

func (d DemoHelper) Update(ctx context.Context, req interface{}, objectID uuid.UUID, updaterID uuid.UUID) (interface{}, error) {
	reqDemo := req.(*models.UpdateDemoRequest)
	logger := logs.WithContext(ctx)
	demoInter, err := d.demoDB.GetOneByID(ctx, objectID.String())
	demo := demoInter.(*models.Demo)
	if err != nil && !utils.IsErrNotFound(err) {
		logger.Errorf("Fail to get project by key: %v", err)
		return nil, utils.Gen("").Internal
	}
	if utils.IsErrNotFound(err) {
		return nil, utils.Demo("").NotFound
	}
	common.Sync(*reqDemo, demo)
	demo.ID = objectID
	demo.UpdaterID = updaterID
	if err := d.demoDB.Update(ctx, demo); err != nil {
		logger.Errorf("Fail to update demo to db: %s", err)
		return nil, utils.Gen("").Database
	}
	return demo, nil
}

func (d DemoHelper) GetOneByID(ctx context.Context, id uuid.UUID) (interface{}, error) {
	logger := logs.WithContext(ctx)
	demo, err := d.demoDB.GetOneByID(ctx, id.String())
	if err != nil && !utils.IsErrNotFound(err) {
		logger.Errorf("Fail to get demo by key: %v", err)
		return nil, utils.Gen("").Internal
	}
	if utils.IsErrNotFound(err) {
		return nil, utils.Demo("").NotFound
	}
	return demo, nil
}

func (d DemoHelper) GetAll(ctx beego.Controller) (interface{}, error) {
	demo, err := d.demoDB.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return demo, nil
}

func (d DemoHelper) Delete(ctx context.Context, id uuid.UUID) error {
	logger := logs.WithContext(ctx)
	productAttr, err := d.demoDB.GetOneByID(ctx, id.String())
	if err != nil && !utils.IsErrNotFound(err) {
		logger.Errorf("Fail to get record by key: %v", err)
		return utils.Gen("").Internal
	}
	if utils.IsErrNotFound(err) {
		return utils.Demo("").NotFound
	}
	if err := d.demoDB.Delete(ctx, productAttr); err != nil {
		return utils.Gen("").Database
	}
	return nil
}
