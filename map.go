package streams

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
	return &Selector{
		source: m.source,
		test: func(i interface{}) bool {
			return f(m.apply(i))
		},
	}
}

func (m *Mapper) ForEach(f func(interface{})) {
}
