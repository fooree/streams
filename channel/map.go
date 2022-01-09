package channel

import (
	"github.com/fooree/streams"
)

type channel struct {
	source streams.Source
	in     chan interface{}
	out    chan interface{}
}

func (c *channel) Out() <-chan interface{} {
	return c.out
}

func (c *channel) In() chan<- interface{} {
	return c.in
}

func (c *channel) To(f func(streams.Source) streams.Sink) streams.Sink {
	return f(c)
}

type Map struct {
	*channel
	action func(interface{}) interface{}
}

func (m *Map) Filter(f func(interface{}) bool) streams.Channel {
	return NewChanFilter(m, f)
}

func (m *Map) Map(f func(interface{}) interface{}) streams.Channel {
	action := m.action
	m.action = func(i interface{}) interface{} {
		return f(action(i))
	}
	return m
}

func (m *Map) do() {
	go func() {
		for x := range m.source.Out() {
			m.in <- x
		}
		close(m.in)
	}()

	for v := range m.in {
		m.out <- m.action(v)
	}
	close(m.out)
}

func NewChanMap(source streams.Source, action func(interface{}) interface{}) *Map {
	f := &Map{
		channel: &channel{
			source: source,
			in:     make(chan interface{}),
			out:    make(chan interface{}),
		},
		action: action,
	}
	go f.do()
	return f
}
