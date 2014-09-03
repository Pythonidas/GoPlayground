package main

import "fmt"
import "math/cmplx"

func Cbrt(x complex128) complex128 {

	var cuberoot, testval complex128 = x / 3, 0

	for cmplx.Abs(cuberoot-testval) > 0.00001 {
		testval = cuberoot
		cuberoot = cuberoot - (cuberoot*cuberoot*cuberoot-x)/(3*cuberoot*cuberoot)
	}
	return cuberoot
}

func main() {
	fmt.Println(Cbrt(2))
	fmt.Println(cmplx.Pow(2, 0.33333))
}
