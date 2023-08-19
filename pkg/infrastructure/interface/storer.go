package pkg

type Storer interface {
	Store(data string) error
	Clean(interval int) error
}
