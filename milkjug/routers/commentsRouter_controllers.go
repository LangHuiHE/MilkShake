package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	// mobile
	beego.GlobalControllerRouter["milkshake/controllers/mobile:MobileEventsController"] = append(beego.GlobalControllerRouter["milkshake/controllers/mobile:MobileEventsController"],
		beego.ControllerComments{
			Method:           "GetRecent",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/mobile:MobileEventsController"] = append(beego.GlobalControllerRouter["milkshake/controllers/mobile:MobileEventsController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/mobile:MobileEventsController"] = append(beego.GlobalControllerRouter["milkshake/controllers/mobile:MobileEventsController"],
		beego.ControllerComments{
			Method:           "ClockIn",
			Router:           "/:id",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	// DateBase
	beego.GlobalControllerRouter["milkshake/controllers/database:EventClockController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:EventClockController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:EventClockController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:EventClockController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:EventClockController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:EventClockController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:EventClockController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:EventClockController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:EventClockController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:EventClockController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:EventHourTypeController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:EventHourTypeController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:EventHourTypeController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:EventHourTypeController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:EventHourTypeController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:EventHourTypeController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:EventHourTypeController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:EventHourTypeController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:EventHourTypeController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:EventHourTypeController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:EventsController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:EventsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:EventsController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:EventsController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:EventsController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:EventsController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:EventsController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:EventsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:EventsController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:EventsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:PaymentController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:PaymentController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:PaymentController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:PaymentController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:PaymentController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:PaymentController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:PaymentController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:PaymentController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:PaymentController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:PaymentController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:PaymentMethodController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:PaymentMethodController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:PaymentMethodController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:PaymentMethodController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:PaymentMethodController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:PaymentMethodController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:PaymentMethodController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:PaymentMethodController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:PaymentMethodController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:PaymentMethodController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:PaymentRunsController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:PaymentRunsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:PaymentRunsController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:PaymentRunsController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:PaymentRunsController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:PaymentRunsController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:PaymentRunsController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:PaymentRunsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:PaymentRunsController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:PaymentRunsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:PaymentTypeController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:PaymentTypeController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:PaymentTypeController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:PaymentTypeController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:PaymentTypeController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:PaymentTypeController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:PaymentTypeController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:PaymentTypeController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:PaymentTypeController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:PaymentTypeController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:SemesterController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:SemesterController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:SemesterController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:SemesterController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:SemesterController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:SemesterController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:SemesterController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:SemesterController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:SemesterController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:SemesterController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:StudentDepartmentController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:StudentDepartmentController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:StudentDepartmentController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:StudentDepartmentController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:StudentDepartmentController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:StudentDepartmentController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:StudentDepartmentController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:StudentDepartmentController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:StudentDepartmentController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:StudentDepartmentController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:StudentMajorController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:StudentMajorController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:StudentMajorController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:StudentMajorController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:StudentMajorController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:StudentMajorController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:StudentMajorController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:StudentMajorController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:StudentMajorController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:StudentMajorController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:StudentsController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:StudentsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:StudentsController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:StudentsController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:StudentsController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:StudentsController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:StudentsController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:StudentsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:StudentsController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:StudentsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:UserStatusTypeController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:UserStatusTypeController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:UserStatusTypeController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:UserStatusTypeController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:UserStatusTypeController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:UserStatusTypeController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:UserStatusTypeController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:UserStatusTypeController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:UserStatusTypeController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:UserStatusTypeController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:UserTypeController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:UserTypeController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:UserTypeController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:UserTypeController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:UserTypeController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:UserTypeController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:UserTypeController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:UserTypeController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:UserTypeController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:UserTypeController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:UsersController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:UsersController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:UsersController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:UsersController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:UsersController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:UsersController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:UsersController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:UsersController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["milkshake/controllers/database:UsersController"] = append(beego.GlobalControllerRouter["milkshake/controllers/database:UsersController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
