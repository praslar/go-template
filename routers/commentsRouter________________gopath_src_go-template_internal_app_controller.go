package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["go-template/internal/app/controller:GraphQLController"] = append(beego.GlobalControllerRouter["go-template/internal/app/controller:GraphQLController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-template/internal/app/controller:GraphQLController"] = append(beego.GlobalControllerRouter["go-template/internal/app/controller:GraphQLController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-template/internal/app/controller:GraphQLController"] = append(beego.GlobalControllerRouter["go-template/internal/app/controller:GraphQLController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-template/internal/app/controller:GraphQLController"] = append(beego.GlobalControllerRouter["go-template/internal/app/controller:GraphQLController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-template/internal/app/controller:GraphQLController"] = append(beego.GlobalControllerRouter["go-template/internal/app/controller:GraphQLController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-template/internal/app/controller:TempController"] = append(beego.GlobalControllerRouter["go-template/internal/app/controller:TempController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-template/internal/app/controller:TempController"] = append(beego.GlobalControllerRouter["go-template/internal/app/controller:TempController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-template/internal/app/controller:TempController"] = append(beego.GlobalControllerRouter["go-template/internal/app/controller:TempController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-template/internal/app/controller:TempController"] = append(beego.GlobalControllerRouter["go-template/internal/app/controller:TempController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-template/internal/app/controller:TempController"] = append(beego.GlobalControllerRouter["go-template/internal/app/controller:TempController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
