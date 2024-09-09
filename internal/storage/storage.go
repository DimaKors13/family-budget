// Пакет storage реализует общие функции для работы с любым storage.
package storage

import (
	"path/filepath"
	"runtime"
)

// CurrentMigrationsPath возвращает полный путь к каталогу migrationsPath
func CurrentMigrationsPath(migrationsPath string) string {

	_, currenttPath, _, _ := runtime.Caller(0)
	currenttPath = filepath.Dir(currenttPath)
	result := filepath.Join(currenttPath, migrationsPath)
	return filepath.ToSlash(result)

}
