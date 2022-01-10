package streams

type Source interface {
	Out() <-chan interface{}
}

type Sink interface {
	In() chan<- interface{}
}

type Channel interface {
	Source
	Sink
	Filter
	Map
	To(func(Source) Sink) Sink
}

type Filter interface {
	Filter(func(interface{}) bool) Channel
}

type Map interface {
	Map(func(interface{}) interface{}) Channel
}

type ForEach interface {
	ForEach(func(interface{}))
}
