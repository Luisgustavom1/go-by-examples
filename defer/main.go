package main
 
import "fmt"

func modifyingReturn() (i int) {
	defer func() { i++ }()
	return 1
}

// The convention in the Go libraries is that even when a package uses panic internally, its external API still presents explicit error return values.
func panicAndRecover() {
	defer func() {	
		if r := recover(); r != nil {
			fmt.Println("Recovered in panicAndRecover: [PANIC]", r)
		}
	}()

	fmt.Println("I will panicAndRecover the panic")
   	panic("panice happened\n") // any deferred function are executed
}

func main() {
     str := "init"
	 defer fmt.Println(str) // defer functions are evaluated when defer statement is evaluated, so this function was evaluated when str is "init"
	 defer fmt.Println("its");
	 defer fmt.Println("a");
	 defer fmt.Println("LIFO");
    
	 fmt.Println(modifyingReturn())

	 panicAndRecover()
			
	 str = "end"
 }
