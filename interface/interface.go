package main

import (
	"strconv"
)

// Interface
type CarInterface interface {
	Drive() bool
	Stop() bool
	Information() string
	GetBrand() string
	GetModel() string
}

func CarExecute(c CarInterface) {
	println(c.Information())

	msg := ""
	if c.Drive() {
		msg += c.GetBrand() + " " + c.GetModel() + " is driving"
	} else {
		msg += c.GetBrand() + " " + c.GetModel() + " is not driving"
	}

	if c.Stop() {
		msg += " and the car is stopped"
	} else {
		msg += " and the car is not stopped"
	}
	println(msg)
}

// Car is a struct
type Car struct {
	Brand string
	Model string
	Color string
	Speed int
	Price float64
}

type SpecialProduction struct {
	isSpecial bool
}

// Ferrari
type Ferrari struct {
	Car
	SpecialProduction
}

func (_ Ferrari) Drive() bool {
	return false
}

func (_ Ferrari) Stop() bool {
	return true
}

func (ferrari Ferrari) Information() string {
	return ferrari.Brand + " " + ferrari.Model + " " + ferrari.Color + " " + strconv.Itoa(ferrari.Speed) + " " + strconv.FormatFloat(ferrari.Price, 'g', -1, 64) + " " + strconv.FormatBool(ferrari.isSpecial)
}

func (ferrari Ferrari) GetBrand() string {
	return ferrari.Brand
}

func (ferrari Ferrari) GetModel() string {
	return ferrari.Model
}

func NewFerrariWithAllArgs(brand string, model string, color string, speed int, price float64, isSpecial bool) Ferrari {
	return Ferrari{Car{brand, model, color, speed, price}, SpecialProduction{isSpecial}}
}

// Mercedes
type Mercedes struct {
	Car
}

func (_ Mercedes) Drive() bool {
	return true
}

func (_ Mercedes) Stop() bool {
	return false
}

func (mercedes Mercedes) Information() string {
	return mercedes.Brand + " " + mercedes.Model + " " + mercedes.Color + " " + strconv.Itoa(mercedes.Speed) + " " + strconv.FormatFloat(mercedes.Price, 'g', -1, 64)
}

func (mercedes Mercedes) GetBrand() string {
	return mercedes.Brand
}

func (mercedes Mercedes) GetModel() string {
	return mercedes.Model
}

func NewMercedesWithAllArgs(brand string, model string, color string, speed int, price float64) Mercedes {
	return Mercedes{Car{brand, model, color, speed, price}}
}

func main() {

	ferrari := NewFerrariWithAllArgs("Ferrari", "F40", "Red", 320, 1000000, true)
	mercedes := NewMercedesWithAllArgs("Mercedes", "S600", "Black", 250, 500000)
	CarExecute(ferrari)
	CarExecute(mercedes)
}
