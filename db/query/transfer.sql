-- This names the SQL query 'CreateTransfer' and indicates it returns a single row
-- name: CreateTransfer :one
INSERT INTO transfers ( -- This command inserts a new row into the 'transfers' table
    from_account_id, -- Column for ID of the account transferring from
    to_account_id, -- Column for ID of the account transferring to
    amount -- Column for the amount being transferred
) VALUES (
             $1, $2, $3 -- Placeholder values for 'from_account_id', 'to_account_id', and 'amount'
         ) RETURNING *; -- Returns all columns of the newly inserted row

-- This names the SQL query 'GetTransfer' and indicates it returns a single row
-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 -- Filters the rows to find where the 'id' column matches the value of $1
LIMIT 1; -- Limits the query to return only one row

-- This names the SQL query 'ListTransfers' and indicates it returns multiple rows
-- name: ListTransfers :many
SELECT * FROM transfers
WHERE
    from_account_id = $1 OR -- Filters to include rows where 'from_account_id' matches $1
    to_account_id = $2 -- OR where 'to_account_id' matches $2
ORDER BY id -- Orders the results by the 'id' column
    LIMIT $3 -- Limits the number of rows returned to the value of $3
OFFSET $4; -- Skips the first $4 rows before beginning to return rows in the result set