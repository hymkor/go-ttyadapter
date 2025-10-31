package ttyadapter

type Tty interface {
	Open(onSize func(int)) error
	GetKey() (string, error)
	Size() (int, int, error)
	Close() error
}
