package streams

type Channel0 chan interface{}

func (c Channel0) Map(f func(interface{}) interface{}) *Mapper {
	return &Mapper{
		source: c,
		apply:  f,
	}
}

func (c Channel0) Select(f func(interface{}) bool) *Selector {
	return &Selector{
		source: c,
		test:   f,
	}
}

func (c Channel0) ForEach(f func(interface{})) {
	for v := range c {
		f(v)
	}
}
