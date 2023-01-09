package helper

// ChanWriter ...
type ChanWriter struct {
	ch chan byte
}

// NewChanWriter ...
func NewChanWriter() *ChanWriter {
	return &ChanWriter{make(chan byte, 1024)}
}

// Chan ...
func (w *ChanWriter) Chan() <-chan byte {
	return w.ch
}

func (w *ChanWriter) Write(p []byte) (int, error) {
	n := 0
	for _, b := range p {
		w.ch <- b
		n++
	}
	return n, nil
}

// Close ...
func (w *ChanWriter) Close() error {
	close(w.ch)
	return nil
}
