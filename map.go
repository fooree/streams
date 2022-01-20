package streams

import "fmt"

type Mapper struct {
	source interface{}
	apply  func(interface{}) interface{}
}

func (m *Mapper) Map(f func(interface{}) interface{}) *Mapper {
	fn := m.apply
	m.apply = func(i interface{}) interface{} {
		return f(fn(i))
	}
	return m
}

func (m *Mapper) Select(f func(interface{}) bool) *Selector {
	s := &Selector{
		source: make(Channel0),
		test:   f,
	}
	go m.flow(s)
	return s
}

func (m *Mapper) flow(s *Selector) {
	ch := s.source.(Channel0)
	defer close(ch)
	switch x := m.source.(type) {
	case Slice:
		for i := 0; i < len(x); i++ {
			ch <- m.apply(x[i])
		}
	case Channel0:
		for v := range x {
			ch <- m.apply(v)
		}
	default:
		panic(fmt.Errorf("unknown source type: %T", s.source))
	}
}

func (m *Mapper) ForEach(each func(interface{})) {
	mapEach(m.source, m.apply, each)
}
