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

	// file, err := os.Create("test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// line, err := file.WriteString("Hello World")
	// if err != nil {
	// 	fmt.Println(err)
	// 	file.Close()
	// 	return
	// }
	// fmt.Println(line, "bytes written successfully")
	// err = file.Close()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}
