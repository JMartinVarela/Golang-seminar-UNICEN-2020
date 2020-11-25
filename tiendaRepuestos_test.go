package main

import "testing"

func TestTiendaRepuestosAdd(t *testing.T) {
	tienda := NewTiendaRepuestos()
	r0 := tienda.FindByID(0)
	if r0 != nil {
		t.Error("El repuesto con ID=0 ya existe en la tienda!!")
	}

	tienda.Add(Repuesto{0, "Juego espejos", "Honda", "Juego de espejos retrovisores", 12843, 680.00, 20})
	r0 = tienda.FindByID(0)
	if r0 == nil {
		t.Error("El repuesto con ID=0 no fue agregado a la tienda!!")
	}

	if r0.Nombre != "Juego espejos" {
		t.Error("El repuesto con ID=0 no tiene el nombre correcto!!")
	}

	if r0.Marca != "Honda" {
		t.Error("El repuesto con ID=0 no tiene la marca correcta!!")
	}

	if r0.Descripcion != "Juego de espejos retrovisores" {
		t.Error("El repuesto con ID=0 no tiene la descripcion correcta!!")
	}

	if r0.NroPieza != 12843 {
		t.Error("El repuesto con ID=0 no tiene el número de pieza correcto!!")
	}

	if r0.Precio != 680.00 {
		t.Error("El repuesto con ID=0 no tiene el precio correcto!!")
	}

	if r0.Stock != 20 {
		t.Error("El repuesto con ID=0 no tiene la cantidad de stock correcta!!")
	}
}

func TestTiendaRepuestosRead(t *testing.T) {
	tienda := NewTiendaRepuestos()
	tienda.Add(Repuesto{0, "Juego espejos", "Honda", "Juego de espejos retrovisores", 12843, 680.00, 20})
	r0 := tienda.FindByID(0)
	if r0 == nil {
		t.Error("El repuesto con ID=0 no se encontró en la tienda!!")
	}
}

func TestTiendaRepuestosDelete(t *testing.T) {
	tienda := NewTiendaRepuestos()
	tienda.Add(Repuesto{0, "Juego espejos", "Honda", "Juego de espejos retrovisores", 12843, 680.00, 20})
	tienda.Delete(0)
	r0 := tienda.FindByID(0)
	if r0 != nil {
		t.Error("El repuesto con ID=0 no se eliminó de la tienda!!")
	}
}

func TestTiendaRepuestosUpdate(t *testing.T) {
	tienda := NewTiendaRepuestos()
	tienda.Add(Repuesto{0, "Juego espejos", "Honda", "Juego de espejos retrovisores", 12843, 680.00, 20})
	tienda.Update(Repuesto{0, "Juego espejos", "Honda", "Juego de espejos retrovisores", 12843, 680.00, 18})
	r0 := tienda.FindByID(0)
	if r0.Stock != 18 {
		t.Error("El repuesto con ID=0 no se actualizó correctamente!!")
	}
}
