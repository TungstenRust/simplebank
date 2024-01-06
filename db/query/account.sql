-- This names the SQL query 'CreateAccount' and indicates it returns a single row
-- name: CreateAccount :one
INSERT INTO accounts (
    owner, balance, currency
) VALUES (
    $1, $2, $3 -- Placeholder values for 'owner', 'balance', and 'currency'
) RETURNING *; -- Returns all columns of the newly inserted row

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 -- Filters the rows to find where the 'id' column matches the value of $1
LIMIT 1; -- Limits the query to return only one row

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id -- Orders the results by the 'id' column
LIMIT $1 -- Limits the number of rows returned to the value of $1
OFFSET $2; -- Skips the first $2 rows before beginning to return rows in the result set

-- This names the SQL query 'UpdateAccount' and indicates it is an execute command without returning rows
-- name: UpdateAccount :exec
UPDATE accounts -- Update 'balance' in 'accounts' table
SET balance = $2 -- Sets the 'balance' column to a new value specified by $2
WHERE id = $1; -- Filters to update the row where 'id' matches the value of $1

-- This names the SQL query 'DeleteAccount' and indicates it is an execute command without returning rows
-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1; -- Filters to delete the row where 'id' matches the value of $1