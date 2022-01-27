package main

import "fmt"

func main() {

	/* colors := map[string]string{
		"Red": "",
	}

	colors["Red"] = "000q" */

	myMap := make(map[string]string)
	myMap["Name"] = "Juan"
	myMap["LastName"] = "Camacho"
	myMap["Email"] = "jcamacholpz@gmail.com"
	myMap["country"] = "Mexico"
	fmt.Println(myMap)
	printMap(myMap)

}

func printMap(m map[string]string) {

	for key, data := range m {
		fmt.Printf("The Key is %v the value is %v\n", key, data)
	}
}
