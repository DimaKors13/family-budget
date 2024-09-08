package data

import "time"

type Account struct {
	Id   int
	Name string
}

type FlowCategory struct {
	Id         int
	Name       string
	Multiplier int
	ParentId   int
}

type FinanceRecord struct {
	Id         int
	Date       *time.Time
	Multiplier int
	AccountId  int
	CategoryId int
	Amount     float64
}
