package models

type Correoselectronicos struct {
	Id              int        `orm:"column(id);pk;auto"`
	ContactoId      *Contactos `orm:"column(contacto_id);rel(fk)"`
	DireccionCorreo string     `orm:"column(direccion_correo)"`
}
