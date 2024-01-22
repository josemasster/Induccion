package models

type Telefonos struct {
	Id             int        `orm:"column(id);pk;auto"`
	ContactoId     *Contactos `orm:"column(contacto_id);rel(fk)"`
	NumeroTelefono string     `orm:"column(numero_telefono)"`
}
