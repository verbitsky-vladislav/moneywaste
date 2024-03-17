## Moneywaste - приложение для всестороннего учета финансов  

1. ### Базовые финансы
   1. _Транзакции_
      1. Доходы
      2. Расходы
   2. _Отчетность по финансам_
2. ### Инвестиции
   1. _Инвестпортфель (ведение)_
      1. _Заметки по инструментам_
   2. _Криптопортфель (ведение)_
      1. _Заметки по инструментам_
   3. _Отчетность по инвестициям_
3. ### Финансовые калькуляторы
   1. _Кредитный калькулятор_ 
   2. _Ипотечный калькулятор_ 
   3. _Кредитный калькулятор с досрочным погашением_ 
   4. _Калькулятор рефинансирования_ 
   5. _Калькулятор микрозаймов_ 
   6. _Калькулятор инфляции_
   7. _Калькулятор вкладов_ 
   8. _Калькулятор инвестиций_ 
   9. _Калькулятор доходности облигаций_
4. ### Финансовый блог
   1. _Посты_ 
   2. _Фото_ 
   3. _Коментарии_ 
   4. _Лайки_ 
5. ### Чат
   1. _Чат между двумя пользователями_


## TODO TASK MANAGER  
Расписать по модулям:
1. Описать функционал каждого модуля
2. Понять какие данные я храню (для каждого из модулей и их подмодулей)
3. Сделать схему базу данных
4. Написать логику
5. Написать хендлеры
6. Протестировать
7. Делать фронт

## Таблицы базы данных

1. User
2. 

"User": {
    "id": "uuid",
    "username": "string",
    "email": "string",
    "passwordHash": "string",
    "createdAt": "datetime",
    "updatedAt": "datetime"
}



"Transaction": {
    "id": "uuid",
    "userId": "uuid",
    "type": "enum['income', 'expense']",
    "amount": "decimal",
    "categoryId": "uuid",
    "date": "datetime",
    "description": "string",
    "createdAt": "datetime",
    "updatedAt": "datetime"
}

"TransactionCategory": {
    "id": "uuid",
    "name": "string",
    "type": "enum['income', 'expense']",
    "userId": "uuid"
}



"InvestmentPortfolio": {
    "id": "uuid",
    "userId": "uuid",
    "name": "string",
    "description": "string",
    "createdAt": "datetime",
    "updatedAt": "datetime"
}

"CryptoPortfolio": {
    "id": "uuid",
    "userId": "uuid",
    "name": "string",
    "description": "string",
    "createdAt": "datetime",
    "updatedAt": "datetime"
}

"InvestmentNote": {
    "id": "uuid",
    "portfolioId": "uuid",
    "note": "text",
    "createdAt": "datetime",
    "updatedAt": "datetime"
}

"FinancialCalculator": {
    "id": "uuid",
    "type": "enum['credit', 'mortgage', 'earlyRepayment', 'refinancing', 'microloan', 'inflation', 'deposit', 'investment', 'bondYield']",
    "parameters": "json",
    "userId": "uuid",
    "createdAt": "datetime",
    "updatedAt": "datetime"
}

"BlogPost": {
    "id": "uuid",
    "userId": "uuid",
    "title": "string",
    "content": "text",
    "createdAt": "datetime",
    "likes": "integer",
    "comments": [
        {
            "userId": "uuid",
            "comment": "text",
            "createdAt": "datetime"
        }
    ]
}

"ChatMessage": {
    "id": "uuid",
    "senderId": "uuid",
    "receiverId": "uuid",
    "message": "text",
    "sentAt": "datetime",
    "readAt": "datetime"
}