package enum

import "fmt"

// type definition
type Brand string

const (
	BMW      Brand = "Bmw"
	MERCEDES Brand = "Mercedes"
)

func PrintBrand(brand Brand) {
	fmt.Print(brand)
}
