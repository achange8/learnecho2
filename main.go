package main

import (
	"fmt"

	"github.com/achange8/learnecho2/router2"
)

func main() {
	fmt.Println("Welcom osh learnecho 2! ")

	e := router2.New()

	e.Logger.Fatal(e.Start(":8081"))

}
