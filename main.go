package main 
import "fmt"

func main() {
	c1 := complex(2,3)
	c2 := 4 + 5i 			   // complex initializer syntax a + ib
	c3 := c1 + c2 	  		  // addition just like other variable 
	fmt.Println("Add: ", c3) // prints "Add: (6+8i)"
	re := real(c3)    		// get real part
	im := imag(c3) 		   // get imaginary part
	fmt.Println(re, im)   // prints 6 8

}