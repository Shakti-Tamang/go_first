// package main Creates an EXECUTABLE
// Library Package vs Executable Package:




// What package main Actually Does
// Technical Work:
// Tells compiler: "Create an executable binary"

// Provides entry point: Must contain func main()

// Generates OS-specific executable: .exe (Windows), binary (Linux/Mac)

package main // ← Creates a runnable program

import "fmt"


// Your Go files (.go)
//     ↓
// go build
//     ↓
// Compiler sees package main
//     ↓
// Creates executable with main() as entry point
//     ↓
// myapp.exe (Windows) or myapp (Linux/Mac)

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

	fmt.Println("go final"+role);
	fmt.Println("the sum is", "\t", total, "\t", "name is", "\t", name, "\t", "second number is", "\t", num2)
	fmt.Println("message", "\t", message)
	fmt.Println("numbers", "\t", n1)
	fmt.Println(verified)

	fmt.Println(namess, salary, age, veri)
	fmt.Println(count, bithNumber)
	fmt.Print(umer, ages)

}

// Basic Data Types:
// Booleans:
// bool: Represents truth values, either true or false.
// Numeric Types:
// Integers:
// int, int8, int16, int32, int64: Signed integers of varying bit sizes. int typically defaults to the system's word size (32 or 64 bits).
// uint, uint8, uint16, uint32, uint64: Unsigned integers of varying bit sizes. uint also typically defaults to the system's word size.
// byte: An alias for uint8, commonly used for raw data.
// rune: An alias for int32, used to represent Unicode code points (characters).
// uintptr: An integer type large enough to hold the bit pattern of any pointer.
// Floating-point Numbers:
// float32, float64: Floating-point numbers with 32-bit and 64-bit precision, respectively. float64 is the default float type.
// Complex Numbers:
// complex64, complex128: Complex numbers with 32-bit and 64-bit floating-point components, respectively.
// String Type:
// string: Represents immutable sequences of UTF-8 encoded bytes (text).
// Composite Data Types:
// These are built using the basic types or other composite types.
// Arrays: Fixed-size sequences of elements of the same type.
// Slices: Dynamically-sized, flexible views into arrays.
// Structs: Collections of fields (members) of potentially different types, grouped under a single name.
// Pointers: Variables that store memory addresses of other variables.
// Functions: First-class citizens in Go, meaning they can be assigned to variables, passed as arguments, and returned from other functions.
// Interfaces: Define a set of method signatures that a type must implement to satisfy the interface.
// Maps: Unordered collections of key-value pairs, where keys are unique and of a comparable type.
// Channels: Used for communication between goroutines (concurrently executing functions).
