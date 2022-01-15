package channel

import (
	"github.com/fooree/streams"
)

type ChanSource struct {
	out chan interface{}
}

func (c *ChanSource) Map(f func(interface{}) interface{}) streams.Channel {
	return NewChanMap(c, f)
}

func (c *ChanSource) Filter(f func(interface{}) bool) streams.Channel {
	return NewChanFilter(c, f)
}

func (c *ChanSource) Out() <-chan interface{} {
	return c.out
}

func (c *ChanSource) ForEach(f func(i interface{})) {
	for v := range c.Out() {
		f(v)
	}
}

func (c *ChanSource) ToSink() *ChanSink {
	return NewSink(c).(*ChanSink)
}

func NewSource(ch chan interface{}) *ChanSource {
	return &ChanSource{out: ch}
}

type ChanSink struct {
	source streams.Source
	Out    chan interface{}
}

func (c *ChanSink) ForEach(fn func(interface{})) {
	for v := range c.Out {
		fn(v)
	}
}

func (c *ChanSink) In() chan<- interface{} {
	return c.Out
}

func (c *ChanSink) do() {
	for v := range c.source.Out() {
		c.Out <- v
	}
	close(c.Out)
}

func NewSink(source streams.Source) streams.Sink {
	s := &ChanSink{
		source: source,
		Out:    make(chan interface{}),
	}
	go s.do()
	return s
}
