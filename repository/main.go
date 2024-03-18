package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

var once sync.Once
var instance *sql.DB

// InitDB инициализирует подключение к базе данных и сохраняет его как Singleton.
func InitDB(dataSourceName string) {
	once.Do(func() {
		var err error
		instance, err = sql.Open("postgres", dataSourceName)
		if err != nil {
			log.Fatalf("Не удалось подключиться к базе данных: %v", err)
		}

		// Проверьте соединение
		if err = instance.Ping(); err != nil {
			log.Fatalf("Не удалось выполнить ping базы данных: %v", err)
		}

		log.Println("Подключение к базе данных успешно установлено")
	})

	NewTransactionsCategory(instance).Init()
}

// GetDB возвращает экземпляр подключения к базе данных.
func GetDB() *sql.DB {
	if instance == nil {
		log.Fatal("Инициализация подключения к базе данных не выполнена")
	}
	return instance
}
