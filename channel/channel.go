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

func NewChanSource(ch chan interface{}) *ChanSource {
	return &ChanSource{out: ch}
}

type ChanSink struct {
	source streams.Source
	Out    chan interface{}
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

func NewChanSink(source streams.Source) streams.Sink {
	s := &ChanSink{
		source: source,
		Out:    make(chan interface{}),
	}
	go s.do()
	return s
}
