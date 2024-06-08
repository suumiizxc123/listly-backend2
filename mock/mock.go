package main

import (
	"kcloudb1/internal/config"
)

func main() {

	config.ConnectDatabase()

	// if err := mockSysUser(); err != nil {
	// 	fmt.Println(" error mock sys user : ", err)
	// }

	// if err := mockOrgAndUser(); err != "" {
	// 	fmt.Println(" error mock user : ", err)
	// }

}
