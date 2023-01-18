package myPackage

type Response struct {
	Status     string
	StatusCode int
	Body       ReadCloser
}

type ReadCloser interface {
	Reader
	Closer
}

type Reader interface {
	Read([]byte) (int, error)
}

type Closer interface {
	Close() error
}
