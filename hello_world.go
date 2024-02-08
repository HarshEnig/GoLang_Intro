package main

import "fmt"

func testingVariables() {
	var i = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)
}
func main() {
	fmt.Println("Hello, World!")
	fmt.Println("go" + " lan")

	//this is a comment
	/*also a comment*/
	var student1 string = "Harsh"
	var student2 = "Jane"
	x := 2
	var a, b, c, d int = 1, 3, 5, 7
	const myConst int = 23

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Print(x, "\n")
	fmt.Print(a+b+c+d+myConst, "\n")
	fmt.Print(a, b)
	testingVariables()
}
