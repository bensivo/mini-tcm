package model

import "fmt"

type TestCase struct {
	Id               string
	Name             string
	SpecMd           string
	StepsMd          string
	ExpectedResultMd string
}

func (tc *TestCase) Print() {
	fmt.Println("-----")
	fmt.Printf("Id: %s\n", tc.Id)
	fmt.Println("-----")
	fmt.Printf("Name: %s\n", tc.Id)
	fmt.Println("-----")
	fmt.Printf("Spec:\n-----\n%s\n", tc.SpecMd)
	fmt.Println("-----")
	fmt.Printf("Steps:\n-----\n%s\n", tc.StepsMd)
	fmt.Println("-----")
	fmt.Printf("ExpectedResult:\n-----\n%s\n", tc.ExpectedResultMd)
	fmt.Println("-----")
}
