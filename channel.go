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

func (c Channel0) forSelect(test func(interface{}) bool, each func(interface{})) {
	for v := range c {
		if test(v) {
			each(v)
		}
	}
}

func (c Channel0) forMap(mapFunc func(interface{}) interface{}, each func(interface{})) {
	for v := range c {
		each(mapFunc(v))
	}
}
