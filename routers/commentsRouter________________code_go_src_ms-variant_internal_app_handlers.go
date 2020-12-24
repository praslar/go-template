package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["gitlab.com/jamalex-ltd/r-d-department/go-template/internal/app/handlers:MetaController"] = append(beego.GlobalControllerRouter["gitlab.com/jamalex-ltd/r-d-department/go-template/internal/app/handlers:MetaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/jamalex-ltd/r-d-department/go-template/internal/app/handlers:VariantController"] = append(beego.GlobalControllerRouter["gitlab.com/jamalex-ltd/r-d-department/go-template/internal/app/handlers:VariantController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/jamalex-ltd/r-d-department/go-template/internal/app/handlers:VariantController"] = append(beego.GlobalControllerRouter["gitlab.com/jamalex-ltd/r-d-department/go-template/internal/app/handlers:VariantController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/jamalex-ltd/r-d-department/go-template/internal/app/handlers:VariantController"] = append(beego.GlobalControllerRouter["gitlab.com/jamalex-ltd/r-d-department/go-template/internal/app/handlers:VariantController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/jamalex-ltd/r-d-department/go-template/internal/app/handlers:VariantController"] = append(beego.GlobalControllerRouter["gitlab.com/jamalex-ltd/r-d-department/go-template/internal/app/handlers:VariantController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/jamalex-ltd/r-d-department/go-template/internal/app/handlers:VariantController"] = append(beego.GlobalControllerRouter["gitlab.com/jamalex-ltd/r-d-department/go-template/internal/app/handlers:VariantController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
