SELECT
	CAST((SELECT value FROM config WHERE key = 'start_balance') AS NUMERIC) as start_balance,
	COALESCE(SUM(CASE WHEN transaction_type = 2 THEN amount ELSE 0 END), 0) AS total_income,
    COALESCE(SUM(CASE WHEN transaction_type = 1 THEN -amount ELSE 0 END), 0) AS total_expense,
    COALESCE(SUM(CASE WHEN transaction_type = 2 THEN amount ELSE -amount END), 0) AS balance_delta,
    CAST((SELECT value FROM config WHERE key = 'start_balance') AS NUMERIC)
    +
    COALESCE(SUM(CASE WHEN transaction_type = 2 THEN amount ELSE -amount END), 0) AS current_balance
FROM
    transactions;