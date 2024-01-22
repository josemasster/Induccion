package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:ContactosController"] = append(beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:ContactosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:ContactosController"] = append(beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:ContactosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:ContactosController"] = append(beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:ContactosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:ContactosController"] = append(beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:ContactosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:ContactosController"] = append(beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:ContactosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:CorreoselectronicosController"] = append(beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:CorreoselectronicosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:CorreoselectronicosController"] = append(beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:CorreoselectronicosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:CorreoselectronicosController"] = append(beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:CorreoselectronicosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:CorreoselectronicosController"] = append(beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:CorreoselectronicosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:CorreoselectronicosController"] = append(beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:CorreoselectronicosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:TelefonosController"] = append(beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:TelefonosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:TelefonosController"] = append(beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:TelefonosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:TelefonosController"] = append(beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:TelefonosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:TelefonosController"] = append(beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:TelefonosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:TelefonosController"] = append(beego.GlobalControllerRouter["github.com/beego/beego/Agenda_api/controllers:TelefonosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
