package main

import (
	"errors"
	"fmt"
	"reflect"
	"unicode/utf8"
)

func main() {
	var intNum uint = 10
	var floatNum float64 = 10000000.8
	fmt.Println("Hello, World!", intNum, floatNum)

	var str string = "d"
	fmt.Println(utf8.RuneCountInString(str))
	fmt.Println(len(str))

	var home rune = 'h'
	fmt.Println(reflect.TypeOf(home))
	fmt.Println(home)

	var1, var2 := 1, 2
	print(var1, var2)

	printme("Hello, World of romil!")
	returnDiv, returnRemain, err := returnDivision(10, 3)
	// printme(fmt.Sprint(returnDiv, returnRemain))
	if err != nil { //we can use || and && operators as well
		fmt.Printf("%s", err.Error())
	} else if returnRemain == 0 {
		fmt.Printf("The result of int division is: %d\n", returnDiv)
	} else {
		fmt.Printf("Division: %d, Remainder: %d\n", returnDiv, returnRemain)
	}

	switch {
	case err != nil:
		fmt.Printf("%s", err.Error())
	case returnRemain == 0:
		fmt.Printf("The result of int division is: %d\n", returnDiv)
	default:
		fmt.Printf("Division: %d, Remainder: %d\n", returnDiv, returnRemain)

	}

	switch returnRemain {
	case 0:
		fmt.Printf("The division was exact, Nothing remaining")
	case 1, 2:
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not exact")
	}

	fmt.Println("\nArray")
	arr := [...]int64{1, 2, 3, 4, 5}
	arr[0] = 1
	arr[1] = 123
	fmt.Println(arr[0])
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Number: %d, Memory Location: %d \n", arr[i],&arr[i])
	}
	fmt.Println(arr[1:5])

	fmt.Println("\nSlice")
	slice := []int64{1, 2, 3, 4, 5}
	slice = append(slice, 6) //creates new array and appends the value, With bigger size
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice)) //length and capacity of the slice, Capacity is more than length


	//create map below
	fmt.Println("\nMap")
	mymap := make(map[string]uint8)
	mymap["key1"] = 1
	mymap["key2"] = 2
	fmt.Println(mymap)

	//check if value exists in map by using ok
	val, ok := mymap["key3"]; 
	if ok {
		fmt.Println(val)
	}

	//delete key from map
	delete(mymap, "key1")
	fmt.Println(mymap)

	mymap["key3"] = 3
	for key, value := range mymap {
		fmt.Println(key, value)
	}

	//while loop using for
	i := 0
	for i<10 {
		fmt.Println(i)
		i++
	}

	//remove condition to make it infinite loop
	for{
		if i == 20 {
			break
		}
		fmt.Println(i)
		i++
	}

	//declare and initialize value in for loop
	for j:=0; j<10; j++ {
		fmt.Printf("Value of j: %d\n", j)
	}

	//struct
	// var myAudi car = car{brand: "Audi", model: "A4", year: 2020}
	// var myAudi car = car{"Audi", "A4", 2020}
	myAudi := car{"Audi", "A4", 2020}
	fmt.Println(myAudi)
	fmt.Println(myAudi.brand)
	pointers()
}

func printme(printValue string) {
	fmt.Println(printValue)
}

func returnDivision(a int, b int) (int, int, error) {
	var err error
	if b == 0 {
		err = fmt.Errorf("Division by zero")
		err = errors.New("Division by zero")
		return 0, 0, err

	}
	return a / b, a % b, err
}

type car struct {
	brand string
	model string
	year int
}

//pointers
func pointers() {
	var a int = 10
	var b *int = &a //b is a pointer to a
	fmt.Println(a, *b)
}