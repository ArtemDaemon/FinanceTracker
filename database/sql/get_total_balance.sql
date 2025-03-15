SELECT
    CAST((SELECT value FROM config WHERE key = 'start_balance') AS NUMERIC)
    +
    COALESCE(SUM(CASE WHEN transaction_type = 2 THEN amount ELSE -amount END), 0) AS current_balance
FROM
    transactions;