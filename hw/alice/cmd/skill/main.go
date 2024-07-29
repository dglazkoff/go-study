// пакеты исполняемых приложений должны называться main
package main

import (
	"database/sql"
	"go.uber.org/zap"
	"gopl.io/hw/alice/internal/logger"
	"gopl.io/hw/alice/internal/store/pg"
	"net/http"
	"strings"
)

func gzipMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// по умолчанию устанавливаем оригинальный http.ResponseWriter как тот,
		// который будем передавать следующей функции
		ow := w

		// проверяем, что клиент умеет получать от сервера сжатые данные в формате gzip
		acceptEncoding := r.Header.Get("Accept-Encoding")
		supportsGzip := strings.Contains(acceptEncoding, "gzip")
		if supportsGzip {
			// оборачиваем оригинальный http.ResponseWriter новым с поддержкой сжатия
			cw := newCompressWriter(w)
			// меняем оригинальный http.ResponseWriter на новый
			ow = cw
			// не забываем отправить клиенту все сжатые данные после завершения middleware
			defer cw.Close()
		}

		// проверяем, что клиент отправил серверу сжатые данные в формате gzip
		contentEncoding := r.Header.Get("Content-Encoding")
		sendsGzip := strings.Contains(contentEncoding, "gzip")
		if sendsGzip {
			// оборачиваем тело запроса в io.Reader с поддержкой декомпрессии
			cr, err := newCompressReader(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			// меняем тело запроса на новое
			r.Body = cr
			defer cr.Close()
		}

		// передаём управление хендлеру
		h.ServeHTTP(ow, r)
	}
}

// функция main вызывается автоматически при запуске приложения
func main() {
	parseFlags()

	if err := run(); err != nil {
		panic(err)
	}
}

// функция run будет полезна при инициализации зависимостей сервера перед запуском
func run() error {
	// это просто мое упражнение на работу с интерфейсами и mux
	//mux := CustomMux{routes: make(map[string]http.Handler)}
	//mux.HandleFunc("/", http.HandlerFunc(webhook))
	//
	//return http.ListenAndServe(":8080", mux)

	if err := logger.Initialize(flagLogLevel); err != nil {
		return err
	}

	// создаём соединение с СУБД PostgreSQL с помощью аргумента командной строки
	conn, err := sql.Open("pgx", flagDatabaseURI)
	if err != nil {
		return err
	}

	// создаём экземпляр приложения, передавая реализацию хранилища pg в качестве внешней зависимости
	appInstance := newApp(pg.NewStore(conn))

	logger.Log.Info("Running server", zap.String("address", flagRunAddr))
	// обернём хендлер webhook в middleware с логированием и поддержкой gzip
	return http.ListenAndServe(flagRunAddr, logger.RequestLogger(gzipMiddleware(appInstance.webhook)))
}

type CustomMux struct {
	routes map[string]http.Handler
}

func (m CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.routes[r.URL.Path].ServeHTTP(w, r)
}

func (m CustomMux) HandleFunc(path string, h http.Handler) {
	m.routes[path] = h
}