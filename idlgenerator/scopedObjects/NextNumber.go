package scopedObjects

type NextNumber struct {
	Number int
}

func NewNextNumber() *NextNumber {
	return &NextNumber{
		Number: 0,
	}
}

func (self *NextNumber) NextNumber() int {
	self.Number++
	return self.Number
}
