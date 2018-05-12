package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// go basic types are:
// 1. bool
// 2. string
// 3. int  int8  int16  int32  int64
// 4. uint uint8 uint16 uint32 uint64 uintptr
// 5. byte // alias for uint8
// 6. rune // alias for int32
//     // represents a Unicode code point
//
// 7. float32 float64
// 8. complex64 complex128

// Constants cannot be declared using the := syntax.
// Numeric constants are high-precision values.

// if with short statement
// the if statement can start with a short statement to execute before the condition.

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func myDate() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Sunday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	fmt.Println(today)
}

func main(){
	a := 4.5
	sum := 3
	fmt.Printf("a is of type %T\n\n",a)
	myDate()
	// for
	for i := 0; i < sum; i++ {
		fmt.Println("Sum is ", sum + i)
	}

	// for-continued
	for ; sum < 100; {
		sum += sum
	}

	fmt.Println(sum)

	// for is Go's while
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println(runtime.GOOS)


}