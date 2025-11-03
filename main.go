package ttyadapter

type Tty interface {
	Open(onResize func(int, int)) error
	GetKey() (string, error)
	Size() (int, int, error)
	Close() error
}
