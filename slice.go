package streams

type Slice []interface{}

func (s Slice) Map(f func(interface{}) interface{}) *Mapper {
	return &Mapper{
		source: s,
		apply:  f,
	}
}

func (s Slice) Select(f func(interface{}) bool) *Selector {
	return &Selector{
		source: s,
		test:   f,
	}
}

func (s Slice) ForEach(f func(interface{})) {
	for i := 0; i < len(s); i++ {
		f(s[i])
	}
}

func (s Slice) forSelect(test func(interface{}) bool, each func(interface{})) {
	for i := 0; i < len(s); i++ {
		v := s[i]
		if test(v) {
			each(v)
		}
	}
}

func (s Slice) ToSlice() []interface{} {
	return s
}
