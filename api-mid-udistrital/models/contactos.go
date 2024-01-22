package models

type Contactos struct {
	Id              int    `orm:"column(id);pk;auto"`
	NombreCompleto  string `orm:"column(nombre_completo)"`
	NumeroDocumento string `orm:"column(numero_documento)"`
	Direccion       string `orm:"column(direccion);null"`
	FechaCreacion   string `orm:"column(fecha_creacion);type(timestamp without time zone);null;auto_now_add"`
}
