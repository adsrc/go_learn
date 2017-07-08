package main

import "fmt"
import "time"

func main() {

	i := "Aditya"

	s := time.Now().Hour()
	fmt.Println(s)

	whoAmI := func(i interface{}) {
		switch t := i.(type) {

		case bool:
			fmt.Println("Boolean")
		case int:
			fmt.Println("Integer")
		default:
			fmt.Println(t)
		}
	}

	whoAmI(true)
	whoAmI(i)

}

// TODO: Learn more about interfaces
// TODO: Learn about packages
