package main

import (
	"fmt"
	"strings"
)

func main() {
	// fizzbuzz()
	// stringExammple()
	// evenAdd()
	// sliceExample()
	// findMaxValue([]int {1, 2 ,45, 2, 33})
	// mapsExample()
	wordCount("Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.")
}

func fizzbuzz() {
	fmt.Println("----------------------")
	fmt.Println("Example for .... if: Fizz buzz")
	fmt.Println("Description: Loop i from 1 to 20, if i divisible by 3 then print fizz istead of i, divisible by 5 print buzz, and divisible both 3 and 5 then print fizz buzz")
	fmt.Println("Result:")
	for i:=1; i<=20; i++ {
		if (i % 15 == 0) {
			fmt.Println("Fizz Buzz")
			continue
		}

		if (i % 3 == 0) {
			fmt.Println("Fizz")
			continue
		}

		if (i % 5 == 0) {
			fmt.Println("Buzz")
			continue
		}

		fmt.Println(i)
	}
}

func stringExammple() {
	fmt.Println("----------------------")
	fmt.Println("Example string")
	x := "Hello Nguyen Tuan Vu"
	fmt.Printf("The full string: %v\n", x)
	fmt.Printf("The hello statement: %v\n", x[:5])
	fmt.Printf("The name: %v\n", x[6:])
}

func evenAdd() {
	fmt.Println("----------------------")
	fmt.Println("Even Add Number Challange")
	fmt.Println("Description: Even Add Number is the number with the same first and last digit. How many even-added number result for mutiply two four-digits numbers")
	fmt.Println("Result:")
	n:=0;
	for i:=1000; i<10000; i++ {
		for j:=i; j<10000; j++ {
			s:= fmt.Sprintf("%d", i*j)
			if (s[0] == s[len(s)-1]) {
				n++
			}
		}
	}
	fmt.Printf("%v", n)
}

func sliceExample() {
	fmt.Println("----------------------")
	fmt.Println("Example slice")
	str := []string{"vu", "huyen", "mai", "hung", "ha"}
	for i:=0; i<len(str); i++ {
		fmt.Printf("%v ", i)
	}
	fmt.Println();
	// loop index 
	fmt.Print("Single index range: ")
	for i:= range str {
		fmt.Printf("%v ", str[i])
	}
	fmt.Println()
	// loop value
	fmt.Print("Single value range: ")
	for _, v:= range str {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
	// loop double 
	fmt.Print("Double value range: ")
	for i, v:= range str {
		fmt.Printf("%v:%v, ", i, v)
	}
}

func findMaxValue(str []int) {
	fmt.Println("----------------------")
	fmt.Println("Max value Challange")
	fmt.Println("Description: find the max value of slice str")
	fmt.Println("Result:")
	x := str[0]
	for _,v := range str {
		if (x < v) {
			x = v
		}
	}
	fmt.Println(x)
}

func mapsExample() {
	fmt.Println("----------------------")
	fmt.Println("Map Example")
	binances := map[string]float64 {
		"BTC": 21326.5,
		"ETH": 1224.9,
		"USDT": 1,
		"BNB": 239.9,
	}

	// length of map 
	fmt.Println("Length of map: ", len(binances))
	// get map key 
	fmt.Println("BTC prices: ", binances["BTC"])
	// set map key
	binances["SOL"] = 40.57 
	fmt.Println("SOL prices add: ", binances["SOL"])
	// loop map key
	fmt.Printf("All coin: ")
	for k := range binances {
		fmt.Printf("%v ", k)
	}
	fmt.Println()
	
	// loop keys and value
	fmt.Printf("All prices: ")
	for k,v := range binances {
		fmt.Printf("(%v, %v) ", k, v)
	}
	fmt.Println()
}

func wordCount(para string) {
	fmt.Println("----------------------")
	fmt.Println("Word Count Challange")
	fmt.Println("Description: count number word of paragraph")
	fmt.Println("Result:")
	para = strings.ReplaceAll(para, ",", "")
	para = strings.ReplaceAll(para, ".", "")
	para = strings.ReplaceAll(para, "  ", "")
	paraTmp := strings.Split(para, " ")
	res := map[string]float64 {}
	for _,v := range paraTmp {
		_, ok := res[v]
		if  ok {
			res[v] ++ 
		} else {
			res[v] = 1
		}
	}
	for k, v := range res {
		fmt.Println(k, ":", v)
	}
}

