package main

import "fmt"

func find_pi(n int) float64{

	var piOver4 float64
	sign := 1.0

	for i:=1; i<=n; i+=2{
		piOver4 += sign*(1/float64(i))
		sign*=-1
	}

	pi := piOver4*4
	return pi
}

func main() {

	var n int = 0
	fmt.Println("Please enter a value for n. \nThe higher the value of n, the more accurate your resulting pi value:")
	fmt.Scanln(&n)
	fmt.Println("Here is Pi: ", find_pi(n))

}
