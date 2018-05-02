package main

import "fmt"

func main() {
	// const pi float64 = 3.141592

	// var (
	// 	varA = 2
	// 	varB = "blah blah \n blah"
	// 	truthy = varA == 3
	// )
	// fmt.Println(varA, varB)		//printLn is basically print line, meaning the argument is printed then goes to next line.
	// 							//Printf is print formatter, gives you the ability to mark where the String variables will go, and pass those variables in with it.
	// fmt.Printf("%.2f \n", pi) 	//%.#f will slice off and return only the specified # off the float.
	// fmt.Printf("%T \n", pi) 	// this should return a data type using %T
	// fmt.Printf("%t \n", truthy) // should return a Boolean, in this case it'll be false.
	// fmt.Printf("%d hehehe \n", 100) 	// this should return an integer passed into it.

//----------------------------//

	// 							//3 logical operators in Go:
	// fmt.Println("true && false =", true && false) // should return false because both need to be true.
	// fmt.Println("true && false =", true && true)	// should return a true b/c both are true here.

	// fmt.Println("True || false =", true || false)	// Is looking for something truthy, so will return true.

	// fmt.Println("!true =", !true) 					// will return the opposite, same as javascript.

//----------------------------//
				// for loops
// i := 1
// for i <= 10 {
// 	fmt.Println(i)
// 	i = i+1 					// can also use i ++, or i += 1. As well as i-- to shorthand decriment. 
// }

								// RELATIONAL operators
								//	==, 1_, <, >, <=, >=

	// for j := 0; j < 5; j++{
	// 	fmt.Println(j)
	// }

	// yourAge := 17
	// if yourAge >= 16 && yourAge >= 18 {
	// 	fmt.Println("You can drive and vote")
	// } else if yourAge >= 16 && yourAge < 18 {
	// 	fmt.Println("You can drive but no voting")
	// } else if yourAge < 18 {
	// 	fmt.Println("You can not drive!")
	// }

//--------------- switch case
	// 	yourAge := 33
	// switch yourAge {
	// case 16 : fmt.Println("You cain't drive!")
	// case 18 : fmt.Println("You can drive but no votage")
	// case 33 : fmt.Println("Get outta here dewdz")
	// default : fmt.Println("Brooooo")
	// }

//--------------- 

		// how to initialize an array:
var favNums2[5] float64
favNums2[0] = 33
favNums2[1] = 345345
favNums2[2] = 353
favNums2[3] = 2.385
favNums2[4] = 545
// fmt.Println(favNums2)

		// another way to inialize an array:
		// favNums3 := [5]float64 {1, 535, 55, 62, 77.7}
		// fmt.Println(favNums3)
		//Loop over the array
		
		// for i, value := range favNums3 { 	// replace i with _ to hide indexing
		// 	fmt.Println(value, i)			// remove i from here to hide indexing
		// }									// then it'll only print out the values int he 
		
		// for j := 0; j < len(favNums3); j++{				// why can't I use this method over the above method? Look ingo ranges more.
		// 	fmt.Println(favNums3[j])					// Is one faster than the other?
		// }


		// Slicing an array.
		numSlice := []int {5, 3535, 25, 63, 365}	//slice, like in javascript, starts at the number specified, but ends at the number before specified.
		numSlice2 := numSlice[1:4]
		numSlice3 := favNums2[0:3]

		fmt.Println(numSlice2, numSlice3)

		

}