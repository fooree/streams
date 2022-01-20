package streams

import "fmt"

func Just(args ...interface{}) Slice {
	return args
}

func From(c chan interface{}) Channel0 {
	return c
}

func ToChannel(do func(chan interface{})) Channel0 {
	ch := make(Channel0)
	go func() {
		defer close(ch)
		do(ch)
	}()
	return ch
}

func selectEach(source interface{}, test func(interface{}) bool, each func(interface{})) {
	switch t := source.(type) {
	case Slice:
		t.forSelect(test, each)
	case Channel0:
		t.forSelect(test, each)
	default:
		panic(fmt.Errorf("unknown type: %T", source))
	}
}

func mapEach(source interface{}, map_ func(interface{}) interface{}, each func(interface{})) {
	switch t := source.(type) {
	case Slice:
		t.forMap(map_, each)
	case Channel0:
		t.forMap(map_, each)
	default:
		panic(fmt.Errorf("unknown type: %T", source))
	}
}
