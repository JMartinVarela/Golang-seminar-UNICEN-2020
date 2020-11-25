package main

import "fmt"

func main() {
	t := NewTiendaRepuestos()
	t.Add(Repuesto{0, "Juego espejos", "Honda", "Juego de espejos retrovisores", 12843, 680.00, 20})
	t.Add(Repuesto{1, "Luz trasera", "Zanella", "Luz de stop trasera", 45433, 920.00, 15})
	t.Add(Repuesto{2, "Embrague", "Yamaha", "Embrague completo con cable de acero", 57383, 1300.50, 5})
	t.Add(Repuesto{3, "Bujia", "Bosch", "Bujia para moto 125cc", 59043, 346.80, 40})

	fmt.Println("La tienda tiene los siguientes repuestos:")
	t.Print()
	fmt.Println("")

	r0 := t.FindByID(0)

	if r0 != nil {
		fmt.Println("El repuesto con ID=0 es:")
		fmt.Printf("  id:[%v]\tnombre:'%v'\tdescripción:'%v'\tnroPieza:'%v'\tprecio:'%.2f'\tstock:'%d'\n", r0.ID, r0.Nombre, r0.Descripcion, r0.NroPieza, r0.Precio, r0.Stock)
	}

	t.Delete(0)

	fmt.Println("La tienda tiene los siguientes repuestos:")
	t.Print()
	fmt.Println("")

	t.Update(Repuesto{2, "Embrague", "Yamaha", "Embrague completo con cable de acero", 57383, 1300.50, 11})

	fmt.Println("La tienda tiene los siguientes repuestos:")
	t.Print()
	fmt.Println("")

}
