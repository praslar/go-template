// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/app/databases"
	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/app/handlers"
	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/app/helpers"
	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/pkg/db/postgres"

	"github.com/astaxie/beego"
)

// Initialize make routes.init to be explicit, we have more control on this process
func Initialize(name, version string) {
	d, err := postgres.GetDatabase("default")
	if err != nil {
		return
	}
	variantDB := databases.NewVariantDatabase(d)
	metadataDB := databases.NewMetadataDatabase(d)

	variantHelper := helpers.NewVariantHelper(variantDB)
	metadataHelper := helpers.NewMetadataHelper(metadataDB)

	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/variant",
			beego.NSInclude(
				&handlers.VariantController{
					VariantHelper: variantHelper,
				},
			),
		),
		beego.NSNamespace("/metadata",
			beego.NSInclude(
				&handlers.MetaController{
					MetadataHelper: metadataHelper,
				},
			),
		),
		beego.NSRouter("/status", handlers.NewHealthController(name, version)),
	)
	beego.AddNamespace(ns)
}
