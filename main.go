package main

import (
	"fmt"
	"time"
)

const globalConst = "Global Constant"

func main() {
	fmt.Println("Hello, World!")

	var name string = "Alice"

	fmt.Println("Name:", name)

	// short hand

	abc := "Alice"
	fmt.Println("abc:", abc)

	var bob string
	bob = "Bob"
	fmt.Println("bob:", bob)

	// constant

	const word string = "Alice"
	// shorthand
	fmt.Println("word:", word)

	fmt.Println("globalConst:", globalConst)

	// grouping

	const (
		port  = 5000
		host1 = "localhost"
	)

	fmt.Println("port:", port)
	fmt.Println("host1:", host1)

	// for looping

	// while loop
	i := 1
	for i <= 3 {
		fmt.Println("i:", i)
		i++
	}

	// infinite loop
	// for {
	// 	fmt.Println("This will run forever")
	// }

	// classic for loop

	for z := 0; z <= 3; z++ {
		if z == 2 {
			continue
		}
		fmt.Println("z:", z)
	}

	// range
	for r := range 11 {
		fmt.Println("r:", r)
	}

	//if else

	age := 70
	if age < 18 {
		fmt.Println("Minor")
	} else if age >= 18 && age < 60 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Senior")
	}

	// or operator

	role := "admin"

	hasPermission := false
	if role == "admin" || hasPermission {
		fmt.Println("Access Granted")
	}

	// and operator

	if role == "admin" && hasPermission {
		fmt.Println("Access Granted")
	}
	// variable declare inside if construct
	if age := 15; age >= 18 {
		fmt.Println("person is adult")
	} else if age >= 15 {
		fmt.Println("person is a teenager")
	}

	// switch

	switch role {
	case "admin":
		fmt.Println("Admin Access")
	case "user":
		fmt.Println("User Access")
	default:
		fmt.Println("No Access")
	}

	// multiple condition switch

	switch time.Now().Weekday() {
	case time.Monday, time.Tuesday, time.Wednesday:
		fmt.Println("Start of the week")
	case time.Friday, time.Thursday:
		fmt.Println("End of the week")
	default:
		fmt.Println("Midweek")
	}

	// type switch

	whoAmI := func(i interface{}) {
		switch v := i.(type) {
		case string:
			fmt.Println("I'm a string:", v)
		case int:
			fmt.Println("I'm an int:", v)

		case bool:
			fmt.Println("I'm a bool:", v)

		case float64:
			fmt.Println("I'm a float64:", v)
		default:
			fmt.Println("I'm something else:", v)
		}
	}

	whoAmI("Hello")
	whoAmI(42)
	whoAmI(3.14)

	// arrays
	var nums [4]int

	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4

	fmt.Println("nums:", nums)
	fmt.Println(len(nums))

	var vals [4]string

	vals[0] = "A"
	vals[1] = "B"
	vals[2] = "C"
	vals[3] = "D"

	fmt.Println("vals:", vals)

	lums := [4]int{1, 2, 3, 4}
	fmt.Println("lums:", lums)

	//2d arrays
	matrix := [2][2]int{
		{1, 2},
		{3, 4},
	}

	fmt.Println("matrix:", matrix)

	// slices --> dynamic

	var sil = make([]int, 0, 5) // length 0, capacity 5
	sil = append(sil, 1)
	sil = append(sil, 2)
	sil = append(sil, 3)

	sil = append(sil, 4)
	sil = append(sil, 5)
	sil = append(sil, 6) // length 6, capacity 10
	sil = append(sil, 7) // length 7, capacity 10

	fmt.Println("sil:", sil)

	fmt.Println("cap(sil):", cap(sil))

	// copy function

}
