-- name: CreateEntry :one // This is a named query 'CreateEntry' that returns a single row
INSERT INTO entries ( -- Insert a new row into 'entries' table
    account_id, -- Column to insert 'account_id'
    amount  -- Column to insert 'amount'
) VALUES (
             $1, $2  -- Placeholder values for 'account_id' and 'amount'
         ) RETURNING *; -- Returns all columns of the newly inserted row

-- name: GetEntry :one // This is a named query 'GetEntry' that returns a single row
SELECT * FROM entries
WHERE id = $1 -- Filter to match 'id' with a placeholder value
    LIMIT 1; -- Limit the result to only one row

-- name: ListEntries :many // This is a named query 'ListEntries' that can return multiple rows
SELECT * FROM entries
WHERE account_id = $1 -- Filter to match 'account_id' with a placeholder value
ORDER BY id -- Order the result by 'id'
    LIMIT $2 -- Limit the number of returned rows by a placeholder value
OFFSET $3; -- Skip a number of rows (offset) before returning results, specified by a placeholder