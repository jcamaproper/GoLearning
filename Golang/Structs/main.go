package main

import "fmt"

type contactInfo struct {
	country string
	zip     int
}

type person struct {
	fisrtName string
	lastName  string
	age       int
	address   contactInfo
}

func main() {
	//DATOS CON UNA ESTRUCTURA NORMAL CON DATOS DE TIPO COMUN
	/* //datos de tipo structura
	var juan, fer person
	//arreglo de datos tipo estructuras
	var personList []person
	juan = person{"Juan", "Camacho", 25}
	fer = person{"Fernando", "Barrera", 25}

	personList = append(personList, juan)
	personList = append(personList, fer)
	juan := person{}
	juan := person{"Juan", "Camacho", 25}
	juan.fisrtName = "Juan"
	juan.lastName = "Camacho"
	juan.age = 25
	fmt.Println(juan)
	fmt.Println(personList) */

	//DATOS CON UNA EMBEDD STRUCTURE
	juan := person{
		fisrtName: "Juan",
		lastName:  "Camacho",
		age:       25,
		address: contactInfo{
			country: "Mexico",
			zip:     07040,
		},
	}

	//se crea un puntero para hacer referencia al espacio de memoria donde está la variable
	//asgino la dirección de memoria donde esta esa variable
	/* j := &juan
	j.updateFirsName("Toño")
	juan.print() */
	juan.updateFirsName("Toño")
	juan.print()
}

func (p person) print() {
	fmt.Println(p)
}

//en lugar de actualizar, asgina el valor en otro espacio de memoria, no en el de la variable que quiero modifcar
/* func (p person) updateFirsName(newName string) {
	p.fisrtName = newName
} */

//recibo un puntero para modificar el espacio de memoria, no copiarlo en otro
func (p *person) updateFirsName(newName string) {
	//con el * recibo el valor que esta en el espacio de memoria de esta variable y lo modifico
	p.fisrtName = newName
}
