-- name: CreateAccount :one
INSERT INTO accounts (
    account_holder_name,
    balance,
    currency
) VALUES (
    $1, $2, $3
)
RETURNING id, account_holder_name, balance, currency, created_at;

-- name: GetAccount :one
SELECT id, account_holder_name, balance, currency, created_at
FROM accounts
WHERE id = $1;

-- name: ListAccounts :many
SELECT id, account_holder_name, balance, currency, created_at
FROM accounts
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateAccountBalance :one
UPDATE accounts
SET balance = $2
WHERE id = $1
RETURNING id, account_holder_name, balance, currency, created_at;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;
