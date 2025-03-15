SELECT
	CAST((SELECT value FROM config WHERE key = 'start_balance') AS NUMERIC)
	+
	SUM(CASE WHEN transaction_type = 2 THEN amount ELSE -amount END) AS current_balance
FROM
	transactions