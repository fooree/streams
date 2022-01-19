package streams

import "fmt"

type Selector struct {
	source interface{}
	test   func(interface{}) bool
}

func (s *Selector) Map(f func(interface{}) interface{}) *Mapper {
	m := &Mapper{
		source: make(Channel0),
		apply:  f,
	}
	go s.convert(m)
	return m
}

func (s *Selector) convert(m *Mapper) {
	ch := m.source.(Channel0)
	defer close(ch)
	switch x := s.source.(type) {
	case Slice:
		for i := 0; i < len(x); i++ {
			if s.test(x[i]) {
				ch <- x[i]
			}
		}
	case Channel0:
		for v := range x {
			if s.test(v) {
				ch <- v
			}
		}
	default:
		panic(fmt.Errorf("unknown source type: %T", s.source))
	}
}

func (s *Selector) Select(f func(interface{}) bool) *Selector {
	fn := s.test
	s.test = func(i interface{}) bool {
		return fn(i) && f(i)
	}
	return s
}

func (s *Selector) ForEach(f func(interface{})) {
	foreach(s.source, s.test, f)
}
