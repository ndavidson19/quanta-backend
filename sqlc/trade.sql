-- name: CreateTrade :one
INSERT INTO trades (
  account_id,
  symbol,
  amount,
  price,
  trade_type,
  status,
  created_at,
  updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: GetTrade :one
SELECT * FROM trades
WHERE id = $1 LIMIT 1;

-- name: ListTrades :many
SELECT * FROM trades
ORDER BY account_id
LIMIT $1
OFFSET $2;
