package main

import "fmt"

var (
	nu1    int
	nu2    float32
	nu3    float64
	n      string
	verify bool
)

const (
	role = "admin"

	age = 20

	grade = 20.5

	notVerify = false
)

func main() {
	// var city = "Kathmandu"
	var name string

	namess, salary, age, veri := "shakti", 23.89, 12, false

	var count, bithNumber = 1, 2

	var umer, ages int = 1, 6

	message := "akjsb"

	n1 := 4

	verified := false

	const human = "amu"

	name = "shakti"

	const num1 = 50

	var num float32

	num = 100

	var num2 = 60

	var total float32
	total = num + num1
	fmt.Println("the sum is", "\t", total, "\t", "name is", "\t", name, "\t", "second number is", "\t", num2)
	fmt.Println("message", "\t", message)
	fmt.Println("numbers", "\t", n1)
	fmt.Println(verified)

	fmt.Println(namess, salary, age, veri)
	fmt.Println(count, bithNumber)
	fmt.Print(umer, ages)

}
