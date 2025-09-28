package main 
import "fmt"

func main() {
	// Signed integers
	var a int8 = 127
	var b int16 = 32767
	var c int32 = 2147483647
	var d int64 = 9223372036854775807
	//Unsigned integers
	var e uint8 = 255
	var f uint16 = 65535
	var g uint32 = 4294967295
	var h uint64 = 18446744073709551615
	//print the values 
	fmt.Println("Signed Integers:")
	fmt.Println("int8: %d\n", a)
	fmt.Println("int16: %d\n", b)
	fmt.Println("int32: %d\n", c)
	fmt.Println("int64: %d\n", d)
	fmt.Println("Unsigned Integers")
	fmt.Println("uint8: %d\n", e)
	fmt.Println("uint16: %d\n", f)
	fmt.Println("uint32: %d\n", g)
	fmt.Println("uint64: %d\n", h)

 }
