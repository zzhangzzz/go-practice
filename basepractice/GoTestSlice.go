package main

import "fmt"

func main() {
	map1 := make(map[int]string)
	map1[1] = "01"
	map1[2] = "02"
	map1[3] = "03"

	map2 := map[int]string {
		1 : "01",
		2 : "02",
		3 : "03",
	}
	fmt.Println(map1[0])

	value,ok := map1[1]; if ok {
		fmt.Println("index 1 is", value)
	} else {
		fmt.Println("index 1 is not exist")
	}
	fmt.Println(map2)
}
