-- name: CreateTransfer :one 
INSERT INTO transfers(
    from_account_id, 
    to_account_id, 
    amount
) VALUES (
    $1, $2, $3
)
RETURNING id, from_account_id, to_account_id, amount, created_at; 

-- name: GetTransfersOfAccount :many
SELECT id, from_account_id, to_account_id, amount, created_at
FROM transfers 
WHERE 
from_account_id = $1 
or 
to_account_id = $2 
ORDER by id 
LIMIT $3 
OFFSET $4; 

-- name: GetTransfer :one 
select id, from_account_id, to_account_id, amount, created_at
FROM transfers
WHERE id = $1 LIMIT 1; 



