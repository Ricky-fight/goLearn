package main

import (
	"fmt"

	snow "github.com/Ricky-fight/goLearn/2/snow"
)

func main() {
	rst := snow.Exec("add", 1.0, 2.0)
	fmt.Println(rst)
}
