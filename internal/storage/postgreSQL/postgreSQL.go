// Пакет postgreSQL реализует функции для работы с БД PostgreSQL.
package postgreSQL

import (
	"database/sql"
	"errors"
	"family-budget/internal/config"
	"family-budget/internal/http-server/data"
	"family-budget/internal/storage"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

//TODO: create DELETE methods of storage

// psqlMigrationsPath содержит локальный путь к папке, содержащей файлы миграции БД.
const psqlMigrationsPath = "./postgreSQL/migrations/"

type Storage struct {
	db *sql.DB
}

// New возвращает соединение с БД PosgreSQL.
//
// Инициализация соединения с БД происходит в несколько этапов:
//  1. Формирование пути подключения к БД и выполнение соединения с БД.
//  2. Сборка драйвера миграции БД. Выполнение миграций БД.
func New(AccessInfo *config.DBAccessInfo) (*Storage, error) {

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s", AccessInfo.DBUser, AccessInfo.DBPassword, AccessInfo.DBAddress)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("could not initialize SQL driver: %w", err)
	}

	migrationsPath := storage.CurrentMigrationsPath(psqlMigrationsPath)
	migration, err := migrate.NewWithDatabaseInstance("file://"+migrationsPath, "postgres", driver)
	if err != nil {
		return nil, fmt.Errorf("could not start migration: %w", err)
	}

	err = migration.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	return &Storage{db: db}, nil
}

// AddAccount добавляет новый Account в БД и  возвращает индекс добавленного элемента
func (storage *Storage) AddAccount(name string) (int, error) {

	query := "INSERT INTO accounts (name) VALUES ($1) RETURNING id"
	row := storage.db.QueryRow(query, name)
	var index int
	err := row.Scan(&index)
	if err != nil {
		return 0, fmt.Errorf("failed to insert into accounts: %w", err)
	}

	return index, nil
}

// GetAccount получает данные об Account из БД по id и возвращает Account, заполненный по данным БД
func (storage *Storage) GetAccount(accountId int) (*data.Account, error) {

	result := data.Account{}
	query := "SELECT id, name FROM accounts WHERE id = $1;"
	row := storage.db.QueryRow(query, sql.Named("account_id", accountId))
	err := row.Scan(&result.Id, &result.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get account by id: %w", err)
	}

	return &result, nil
}

// AddFlowCategory добавляет новый FlowCategory в БД и  возвращает индекс добавленного элемента
func (storage *Storage) AddFlowCategory(name string, multiplier, parentId int) (int, error) {

	queryParams := fmt.Sprintf("'%s', %d, parentID", name, multiplier)

	if parentId == 0 {
		queryParams = strings.Replace(queryParams, "parentID", "NULL", 1)
	} else {
		queryParams = strings.Replace(queryParams, "parentID", strconv.Itoa(parentId), 1)
	}

	query := "INSERT INTO flow_categories(name, multiplier, parent_id) VALUES (%s) RETURNING id;"
	row := storage.db.QueryRow(fmt.Sprintf(query, queryParams))

	var index int
	err := row.Scan(&index)
	if err != nil {
		return 0, fmt.Errorf("failed to insert into flow_categories: %w", err)
	}

	return index, nil
}

// GetFlowCategory получает данные о FlowCategory из БД по id и возвращает FlowCategory, заполненный по данным БД
func (storage *Storage) GetFlowCategory(categoryId int) (*data.FlowCategory, error) {

	query := queryGetFlowCategoryById()
	row := storage.db.QueryRow(query, categoryId)

	category := data.FlowCategory{}
	err := row.Scan(&category.Id, &category.Name, &category.Multiplier, &category.ParentId)
	if err != nil {
		return nil, fmt.Errorf("failed to get flow categoty by id: %w", err)
	}

	return &category, nil
}

// AddRecordToCashFlow добавляет новый FinanceRecord в БД и  возвращает индекс добавленного элемента
func (storage *Storage) AddRecordToCashFlow(record *data.FinanceRecord) error {

	query := "INSERT INTO cash_flow(record_date, multiplier, category_id, account_id, amount) VALUES ($1, $2, $3, $4, $5) RETURNING id;"
	row := storage.db.QueryRow(query, record.Date.Format(time.DateTime), record.Multiplier, record.CategoryId, record.AccountId, record.Amount)

	err := row.Scan(&record.Id)
	if err != nil {
		return fmt.Errorf("failed to insert into cash flow: %w", err)
	}

	return nil

}

// GetRecordFromCashFlow получает данные о FinanceRecord из БД по id и возвращает FinanceRecord, заполненный по данным БД
func (storage *Storage) GetRecordFromCashFlow(recordId int) (*data.FinanceRecord, error) {

	query := "SELECT id, record_date, multiplier, category_id, account_id, amount FROM public.cash_flow WHERE id = $1;"
	row := storage.db.QueryRow(query, recordId)

	result := data.FinanceRecord{}
	err := row.Scan(&result.Id, &result.Date, &result.Multiplier, &result.CategoryId, &result.AccountId, &result.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to get record from cash_flow by id: %w", err)
	}

	return &result, nil
}

// queryGetFlowCategoryById возвращает текст запроса к БД для получения данных FlowCategory по id
func queryGetFlowCategoryById() string {

	return (`SELECT
		ID,
		NAME,
		MULTIPLIER,
		CASE
			WHEN PARENT_ID IS NULL THEN 0
			ELSE PARENT_ID
		END
	FROM
		FLOW_CATEGORIES WHERE
		ID = $1;`)
}
