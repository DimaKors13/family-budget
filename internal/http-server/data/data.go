// Пакет data содержит описание структур основных функциональных объектов приложения.
package data

import "time"

// Счет
// Описание структуры:
//	Id - Идентификатор текущего счета в БД
//	Name - Наименование счета
type Account struct {
	Id   int    `json:"id"`
	Name string `json:"name" validate:"required"`
}

// Категория доходов/расходов
// Описание структуры:
//	Id - Идентификатор текущей категории доходов/расходов в БД
//	Name - Наименование категории
//	Multiplier - Коэфициент движения ДС. Определяет принадлежность категории к записям доходов или расходов.
//	ParentId - Идентификатор категории-родителя в БД
type FlowCategory struct {
	Id         int
	Name       string
	Multiplier int
	ParentId   int
}

// Запись с отражением движения денежных средств
// Описание структуры:
//	Id - Идентификатор текущей записи в БД
//	Date - Дата движения ДС
//	Multiplier - Коэфициент движения ДС, определяющий отражение дохода или расхода
//	AccountId - Идентификатор счета в БД
//	CategoryId - Идентификатор категории доходов/расходов в БД
//	Amount - Сумма записи
type FinanceRecord struct {
	Id         int
	Date       *time.Time
	Multiplier int
	AccountId  int
	CategoryId int
	Amount     float64
}
