package main

import (
	"fmt"

	"github.com/dopaemon/artus/internal/libutils"
)

func main() {
	if name, err := libutils.GetCPUName(); err == nil {
		fmt.Println("CPU Name: ", name)
	} else {
		fmt.Println("Err: ", err)
	}
}
