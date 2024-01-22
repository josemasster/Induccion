package controllers

import (
	"github.com/beego/beego/api-mid-udistrital/helpers"

	"github.com/astaxie/beego"
)

// ContactosController operations for Contactos
type Contactos struct {
	beego.Controller
}

// URLMapping ...
func (c *Contactos) URLMapping() {

}

// GetAll ...
// @Title GetAll
// @Description get obetener contactos
// @Success 200 {object} []models.ResolucionesParametros
// @Failure 400 bad request
// @Failure 500 Internal server error
// @router / [get]
func (c *Contactos) GetAll() {
	defer helpers.ErrorController(c.Controller, "Contactos")

	if v, err := helpers.ListarContactos(); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Succes": true, "Status": 200, "Message": "Listado consultado con exito", "Data": v}

	} else {
		panic(err)
	}
	c.ServeJSON()

}

// Post ...
// @Title Post
// @Description Crear obetener contactos
// @Success 200 {object} []models.ResolucionesParametros
// @Failure 400 bad request
// @Failure 500 Internal server error
// @router / [post]

/*func (c *Contactos) Post() {
	defer helpers.ErrorController(c.Controller, "Contactos")

	if v, err := helpers.CrearContacto(); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Succes": true, "Status": 200, "Message": "Listado consultado con exito", "Data": v}

	} else {
		panic(err)
	}
	c.ServeJSON()

}*/
