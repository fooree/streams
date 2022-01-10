package main

import (
	"fmt"
	"github.com/fooree/streams/channel"
)

func main() {

	ch := channel.Int(1, 2, 3, 4, 5).
		Map(func(i interface{}) interface{} {
			return i.(int) + 1 + 2
		}).
		Filter(func(t interface{}) bool {
			return t.(int)%2 == 0
		}).
		Filter(func(t interface{}) bool {
			return t.(int) > 3
		}).
		Map(func(t interface{}) interface{} {
			return fmt.Sprintf("s-%d", t)
		}).
		To(channel.NewSink).(*channel.ChanSink).Out

	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("finished0")

	src := make(chan interface{})

	go func() {
		for i := 0; i < 5; i++ {
			src <- i
		}
		close(src)
	}()

	channel.NewSource(src).
		Map(func(i interface{}) interface{} {
			return i.(int) + 1 + 2
		}).
		Filter(func(t interface{}) bool {
			return t.(int)%2 == 0
		}).
		Filter(func(t interface{}) bool {
			return t.(int) > 3
		}).
		Map(func(t interface{}) interface{} {
			return fmt.Sprintf("s-%d", t)
		}).
		To(channel.NewSink).(*channel.ChanSink).
		ForEach(func(i interface{}) {
			fmt.Println(i)
		})

	fmt.Println("finished1")
}
