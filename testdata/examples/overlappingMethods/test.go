package overlappingMethods

type ReadCloser interface {
	Read([]byte) (int, error)
	Close() error
}

type WriteCloser interface {
	Write([]byte) (int, error)
	Close() error
}

//wrapgen:generate -template gomock -destination gomock.go
type ReadWriteCloser interface {
	ReadCloser
	WriteCloser
}
