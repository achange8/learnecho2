package main

import (
	"fmt"

	"github.com/achange8/learnecho2/router"
)

func main() {
	fmt.Println("Welcom osh learnecho 2! ")

	e := router.New()

	e.Logger.Fatal(e.Start(":8081"))

}
