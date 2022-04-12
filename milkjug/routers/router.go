// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	apiControllers "milkshake/controllers/database"
	mobileControllers "milkshake/controllers/mobile"
	webControllers "milkshake/controllers/web"

	beego "github.com/beego/beego/v2/server/web"
)

/*
var FilterMobileUser = func(ctx *context.Context) {
	if ctx.Request.RequestURI != "/mobile/login" {
		if !jwt.VerifyApiCredentials(ctx) {
			ctx.ResponseWriter.Write([]byte("Unauthorized Request"))
			ctx.Output.SetStatus(401)
		}
	}
}
*/
func init() {
	webNS := beego.NewNamespace("/web",
		beego.NSRouter("/home", &webControllers.HomePageController{}, "get:Index"),
		beego.NSRouter("/login", &webControllers.WebLoginController{}, "get,post:Login"),
		beego.NSRouter("/loginQrCode", &webControllers.WebLoginController{}, "get,post:LoginWithQrCode"),
		beego.NSRouter("/admin", &webControllers.WebAdminController{}, "get:Index"),
		beego.NSRouter("/event", &webControllers.EventController{}, "get:Index"),
	)
	beego.AddNamespace(webNS)

	mobilesNS := beego.NewNamespace("/mobile",
		beego.NSRouter("/login", &mobileControllers.MobileLoginController{}, "get,post:Login"),
		beego.NSRouter("/user", &mobileControllers.MobileUserController{}),
		beego.NSRouter("/accounts", &mobileControllers.MobileAccountController{}),

		beego.NSNamespace("/events",
			beego.NSInclude(
				&mobileControllers.MobileEventsController{},
			),
		),
	)

	// beego.InsertFilter("/mobile/*", beego.BeforeRouter, FilterMobileUser)

	beego.AddNamespace(mobilesNS)

	ns := beego.NewNamespace("/database",

		beego.NSNamespace("/eventClock",
			beego.NSInclude(
				&apiControllers.EventClockController{},
			),
		),

		beego.NSNamespace("/eventHourType",
			beego.NSInclude(
				&apiControllers.EventHourTypeController{},
			),
		),

		beego.NSNamespace("/events",
			beego.NSInclude(
				&apiControllers.EventsController{},
			),
		),

		beego.NSNamespace("/payment",
			beego.NSInclude(
				&apiControllers.PaymentController{},
			),
		),

		beego.NSNamespace("/paymentMethod",
			beego.NSInclude(
				&apiControllers.PaymentMethodController{},
			),
		),

		beego.NSNamespace("/paymentRuns",
			beego.NSInclude(
				&apiControllers.PaymentRunsController{},
			),
		),

		beego.NSNamespace("/paymentType",
			beego.NSInclude(
				&apiControllers.PaymentTypeController{},
			),
		),

		beego.NSNamespace("/semester",
			beego.NSInclude(
				&apiControllers.SemesterController{},
			),
		),

		beego.NSNamespace("/studentDepartment",
			beego.NSInclude(
				&apiControllers.StudentDepartmentController{},
			),
		),

		beego.NSNamespace("/studentMajor",
			beego.NSInclude(
				&apiControllers.StudentMajorController{},
			),
		),

		beego.NSNamespace("/students",
			beego.NSInclude(
				&apiControllers.StudentsController{},
			),
		),

		beego.NSNamespace("/users",
			beego.NSInclude(
				&apiControllers.UsersController{},
			),
		),

		beego.NSNamespace("/userStatusType",
			beego.NSInclude(
				&apiControllers.UserStatusTypeController{},
			),
		),

		beego.NSNamespace("/userType",
			beego.NSInclude(
				&apiControllers.UserTypeController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
