package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/app/types"
	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/pkg/common"
	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/pkg/logs"
	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/pkg/response"
	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/pkg/utils"

	"github.com/astaxie/beego"
	"github.com/google/uuid"
)

//  VariantController operations for variant

type VariantHelper interface {
	Create(ctx context.Context, req *types.CreateVariant, creatorID uuid.UUID) (*types.Variant, error)
	Update(ctx context.Context, id uuid.UUID, updaterID uuid.UUID, req *types.UpdateVariant) (*types.Variant, error)
	GetOneByID(ctx context.Context, id uuid.UUID) (*types.Variant, error)
	GetAll(ctx beego.Controller) ([]*types.Variant, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type VariantController struct {
	beego.Controller
	VariantHelper VariantHelper
}

// URLMapping ...
func (c *VariantController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Put", c.Put)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create variant
// @Param	Authorization	header	string	true	"authen key"
// @Param	x-user-id	header	string	true	"user id call request"
// @Param	body		body 	models.Variant	true		"body for variant content"
// @Success 200 {object} models.Variant
// @Failure 403  forbidden
// @Failure 500  error server
// @Failure 401  Unauthorized
// @router / [post]
func (c *VariantController) Post() {
	logger := logs.WithContext(c.Ctx.Request.Context())
	creatorID, err := common.CurrentUser(c.Controller)
	if err != nil {
		logger.Errorf("Invalid x-user-id, %s", err)
		response.Error(c.Ctx.ResponseWriter, utils.Gen("").Unauthorized, http.StatusUnauthorized)
		return
	}
	var req types.CreateVariant
	if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&req); err != nil {
		logger.Errorf("Fail to parse create variant request, %s", err)
		response.Error(c.Ctx.ResponseWriter, err, http.StatusBadRequest)
		return
	}

	variant, err := c.VariantHelper.Create(c.Controller.Ctx.Request.Context(), &req, creatorID)
	if err != nil {
		logger.Errorf("Fail to create variant attribute due to %v", err)
		response.Error(c.Ctx.ResponseWriter, err, http.StatusInternalServerError)
		return
	}

	response.JSON(c.Ctx.ResponseWriter, http.StatusOK, response.BaseResponse{
		Data: variant,
	})
}

// Put ...
// @Title Put
// @Description update the variant
// @Param	Authorization	header	string	true	"authen key"
// @Param	x-user-id	header	string	true	"user id call request"
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Variant	true		"body for station content"
// @Success 200 {object} models.Variant
// @Failure 403  forbidden
// @Failure 500  error server
// @Failure 401  Unauthorized
// @router /:id [put]
func (c *VariantController) Put() {

	logger := logs.WithContext(c.Ctx.Request.Context())
	updaterID, err := common.CurrentUser(c.Controller)
	if err != nil {
		logger.Errorf("Invalid x-user-id, %s", err)
		response.Error(c.Ctx.ResponseWriter, utils.Gen("").Unauthorized, http.StatusUnauthorized)
		return
	}

	// Get variant id from param
	variantID, err := common.With(c.Controller).On("path").Uuid("id")
	if err != nil {
		logger.Errorf("Invalid variant , %s", err)
		response.Error(c.Ctx.ResponseWriter, utils.Gen("Invalid variant  ID").BadRequest, http.StatusBadRequest)
		return
	}

	var req types.UpdateVariant
	if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&req); err != nil {
		logger.Errorf("Fail to parse create variant request, %s", err)
		response.Error(c.Ctx.ResponseWriter, err, http.StatusBadRequest)
		return
	}

	updateVariant, err := c.VariantHelper.Update(c.Controller.Ctx.Request.Context(), variantID, updaterID, &req)
	if err != nil {
		logger.Errorf("Fail to update variant due to %v", err)
		response.Error(c.Ctx.ResponseWriter, err, http.StatusInternalServerError)
		return
	}

	response.JSON(c.Ctx.ResponseWriter, http.StatusOK, response.BaseResponse{
		Data: updateVariant,
	})
}

// GetOne ...
// @Title GetOne
// @Description get variant by id
// @Param	Authorization	header	string	true	"authen key"
// @Param	id		path 	string	true		"id variant"
// @Success 200 {object} models.Variant
// @Failure 403  forbidden
// @Failure 500  error server
// @Failure 401  Unauthorized
// @router /:id [get]
func (c *VariantController) GetOne() {
	logger := logs.WithContext(c.Ctx.Request.Context())
	variantID, err := common.With(c.Controller).On("path").Uuid("id")
	if err != nil {
		logger.Errorf("Invalid variant ID, %s", err)
		response.Error(c.Ctx.ResponseWriter, utils.Gen("Invalid variant ID").BadRequest, http.StatusBadRequest)
		return
	}

	variant, err := c.VariantHelper.GetOneByID(c.Controller.Ctx.Request.Context(), variantID)
	if err != nil {
		logger.Errorf("Fail to get variant by ID due to %v", err)
		response.Error(c.Ctx.ResponseWriter, err, http.StatusInternalServerError)
		return
	}

	response.JSON(c.Ctx.ResponseWriter, http.StatusOK, response.BaseResponse{
		Data: variant,
	})
}

// GetAll ...
// @Title GetAll
// @Description get variants
// @Param	Authorization	header	string	true	"authen key"
// @Param	size	query	string	true	"size per page"
// @Param	page	query	string	true	"page number (> 0)"
// @Param	platform_key	query	string	true	"platform code"
// @Param	search	query	string	true	"search"
// @Param	admin_name	query	string	true	"admin name of variant"
// @Success 200 {object} []models.Variant
// @Failure 403  forbidden
// @Failure 500  error server
// @Failure 401  Unauthorized
// @router / [get]
func (c *VariantController) GetAll() {
	logger := logs.WithContext(c.Ctx.Request.Context())
	variant, err := c.VariantHelper.GetAll(c.Controller)
	if err != nil {
		logger.Errorf("Fail to get productAttrs due to %v", err)
		response.Error(c.Ctx.ResponseWriter, err, http.StatusInternalServerError)
		return
	}

	response.JSON(c.Ctx.ResponseWriter, http.StatusOK, response.BaseResponse{
		Data: variant,
	})
}

// Delete ...
// @Title Delete
// @Description delete the variant
// @Param	Authorization	header	string	true	"authenticate key"
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete product attribute id
// @Failure 403  forbidden
// @Failure 401  Unauthorized
// @router /:id [delete]
func (c *VariantController) Delete() {
	// Get x-user-id from header
	logger := logs.WithContext(c.Ctx.Request.Context())
	_, err := common.CurrentUser(c.Controller)
	if err != nil {
		logger.Errorf("Invalid x-user-id, %s", err)
		response.Error(c.Ctx.ResponseWriter, utils.Gen("").Unauthorized, http.StatusUnauthorized)
		return
	}

	// Get product attribute id from param
	variantID, err := common.With(c.Controller).On("path").Uuid("id")
	if err != nil {
		logger.Errorf("Invalid product attribute ID, %s", err)
		response.Error(c.Ctx.ResponseWriter, utils.Gen("Invalid product attribute ID").BadRequest, http.StatusBadRequest)
		return
	}

	if err := c.VariantHelper.Delete(c.Controller.Ctx.Request.Context(), variantID); err != nil {
		logger.Errorf("Fail to delete variant due to %v", err)
		response.Error(c.Ctx.ResponseWriter, err, http.StatusInternalServerError)
		return
	}

	response.JSON(c.Ctx.ResponseWriter, http.StatusOK, response.BaseResponse{
		Data: variantID,
	})
}
