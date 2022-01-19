package streams

import "fmt"

func Just(args ...interface{}) Slice {
	return args
}

func From(c chan interface{}) Channel0 {
	return c
}

func foreach(source interface{}, test func(interface{}) bool, apply func(interface{})) {
	switch t := source.(type) {
	case Slice:
		t.forSelect(test, apply)
	case Channel0:
		for v := range t {
			if test(v) {
				apply(v)
			}
		}
	default:
		panic(fmt.Errorf("unknown type: %T", source))
	}
}
