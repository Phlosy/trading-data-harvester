package datamodel

type ValidatorTimeGap struct {
	Symbol      string
	Interval    string
	MissingFrom uint64
	MissingTo   uint64
}
