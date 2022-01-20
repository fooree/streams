package streams

import (
	"fmt"
	"strings"
	"testing"
)

func TestJust(t *testing.T) {
	Just(1, 2, 3, 4, 5).Select(func(i interface{}) bool {
		return i.(int)%2 == 1
	}).Select(func(i interface{}) bool {
		return i.(int) > 2
	}).Map(func(i interface{}) interface{} {
		return fmt.Sprintf("%v", i)
	}).Select(func(i interface{}) bool {
		return i.(string) == "5"
	}).Map(func(i interface{}) interface{} {
		return "s" + i.(string)
	}).ForEach(func(i interface{}) {
		fmt.Println(i)
	})
}

func TestFrom(t *testing.T) {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	From(ch).Select(func(i interface{}) bool {
		return i.(int)%2 == 1
	}).Select(func(i interface{}) bool {
		return i.(int) > 2
	}).Map(func(i interface{}) interface{} {
		return fmt.Sprintf("%v", i.(int)*10)
	}).Select(func(i interface{}) bool {
		return len(i.(string)) > 0
	}).Map(func(i interface{}) interface{} {
		return "s" + i.(string)
	}).ForEach(func(i interface{}) {
		fmt.Println(i)
	})
}

func TestToChannel(t *testing.T) {
	slice := []string{"a", "b", "c", "c", "e"}
	ToChannel(func(c chan interface{}) {
		for _, s := range slice {
			c <- strings.ToUpper(s)
		}
	}).Map(func(i interface{}) interface{} {
		return strings.Repeat(i.(string), 3)
	}).ForEach(func(i interface{}) {
		fmt.Println(i)
	})
}
