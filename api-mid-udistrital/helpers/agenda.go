package helpers

import (
	"fmt"

	"github.com/beego/beego/api-mid-udistrital/models"

	"github.com/beego/beego/logs"
)

func ListarContactos() (agenda models.Agenda, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "ListadeResoluciones", "err": err, "status": "500"}
			panic(outputError)
		}

	}()
	var datos1 []models.Contactos
	url1 := "contactos"

	if err := GetReqNew("UrlCrudAgenda", url1, &datos1); err != nil {
		logs.Error(err)
		panic(err.Error)
	}
	var datos2 []models.Correoselectronicos
	url2 := "correoselectronicos"

	if err := GetReqNew("UrlCrudAgenda", url2, &datos2); err != nil {
		logs.Error(err)
		panic(err.Error)
	}

	var datos3 []models.Telefonos
	url3 := "telefonos"
	if err := GetReqNew("UrlCrudAgenda", url3, &datos3); err != nil {
		logs.Error(err)
		panic(err.Error)
	}

	fmt.Println(datos3)

	agenda.DatosContactos = datos1
	agenda.Correoselectronicos = datos2
	agenda.Telefono = datos3

	return agenda, outputError
}

/*func CrearContacto() (agenda models.Agenda, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "ListadeResoluciones", "err": err, "status": "500"}
			panic(outputError)
		}

	}()
	var datos []models.Contactos
	url := "contactos"

	if err := Post("UrlCrudAgenda", url, &datos); err != nil {
		logs.Error(err)
		panic(err.Error)
	}
	fmt.Println("Datos", datos)

	agenda.DatosContactos = datos

	return agenda, outputError
}*/
