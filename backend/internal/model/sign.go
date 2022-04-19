package model

type Sign struct {
	ID       int
	Status   bool
	FromWho  int
	ToWhom   int
	SignId   int
	Document // doc id
}
