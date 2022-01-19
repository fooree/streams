package streams

import (
	"fmt"
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
	}).ForEach(func(i interface{}) {
		fmt.Println(i)
	})
}
