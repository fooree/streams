package channel

var _ = Any

func Int(slice ...int) *ChanSource {
	ch := make(chan interface{})
	go func() {
		for _, t := range slice {
			ch <- t
		}
		close(ch)
	}()
	c := &ChanSource{ch}
	return c
}

func Any(slice ...interface{}) *ChanSource {
	ch := make(chan interface{})
	go func() {
		for _, t := range slice {
			ch <- t
		}
		close(ch)
	}()
	c := &ChanSource{ch}
	return c
}
