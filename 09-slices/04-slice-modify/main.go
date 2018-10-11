package main

import "fmt"

func main() {
	a := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	slice1 := a[1:]
	slice2 := a[3:]

	fmt.Println("------- Before Modifications -------")
	fmt.Println("a  = ", a) // a  =  [Mon Tue Wed Thu Fri Sat Sun]
	fmt.Println("slice1 = ", slice1) // slice1 =  [Tue Wed Thu Fri Sat Sun]
	fmt.Println("slice2 = ", slice2) // slice2 =  [Thu Fri Sat Sun]

	slice1[0] = "TUE"
	slice1[1] = "WED"
	slice1[2] = "THU"

	slice2[1] = "FRIDAY"

	fmt.Println("\n-------- After Modifications --------")
	fmt.Println("a  = ", a) // a  =  [Mon TUE WED THU FRIDAY Sat Sun]
	fmt.Println("slice1 = ", slice1) // slice1 =  [TUE WED THU FRIDAY Sat Sun]
	fmt.Println("slice2 = ", slice2) // slice2 =  [THU FRIDAY Sat Sun]
}
