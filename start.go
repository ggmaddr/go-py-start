package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
	"unicode/utf8"
)

var poems = `those who do not feel this love
pulling the mlike a river
those who do not drink dawn
like a cup of spring water
or take in sunset like supper
those who do not want to change
let the sleep`

func main() {
	//decare and init in 2 lines
	var ass string
	ass = "assigned var"

	var b = "init " //init var in one line
	b = " cat "
	const c = 20
	a := "initialized" //init var in one line
	fmt.Println("Testing variables", b, a, "and "+ass)
	fmt.Printf("Use multiple.. %v..container..%v..", c, b) //respectively

	//ask user input
	var userName string
	fmt.Println("What's your name")
	fmt.Scan(&userName) //pass reference

	//string Functions UTF-8
	fmt.Println("Length of username", userName, " is ", utf8.RuneCountInString(userName))

	//use imported time
	fmt.Print("It is", time.Now())

	fmt.Println(Hello("London")) //use function Hello from another file in same dir

	//loop
	var total int
	for i := 1; i <= c; i++ {
		if i%5 == 0 || i%7 == 0 {
			fmt.Print("This is a loop ", i)
			total += i
			fmt.Println(" with total: ", float32(total)) //cast to float
		}
	}
	//range for loop
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	//i = index, v = value: of pow array
	for i, v := range pow {
		fmt.Printf("2**%d = %d \n", i, v)
	}

	//array
	var empty [5]float64              //empty
	floats := []float64{3.14, 2.9, 5} //init without cap

	sort.Float64s(floats) //sorting

	intss := [10]int{4, 8, 7, 4, 9} //cap: 10, len: 5, autofill with 0s
	intss[8] = 99
	copy(empty[0:5], floats) //copy function
	// [2.9 3.14 5 0 0] [2.9 3.14 5] [4 8 7 4 9 0 0 0 99 0]
	fmt.Println(empty, floats, intss)

	//slice
	var s1 []int = intss[1:4] //create slice from arr

	s2 := make([]int, 4, 6) //create slice with make, len=4;cap=6
	s3 := make([]int, 4)    //len=4;cap=4

	arrS := [][]int{s1, s2, s3}
	arrS1 := [5][]int{s1, s2, s3}

	fmt.Println(s2, " s2 has length ", len(s2), " and cap: ", cap(s2))
	fmt.Println(s3, " s2 has length ", len(s3), " and cap: ", cap(s3))
	fmt.Println("Array of slices:", arrS, "arr with declared cap", arrS1)

	//func Frequency
	fmt.Println("Most frequenty word: ", Frequency(poems))

}

func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Function: Hi, %v. Welcome!", name)
	return message
}
func Frequency(paragraph string) string {
	//init map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map init:", n)

	//make(map[key-type]val-type)
	frequency := make(map[string]int)

	//strings.Fields : string.split()
	for _, word := range strings.Fields(paragraph) {
		frequency[word]++
	}
	maxW, maxC := "", 0 //init multi var in one line

	for w, c := range frequency {
		if c > maxC {
			maxW, maxC = w, c //multi var assignment
		}
	}
	return maxW
}
