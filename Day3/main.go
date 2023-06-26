package main

import "fmt"

func main() {

var point = 10

if point == 10 {
		fmt.Println("Perfect")
	} else if point > 5 {
		fmt.Println("Awesome")
	} else if point == 4 {
		fmt.Println("Not Bad")
	} else {
		fmt.Printf("You're Failed. your grade is %d\n", point)
	}

var point2 = 9420.0

	if percent := point2 / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

var point3 = 8

	switch point3 {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not Bad")
	}

var point4 = 6

	switch point4 {
	case 8:
		fmt.Println("Perfect")
	case 7, 6, 5, 4:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not Bad")
	}

var point5 = 6

	switch point5 {
		case 8:
			fmt.Println("Perfect")
		case 7, 6, 5, 4:
			fmt.Println("Awesome")
		default:
			{
				fmt.Println("Not Bad")
				fmt.Println("You can be better!")
			}
	}

var point6 = 2

	switch {
	case point6 == 8:
			fmt.Println("Perfect")
		case (point6 < 8) && (point6 > 3):
			fmt.Println("Awesome")
		default:
			{
				fmt.Println("Not Bad")
				fmt.Println("You can be better!")
			}
	}

var point7 = 3

	switch {
		case point7 == 8:
			fmt.Println("Perfect")
		case (point7 < 8) && (point7 > 3):
			fmt.Println("Awesome")
			fallthrough
		case point7 < 5:
			fmt.Println("You need to learn more")
		default:
			{
				fmt.Println("Not Bad")
				fmt.Println("You can be better!")
			}
	}

var point8 = 2

	if point8 > 7 {
		switch point {
		case 10:
			fmt.Println("Perfect!")
		default:
			fmt.Println("Nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("Not Bad")
		} else if point == 3 {
			fmt.Println("Keep Trying")
		} else {
			fmt.Println("You can do it")
			if point == 0 {
				fmt.Println("Try Harder!")
			}
		}
	}
}