package main

import "fmt"

// TiendaRepuestos es una colección de repuestos de motos
type TiendaRepuestos struct {
	repuestos map[int]*Repuesto
}

//Repuesto se refiere a un repuesto de una moto en particular
type Repuesto struct {
	ID          int
	Nombre      string
	Marca       string
	Descripcion string
	NroPieza    int
	Precio      float64
	Stock       int
}

//NewTiendaRepuestos crea una colección de repuestos
func NewTiendaRepuestos() TiendaRepuestos {
	repuestos := make(map[int]*Repuesto)
	return TiendaRepuestos{
		repuestos,
	}
}

//Add agrega un repuesto a la tienda
func (t TiendaRepuestos) Add(r Repuesto) {
	t.repuestos[r.ID] = &r
}

//Print muestra por pantalla los elementos de un repuesto
func (t TiendaRepuestos) Print() {
	for _, v := range t.repuestos {
		fmt.Printf("id:[%v]\tnombre:'%v'\tdescripción:'%v'\tnroPieza:'%v'\tprecio:'%.2f'\tstock:'%d'\n", v.ID, v.Nombre, v.Descripcion, v.NroPieza, v.Precio, v.Stock)
	}
}

//FindByID devuelve el elemento correspondiente a la key dada si es que existe
func (t TiendaRepuestos) FindByID(ID int) *Repuesto {
	return t.repuestos[ID]
}

//Delete elimina de la tienda un repuesto dado
func (t TiendaRepuestos) Delete(ID int) {
	delete(t.repuestos, ID)
}

//Update actualiza los valores de un repuesto dado
func (t TiendaRepuestos) Update(r Repuesto) {
	t.repuestos[r.ID] = &r
}
