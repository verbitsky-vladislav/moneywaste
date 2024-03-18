package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"moneywaste/repository/models"
)

type TransactionsCategory struct {
	db *sql.DB
}

func NewTransactionsCategory(db *sql.DB) *TransactionsCategory {
	return &TransactionsCategory{
		db: db,
	}
}

// Init
// Creating all main categories
func (s *TransactionsCategory) Init() {
	var categories = []models.TransactionCategory{
		{Name: "Продукты", Description: "Покупка продуктов питания"},
		{Name: "Транспорт", Description: "Расходы на общественный транспорт и такси"},
		{Name: "Жилье", Description: "Оплата аренды, ипотеки, коммунальных услуг"},
		{Name: "Развлечения", Description: "Кино, театры, концерты и другие мероприятия"},
		{Name: "Здоровье", Description: "Расходы на лекарства, визиты к врачам"},
		{Name: "Одежда", Description: "Покупка одежды и обуви"},
		{Name: "Кафе и рестораны", Description: "Питание вне дома"},
		{Name: "Подарки", Description: "Подарки друзьям и близким"},
		{Name: "Техника", Description: "Покупка бытовой техники, электроники"},
		{Name: "Спорт", Description: "Расходы на спортзал, тренажерный зал, фитнес"},
		{Name: "Образование", Description: "Курсы, обучающие программы, учебники"},
		{Name: "Дом и сад", Description: "Инструменты, садовая техника, ремонт и улучшение жилища"},
		{Name: "Путешествия", Description: "Расходы на отпуск, билеты, гостиницы"},
		{Name: "Сбережения", Description: "Отчисления на личные или семейные сбережения"},
		{Name: "Инвестиции", Description: "Покупка акций, облигаций, криптовалюты"},
		{Name: "Питомцы", Description: "Товары для животных, ветеринарные услуги"},
		{Name: "Красота и уход", Description: "Салоны красоты, косметика, уходовая продукция"},
		{Name: "Ремонт автомобиля", Description: "Обслуживание и ремонт транспортных средств"},
		{Name: "Страхование", Description: "Оплата страховых взносов"},
		{Name: "Благотворительность", Description: "Пожертвования на благотворительные и социальные нужды"},
		{Name: "Развитие бизнеса", Description: "Расходы на развитие собственного дела"},
		{Name: "Услуги связи", Description: "Мобильная связь, интернет"},
		{Name: "Культура", Description: "Посещение музеев, выставок, библиотек"},
		{Name: "Игрушки и игры", Description: "Видеоигры, настольные игры, игрушки для детей"},
		{Name: "Хобби", Description: "Материалы и принадлежности для хобби"},
	}

	// Демонстрация списка категорий
	for _, category := range categories {
		_, err := s.Create(category)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func (s *TransactionsCategory) Create(transactionCategory models.TransactionCategory) (*models.TransactionCategory, error) {
	var category models.TransactionCategory

	err := s.db.QueryRow(`SELECT id, name, description FROM "TransactionCategory" WHERE name = $1`, transactionCategory.Name).Scan(&category.Id, &category.Name, &category.Description)
	if err == nil {
		return &category, nil
	} else if !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	err = s.db.QueryRow(`INSERT INTO "TransactionCategory" (name, description) VALUES ($1, $2) RETURNING id, name, description`,
		transactionCategory.Name, transactionCategory.Description).Scan(&category.Id, &category.Name, &category.Description)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (s *TransactionsCategory) Update() {

}

func (s *TransactionsCategory) Delete() {

}

func (s *TransactionsCategory) GetOneById(id string) (*models.TransactionCategory, error) {
	var transaction models.TransactionCategory

	query := fmt.Sprintf(`SELECT * FROM "TransactionCategory" WHERE id = '%s'`, id)

	rows, err := s.db.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(rows)
	for rows.Next() {
		err := rows.Scan(&transaction.Id, &transaction.Name, &transaction.Description)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &transaction, nil
}

func (s *TransactionsCategory) GetAll() {

}
