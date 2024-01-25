package main

import (
	"fmt"

	adminLoginRepository "github.com/amerikarno/icoApi/repository/admin"
)

func main() {
	pass := adminLoginRepository.NewAdminPassword()
	password := pass.GeneratePassword(13)
	fmt.Printf("password is %s", password)
}
