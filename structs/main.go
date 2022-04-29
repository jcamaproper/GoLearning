package main

import (
	"fmt"
	"strconv"
	"strings"
)

type names struct {
	fName string
	lName string
}

func main() {
	var friends []names

	friend1 := names{
		fName: "Juan",
		lName: "Camacho",
	}
	friends = append(friends, friend1)
	friend2 := names{
		fName: "Fer",
		lName: "Barrera",
	}
	friends = append(friends, friend2)
	friend3 := names{
		fName: "Cesar",
		lName: "Alcibar",
	}
	friends = append(friends, friend3)

	fmt.Println(friends[1].fName)

	depositAmount := strings.Replace("1,110.00", ",", "", -2)
	fmt.Println(depositAmount)
	amount, res := strconv.ParseFloat(depositAmount, 64)
	fmt.Println(amount, res)
}
