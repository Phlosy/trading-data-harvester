package datamodel

type ValidatorTimeGap struct {
	Exchange    string
	Symbol      string
	Interval    string
	MissingFrom uint64
	MissingTo   uint64
}
