package main

import "fmt"

func sumInt(m map[string]int64)int64 {
	var res int64
	for _,v := range m {
		res += v
	}
	return res
}

func sumFloat(m map[string]float64)float64{
	var res float64
	for _,v := range m {
		res += v
	}
	return res
}


func genericSum[K comparable, V int64 | float64](m map[K]V) V{
	var res V
	for _,v := range m {
		res += v
	}
	return res
}


type Number interface{
	int64 | float64
}

func streamlineGenericSum[K comparable, V Number](m map[K]V) V{
	var res V
	for _,v := range m {
		res += v
	}
	return res
}

//error
// func streamlineGenericSum1[K comparable](m map[K]Number) Number{
// 	var res Number
// 	for _,v := range m {
// 		res += v
// 	}
// 	return res
// }

func main(){
	intmap := map[string]int64 {
		"x": 123,
		"y": 123,
	} 
	floatmap := map[string]float64 {
		"x": 123.12,
		"y": 123.89,
	} 
	fmt.Printf("no generic %v \n", sumInt(intmap))
	fmt.Printf("no generic %v \n", sumFloat(floatmap))
	fmt.Printf("no generic %v \n", genericSum[string, int64](intmap))
	fmt.Printf("no generic %v \n", genericSum[string, float64](floatmap))
	fmt.Printf("generic %v \n", genericSum(intmap))
	fmt.Printf("generic %v \n", genericSum(floatmap))
	fmt.Printf("streamgeneric %v \n", streamlineGenericSum(intmap))
	fmt.Printf("steamgeneric %v \n", streamlineGenericSum(floatmap))
}