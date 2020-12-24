package handlers

import (
	"net/http"

	"github.com/astaxie/beego"

	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/app/types"
	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/pkg/logs"
	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/pkg/response"
)

type MetadataHelper interface {
	GetAll(ctx beego.Controller) (res []*types.MetaData, err error)
}

// MetaController operations for Meta
type MetaController struct {
	beego.Controller
	MetadataHelper MetadataHelper
}

// URLMapping ...
func (c *MetaController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

// GetAll ...
// @Title Get all data of metadata
// @Description get Meta
// @Param	type	query	string	false	"type"
// @Param	platform_key	query	string	false	"platform_key"
// @Success 200 {object} []models.MetaData
// @Failure 403
// @router / [get]
func (c *MetaController) GetAll() {
	logger := logs.WithContext(c.Ctx.Request.Context())
	variant, err := c.MetadataHelper.GetAll(c.Controller)
	if err != nil {
		logger.Errorf("Fail to get productAttrs due to %v", err)
		response.Error(c.Ctx.ResponseWriter, err, http.StatusInternalServerError)
		return
	}

	response.JSON(c.Ctx.ResponseWriter, http.StatusOK, response.BaseResponse{
		Data: variant,
	})
}
