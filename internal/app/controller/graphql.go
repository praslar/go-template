package controller

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"go-template/internal/app/models"
	"go-template/internal/pkg/logs"
	"go-template/internal/pkg/response"
	"net/http"
)

type GraphQLController struct {
	beego.Controller
	GraphQLHelper GraphQLHelper
}

// URLMapping ...
func (c *GraphQLController) URLMapping() {
	c.Mapping("Login", c.Login)
	c.Mapping("Put", c.Put)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Delete", c.Delete)
}

type GraphQLHelper interface {
	models.BaseHelper
	// some func of Temp
	Login(ctx context.Context, req interface{}) (interface{}, error)
}

var (
	// compile time type assertion
	_ models.BaseController = &GraphQLController{}
)

// Post ...
// @Title Create
// @Description create variant
// @Param	Authorization	header	string	true	"authen key"
// @Param	x-user-id	header	string	true	"user id call request"
// @Param	body		body 	models.Temp	true		"body for variant content"
// @Success 200 {object} models.Temp
// @Failure 403  forbidden
// @Failure 500  error server
// @Failure 401  Unauthorized
// @router /login [post]
func (t *GraphQLController) Login() {
	logger := logs.WithContext(t.Ctx.Request.Context())
	req := models.ActionPayload{}
	if err := json.NewDecoder(t.Ctx.Request.Body).Decode(&req); err != nil {
		logger.Errorf("Fail to login %s", err)
		response.Error(t.Ctx.ResponseWriter, err, http.StatusBadRequest)
		return
	}
	res, err := t.GraphQLHelper.Login(t.Ctx.Request.Context(), req)
	if err != nil {
		logger.Errorf("Fail to login %s", err)
		response.JSON(t.Ctx.ResponseWriter, http.StatusBadRequest, models.GraphQLError{Message: err.Error()})
	} else {
		response.JSON(t.Ctx.ResponseWriter, http.StatusOK, res)
	}
}

// Put ...
// @Title Put
// @Description update the variant
// @Param	Authorization	header	string	true	"authen key"
// @Param	x-user-id	header	string	true	"user id call request"
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Temp	true		"body for station content"
// @Success 200 {object} models.Temp
// @Failure 403  forbidden
// @Failure 500  error server
// @Failure 401  Unauthorized
// @router /:id [put]
func (c *GraphQLController) Put() {
	//logger := logs.WithContext(c.Ctx.Request.Context())
	//updaterID, err := common.CurrentUser(c.Controller)
	//if err != nil {
	//	logger.Errorf("Invalid x-user-id, %s", err)
	//	response.Error(c.Ctx.ResponseWriter, utils.Gen("").Unauthorized, http.StatusUnauthorized)
	//	return
	//}
	//
	//// Get variant id from param
	//variantID, err := common.With(c.Controller).On("path").Uuid("id")
	//if err != nil {
	//	logger.Errorf("Invalid variant , %s", err)
	//	response.Error(c.Ctx.ResponseWriter, utils.Gen("Invalid variant  ID").BadRequest, http.StatusBadRequest)
	//	return
	//}
	//
	//var req models.UpdateTemp
	//if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&req); err != nil {
	//	logger.Errorf("Fail to parse create variant request, %s", err)
	//	response.Error(c.Ctx.ResponseWriter, err, http.StatusBadRequest)
	//	return
	//}
	//
	//updateVariant, err := c.VariantHelper.Update(c.Controller.Ctx.Request.Context(), variantID, updaterID, &req)
	//if err != nil {
	//	logger.Errorf("Fail to update variant due to %v", err)
	//	response.Error(c.Ctx.ResponseWriter, err, http.StatusInternalServerError)
	//	return
	//}
	//
	//response.JSON(c.Ctx.ResponseWriter, http.StatusOK, response.BaseResponse{
	//	Data: updateVariant,
	//})
}

// GetOne ...
// @Title GetOne
// @Description get variant by id
// @Param	Authorization	header	string	true	"authen key"
// @Param	id		path 	string	true		"id variant"
// @Success 200 {object} models.Temp
// @Failure 403  forbidden
// @Failure 500  error server
// @Failure 401  Unauthorized
// @router /:id [get]
func (c *GraphQLController) GetOne() {
	//logger := logs.WithContext(c.Ctx.Request.Context())
	//variantID, err := common.With(c.Controller).On("path").Uuid("id")
	//if err != nil {
	//	logger.Errorf("Invalid variant ID, %s", err)
	//	response.Error(c.Ctx.ResponseWriter, utils.Gen("Invalid variant ID").BadRequest, http.StatusBadRequest)
	//	return
	//}
	//
	//variant, err := c.VariantHelper.GetOneByID(c.Controller.Ctx.Request.Context(), variantID)
	//if err != nil {
	//	logger.Errorf("Fail to get variant by ID due to %v", err)
	//	response.Error(c.Ctx.ResponseWriter, err, http.StatusInternalServerError)
	//	return
	//}
	//
	//response.JSON(c.Ctx.ResponseWriter, http.StatusOK, response.BaseResponse{
	//	Data: variant,
	//})
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
// @Success 200 {object} []models.Temp
// @Failure 403  forbidden
// @Failure 500  error server
// @Failure 401  Unauthorized
// @router / [get]
func (c *GraphQLController) GetAll() {
	//logger := logs.WithContext(c.Ctx.Request.Context())
	//variant, err := c.VariantHelper.GetAll(c.Controller)
	//if err != nil {
	//	logger.Errorf("Fail to get productAttrs due to %v", err)
	//	response.Error(c.Ctx.ResponseWriter, err, http.StatusInternalServerError)
	//	return
	//}
	//
	//response.JSON(c.Ctx.ResponseWriter, http.StatusOK, response.BaseResponse{
	//	Data: variant,
	//})
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
func (c *GraphQLController) Delete() {
	// Get x-user-id from header
	//logger := logs.WithContext(c.Ctx.Request.Context())
	//_, err := common.CurrentUser(c.Controller)
	//if err != nil {
	//	logger.Errorf("Invalid x-user-id, %s", err)
	//	response.Error(c.Ctx.ResponseWriter, utils.Gen("").Unauthorized, http.StatusUnauthorized)
	//	return
	//}
	//
	//// Get product attribute id from param
	//variantID, err := common.With(c.Controller).On("path").Uuid("id")
	//if err != nil {
	//	logger.Errorf("Invalid product attribute ID, %s", err)
	//	response.Error(c.Ctx.ResponseWriter, utils.Gen("Invalid product attribute ID").BadRequest, http.StatusBadRequest)
	//	return
	//}
	//
	//if err := c.VariantHelper.Delete(c.Controller.Ctx.Request.Context(), variantID); err != nil {
	//	logger.Errorf("Fail to delete variant due to %v", err)
	//	response.Error(c.Ctx.ResponseWriter, err, http.StatusInternalServerError)
	//	return
	//}
	//
	//response.JSON(c.Ctx.ResponseWriter, http.StatusOK, response.BaseResponse{
	//	Data: variantID,
	//})
}
