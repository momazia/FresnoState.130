/*
	Source: 		Column AP in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	5) Write a function, foo, which can be called in all of these ways:
					func main() {
						foo(1, 2)
						foo(1, 2, 3)
						aSlice := []int{1, 2, 3, 4}
						foo(aSlice...)
						foo()
					}
*/
package main

func main() {

	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}

func foo(input ...int) {

}
