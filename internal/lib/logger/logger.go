// Пакет logger реализует вспомогательные функции для работы slog логгера.
package logger

import "log/slog"

// Err возвращает описание ошибки, обернутую в формат, подходящий для вставки в логгер.
func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
