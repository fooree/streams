package channel

import (
	"github.com/fooree/streams"
)

type Filter struct {
	source streams.Source
	action func(interface{}) bool
	in     chan interface{}
	out    chan interface{}
}

func (f *Filter) In() chan<- interface{} {
	return f.in
}

func (f *Filter) To(fn func(source streams.Source) streams.Sink) streams.Sink {
	return fn(f)
}

func (f *Filter) Filter(fn func(interface{}) bool) streams.Channel {
	act := f.action
	f.action = func(i interface{}) bool {
		return act(i) && fn(i)
	}
	return f
}

func (f *Filter) Map(m func(interface{}) interface{}) streams.Channel {
	return NewChanMap(f, m)
}

func (f *Filter) Out() <-chan interface{} {
	return f.out
}

func (f *Filter) do() {
	go func() {
		for x := range f.source.Out() {
			if f.action(x) {
				f.in <- x
			}
		}
		close(f.in)
	}()

	for v := range f.in {
		f.out <- v
	}
	close(f.out)
}

func NewChanFilter(source streams.Source, action func(interface{}) bool) *Filter {
	f := Filter{
		source: source,
		action: action,
		in:     make(chan interface{}),
		out:    make(chan interface{}),
	}
	go f.do()
	return &f
}
