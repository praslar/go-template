// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"go-template/internal/app/databases"
	"go-template/internal/app/controller"
	"go-template/internal/app/helpers"
	"go-template/internal/pkg/db/postgres"

	"github.com/astaxie/beego"
)

// Initialize make routes.init to be explicit, we have more control on this process
func Initialize(name, version string) {
	d, err := postgres.GetDatabase("default")
	if err != nil {
		return
	}
	tempDB := databases.NewTempDatabase(d)
	tempHelper := helpers.NewTempHelper(*tempDB)

	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/temp",
			beego.NSInclude(
				&controller.TempController{
					TempHelper: tempHelper,
				},
			),
		),
		beego.NSRouter("/status", controller.NewHealthController(name, version)),
	)
	beego.AddNamespace(ns)
}
