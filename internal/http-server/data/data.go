package data

import "time"

type Account struct {
	Id   int    `json:"id"`
	Name string `json:"name" validate:"required"`
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
