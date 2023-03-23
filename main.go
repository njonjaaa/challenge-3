package main

import "fmt"

func main() {

	str1 := "selamat malam"

	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c\n", str1[i])
	}

	countMap := map[string]int{}

	for _, s := range str1 {
		countMap[string(s)] += 1
	}

	fmt.Println(countMap)

}
