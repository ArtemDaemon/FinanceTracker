CREATE TABLE IF NOT EXISTS categories (
	category_id INTEGER PRIMARY KEY AUTOINCREMENT,
	category_name TEXT NOT NULL UNIQUE
);

INSERT OR IGNORE INTO categories (category_name) 
VALUES ('Продукты');

CREATE TABLE IF NOT EXISTS config (
	key	TEXT NOT NULL UNIQUE,
	value	TEXT,
	PRIMARY KEY(key)
);

INSERT OR IGNORE INTO config (key, value) VALUES ('start_balance', '0');

CREATE TABLE IF NOT EXISTS transactions (
	transaction_id	INTEGER NOT NULL UNIQUE,
	category_id	INTEGER NOT NULL,
	amount	REAL NOT NULL DEFAULT 0,
	transaction_type	INTEGER NOT NULL DEFAULT 1,
	"ate"TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
	PRIMARY KEY(transaction_id AUTOINCREMENT),
	CONSTRAINT category_fk FOREIGN KEY(category_id) REFERENCES categories(category_id)
);