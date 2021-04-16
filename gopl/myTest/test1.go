package myTest

type TestStr struct {
	Num 	int
	Name 	string
	size	int
}

func (t *TestStr) SetSize(size int) {
	t.size = 10
}

func (t *TestStr) GetSize() int {
	return t.size
}