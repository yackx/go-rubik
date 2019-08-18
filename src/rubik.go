package main

import (
	"fmt"
	"rubik"
)

func main() {
	cube := rubik.NewCube("sssssssssqqqqqqqqqsssssssssqqqqqqqqqsssssssssqqqqqqqqq")
	fmt.Printf("Cube:\n%s %v\n", cube, cube.IsSolved())

	solved := rubik.NewSolvedCube()
	fmt.Printf("Solved cube\n%s %v\n", solved, solved.IsSolved())

	copied := solved.Copy()
	fmt.Printf("copied:\n%s %v\n", copied, copied.IsSolved())

	f := solved.F()
	fmt.Printf("F:\n%s %v\n", f, f.IsSolved())

	fc := solved.Fc()
	fmt.Printf("F':\n%s %v\n", fc, fc.IsSolved())

	back := solved.F().Fc()
	fmt.Printf("FF'\n%s %v\n", back, back.IsSolved())

	d := solved.D()
	fmt.Printf("o D:\n%s\n", d)
}
