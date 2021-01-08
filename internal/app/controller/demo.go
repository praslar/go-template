package controller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"go-template/internal/app/models"
	"go-template/internal/pkg/common"
	"go-template/internal/pkg/logs"
	"go-template/internal/pkg/response"
	"go-template/internal/pkg/utils"
	"net/http"
)

type DemoController struct {
	beego.Controller
	DemoHelper DemoHelper
}

// URLMapping ...
func (c *DemoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Put", c.Put)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Delete", c.Delete)
}

type DemoHelper interface {
	models.BaseHelper
	// some func of Demo
}

var (
	// compile time type assertion
	_ models.BaseController = &DemoController{}
)

// Post ...
// @Title Create
// @Description create demo
// @Param	x-user-id	header	string	true	"user id call request"
// @Param	body		body 	models.CreateDemoRequest	true		"body for demo content"
// @Success 200 {object} models.Demo
// @Failure 403  forbidden
// @Failure 500  error server
// @Failure 401  Unauthorized
// @router / [post]
func (c *DemoController) Post() {
	logger := logs.WithContext(c.Ctx.Request.Context())
	creatorID, err := common.CurrentUser(c.Controller)
	if err != nil {
		logger.Errorf("Invalid x-user-id, %s", err)
		response.Error(c.Ctx.ResponseWriter, utils.Gen("").Unauthorized, http.StatusUnauthorized)
		return
	}
	var req models.CreateDemoRequest
	if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&req); err != nil {
		logger.Errorf("Fail to parse create demo request, %s", err)
		response.Error(c.Ctx.ResponseWriter, err, http.StatusBadRequest)
		return
	}
	demo, err := c.DemoHelper.Create(c.Controller.Ctx.Request.Context(), &req, creatorID)
	if err != nil {
		logger.Errorf("Fail to create demo due to %v", err)
		response.Error(c.Ctx.ResponseWriter, err, http.StatusInternalServerError)
		return
	}
	
	response.JSON(c.Ctx.ResponseWriter, http.StatusOK, response.BaseResponse{
		Data: demo,
	})
}

// Put ...
// @Title Put
// @Description update the demo
// @Param	x-user-id	header	string	true	"user id call request"
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.UpdateDemoRequest	true		"body for station content"
// @Success 200 {object} models.Demo
// @Failure 403  forbidden
// @Failure 500  error server
// @Failure 401  Unauthorized
// @router /:id [put]
func (c *DemoController) Put() {
	
	logger := logs.WithContext(c.Ctx.Request.Context())
	updaterID, err := common.CurrentUser(c.Controller)
	if err != nil {
		logger.Errorf("Invalid x-user-id, %s", err)
		response.Error(c.Ctx.ResponseWriter, utils.Gen("").Unauthorized, http.StatusUnauthorized)
		return
	}
	
	// Get demo id from param
	demoID, err := common.With(c.Controller).On("path").Uuid("id")
	if err != nil {
		logger.Errorf("Invalid demo , %s", err)
		response.Error(c.Ctx.ResponseWriter, utils.Gen("Invalid demo  ID").BadRequest, http.StatusBadRequest)
		return
	}
	
	var req models.UpdateDemoRequest
	if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&req); err != nil {
		logger.Errorf("Fail to parse create demo request, %s", err)
		response.Error(c.Ctx.ResponseWriter, err, http.StatusBadRequest)
		return
	}
	updateDemo, err := c.DemoHelper.Update(c.Controller.Ctx.Request.Context(), &req, demoID, updaterID)
	if err != nil {
		logger.Errorf("Fail to update demo due to %v", err)
		response.Error(c.Ctx.ResponseWriter, err, http.StatusInternalServerError)
		return
	}
	
	response.JSON(c.Ctx.ResponseWriter, http.StatusOK, response.BaseResponse{
		Data: updateDemo,
	})
}

// GetOne ...
// @Title GetOne
// @Description get demo by id
// @Param	id		path 	string	true		"id demo"
// @Success 200 {object} models.Demo
// @Failure 403  forbidden
// @Failure 500  error server
// @Failure 401  Unauthorized
// @router /:id [get]
func (c *DemoController) GetOne() {
	logger := logs.WithContext(c.Ctx.Request.Context())
	demoID, err := common.With(c.Controller).On("path").Uuid("id")
	if err != nil {
		logger.Errorf("Invalid demo ID, %s", err)
		response.Error(c.Ctx.ResponseWriter, utils.Gen("Invalid demo ID").BadRequest, http.StatusBadRequest)
		return
	}
	
	demo, err := c.DemoHelper.GetOneByID(c.Controller.Ctx.Request.Context(), demoID)
	if err != nil {
		logger.Errorf("Fail to get demo by ID due to %v", err)
		response.Error(c.Ctx.ResponseWriter, err, http.StatusInternalServerError)
		return
	}
	
	response.JSON(c.Ctx.ResponseWriter, http.StatusOK, response.BaseResponse{
		Data: demo,
	})
}

// GetAll ...
// @Title GetAll
// @Description get demos
// @Param	size	query	string	true	"size per page"
// @Param	page	query	string	true	"page number (> 0)"
// @Param	search	query	string	false	"search"
// @Param	name	query	string	false	"name of demo"
// @Param	age	query	string	false	"age of demo"
// @Success 200 {object} []models.Demo
// @Failure 403  forbidden
// @Failure 500  error server
// @Failure 401  Unauthorized
// @router / [get]
func (c *DemoController) GetAll() {
	logger := logs.WithContext(c.Ctx.Request.Context())
	demo, err := c.DemoHelper.GetAll(c.Controller)
	if err != nil {
		logger.Errorf("Fail to get demo due to %v", err)
		response.Error(c.Ctx.ResponseWriter, err, http.StatusInternalServerError)
		return
	}
	
	response.JSON(c.Ctx.ResponseWriter, http.StatusOK, response.BaseResponse{
		Data: demo,
	})
}

// Delete ...
// @Title Delete
// @Description delete the demo
// @Param	x-user-id	header	string	true	"user id call request"
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete demo id
// @Failure 403  forbidden
// @Failure 404  Not Found
// @Failure 401  Unauthorized
// @router /:id [delete]
func (c *DemoController) Delete() {
	// Get x-user-id from header
	logger := logs.WithContext(c.Ctx.Request.Context())
	_, err := common.CurrentUser(c.Controller)
	if err != nil {
		logger.Errorf("Invalid x-user-id, %s", err)
		response.Error(c.Ctx.ResponseWriter, utils.Gen("").Unauthorized, http.StatusUnauthorized)
		return
	}
	
	// Get demo id from param
	demoID, err := common.With(c.Controller).On("path").Uuid("id")
	if err != nil {
		logger.Errorf("Invalid product attribute ID, %s", err)
		response.Error(c.Ctx.ResponseWriter, utils.Gen("Invalid demo ID").BadRequest, http.StatusBadRequest)
		return
	}
	
	if err := c.DemoHelper.Delete(c.Controller.Ctx.Request.Context(), demoID); err != nil {
		logger.Errorf("Fail to delete demo due to %v", err)
		response.Error(c.Ctx.ResponseWriter, err, http.StatusInternalServerError)
		return
	}
	
	response.JSON(c.Ctx.ResponseWriter, http.StatusOK, response.BaseResponse{
		Data: demoID,
	})
}
