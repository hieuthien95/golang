package main

// . để build thành file *.a: GO INSTALL
// . để build thành file .exe trong bin: GO build main.go
// . de chay file .exe: go run ...

import (
	"fmt"
	// 2. my package
	"nordiccoder/week1/variable-package/weekpkg"
	// 3. my package global
	myvalidator "my-validator-pkg"
)

// 1. bien global
var (
	gA = 1
	gB = ""
)

// 1. Bien global
//	  khong yeu cau phai su dung
// 2. my package insite project
// 3. my package global
func main() {
	// 1. Bien global
	fmt.Println("Global:", gA)

	// 2. my package
	fmt.Println("CONST:", weekpkg.DOMAIN, weekpkg.HOST, weekpkg.PORT)
	weekpkg.HelperFunc()

	// 3. my package global
	myvalidator.ValidatorA()
}
