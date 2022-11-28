package main

import "fmt"

func main() {
	var width, height, area float64
	width = 4.2
	height = 3.0
	area = width * height
	fmt.Println(fmt.Sprintf("%0.2f", area))
	width = 1.2
	height = 4.1
	area = width * height
	fmt.Println(area)
}

func show() {

}
