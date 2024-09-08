package storage

import (
	"path/filepath"
	"runtime"
)

func CurrentMigrationsPath(migrationsPath string) string {

	_, currenttPath, _, _ := runtime.Caller(0)
	currenttPath = filepath.Dir(currenttPath)
	result := filepath.Join(currenttPath, migrationsPath)
	return filepath.ToSlash(result)
	//TODO: Найти более простой способо сформировать путь к файлам миграций

}
