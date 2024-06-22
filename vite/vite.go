package vite_boil

import (
	"fmt"

	"sam.crider/boilerplate-script/utils"
)


func Vite() {
	cmd := utils.BoundCommand("npm", "create", "vite@latest")

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
