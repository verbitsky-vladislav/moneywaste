-- Создание функции, которая будет вызываться триггером для обновления updated_at
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;



CREATE TABLE "User" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    fio VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    passwordHash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE TRIGGER set_updated_at
    BEFORE UPDATE ON "User"
    FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();


CREATE TYPE transaction_type AS ENUM ('income', 'expense');
CREATE TABLE "TransactionCategory" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255)
);
CREATE TABLE "Transaction" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    userId UUID REFERENCES "User" (id) ON DELETE CASCADE,
    type transaction_type NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    categoryId UUID REFERENCES "TransactionCategory" (id) ON DELETE SET NULL,
    date DATE NOT NULL,
    description VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE TRIGGER set_updated_at
    BEFORE UPDATE ON "Transaction"
    FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();



-- "InvestmentPortfolio": {
--     "id": "uuid",
--     "userId": "uuid",
--     "name": "string",
--     "description": "string",
--     "createdAt": "datetime",
--     "updatedAt": "datetime"
-- }
--
-- "CryptoPortfolio": {
--     "id": "uuid",
--     "userId": "uuid",
--     "name": "string",
--     "description": "string",
--     "createdAt": "datetime",
--     "updatedAt": "datetime"
-- }
--
-- "InvestmentNote": {
--     "id": "uuid",
--     "portfolioId": "uuid",
--     "note": "text",
--     "createdAt": "datetime",
--     "updatedAt": "datetime"
-- }
--
-- "FinancialCalculator": {
--     "id": "uuid",
--     "type": "enum['credit', 'mortgage', 'earlyRepayment', 'refinancing', 'microloan', 'inflation', 'deposit', 'investment', 'bondYield']",
--     "parameters": "json",
--     "userId": "uuid",
--     "createdAt": "datetime",
--     "updatedAt": "datetime"
-- }
--
-- "BlogPost": {
--     "id": "uuid",
--     "userId": "uuid",
--     "title": "string",
--     "content": "text",
--     "createdAt": "datetime",
--     "likes": "integer",
--     "comments": [
--         {
--             "userId": "uuid",
--             "comment": "text",
--             "createdAt": "datetime"
--         }
--     ]
-- }
--
-- "ChatMessage": {
--     "id": "uuid",
--     "senderId": "uuid",
--     "receiverId": "uuid",
--     "message": "text",
--     "sentAt": "datetime",
--     "readAt": "datetime"
-- }