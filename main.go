package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	//r := rand.New(rand.NewSource(99))

	ch := make(chan int)

	go func() {
		nums := []int{20, 30, 40}
		nums = append(nums, 90, 100, 110, 120, 130, 140, rand.Intn(50))
		fmt.Println("Slice capacity: ", cap(nums), "Slice len: ", len(nums))

		for index, num := range nums {
			time.Sleep(1 * time.Millisecond)
			fmt.Println("Index: ", index, ", Value: ", num)
		}

		fmt.Println("result from go routine")
		ch <- nums[len(nums)-1]
	}()

	fmt.Printf("hello, world\n")
	fmt.Printf(textout("hello"))

	l := list.New()
	l.PushBack(21)

	fmt.Printf("\r\n" + strconv.Itoa(l.Len()))
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value) // print out the elements
	}

	fmt.Println("From channel: ", <-ch)

}

func textout(text string) string {
	return text + " Done"
}

func add(x int, y int) int {
	return x + y
}
