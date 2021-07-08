package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 6)
	fmt.Printf("长度：%d \n", len(slice))
	fmt.Printf("容量：%d \n", cap(slice))
	for _, s := range slice {
		fmt.Printf("%d \n", s)
	}

	dict := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
		"Red":         "#da1337",
		"Orange":      "#e95a22"}
	value, exists := dict["Red"]
	if exists {
		fmt.Printf("%s \n", value)
	}

	for key, value := range dict {
		fmt.Printf("%s：%s \n", key, value)
	}

	removeColor(dict, "Red")

	for key, value := range dict {
		fmt.Printf("%s：%s \n", key, value)
	}
}

func foo(slice []int) []int {
	return slice
}

func removeColor(colors map[string]string, key string) {
	delete(colors, key)
	fmt.Printf("remove color %s success \n", key)
}
