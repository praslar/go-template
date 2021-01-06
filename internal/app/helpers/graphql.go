package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/google/uuid"
	"github.com/sendgrid/rest"
	"go-template/internal/app/models"
	"go-template/internal/pkg/utils"
	"net/http"
)

type graphqlDatabase interface {
	models.BaseDatabase
	// some func of Temp
}

var (
	// compile time type assertion
	_ models.BaseHelper = &GraphQLHelper{}
)

type GraphQLHelper struct {
	tempDB graphqlDatabase
}

func NewGraphQlHelper(db graphqlDatabase) *GraphQLHelper {
	return &GraphQLHelper{
		db,
	}
}

func (g *GraphQLHelper) LoginHandler(req models.LoginRequest) (res models.LoginResponse, err error) {
	tmpReq, err := json.Marshal(req)
	if err != nil {
		return res, err
	}
	request := rest.Request{
		Method:  rest.Post,
		BaseURL: utils.URL_LOGIN_API,
		Body:    tmpReq,
	}
	response, err := rest.Send(request)
	if err != nil {
		return res, err
	}
	if response.StatusCode != http.StatusOK {
		return res, fmt.Errorf("Error when call Login API %s", response.Body)
	}
	var esBody = models.ResponseDefault{}
	err = json.Unmarshal([]byte(response.Body), &esBody)
	if err != nil {
		return res, err
	}
	return esBody.Data.(models.LoginResponse), err
}

func (g *GraphQLHelper) Login(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(models.ActionPayload)
	res, err := g.LoginHandler(r.Input)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (g *GraphQLHelper) Create(ctx context.Context, req interface{}) (interface{}, error) {
	panic("implement me")
}

func (g *GraphQLHelper) Update(ctx context.Context, req interface{}) (interface{}, error) {
	panic("implement me")
}

func (g *GraphQLHelper) GetOneByID(ctx context.Context, id uuid.UUID) (interface{}, error) {
	panic("implement me")
}

func (g *GraphQLHelper) GetAll(ctx beego.Controller) ([]interface{}, error) {
	panic("implement me")
}

func (g *GraphQLHelper) Delete(ctx context.Context, id uuid.UUID) error {
	panic("implement me")
}
