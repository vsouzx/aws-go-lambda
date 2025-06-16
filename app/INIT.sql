
CREATE TABLE IF NOT EXISTS transactions (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    from_user_id INT NOT NULL,
    to_user_id INT NOT NULL,   
    amount DECIMAL(10, 2) NOT NULL,
    type VARCHAR(50) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    description TEXT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO transactions (from_user_id, to_user_id, amount, type, description)
VALUES (1, 2, 120.93, 'PIX', 'Pagando divida');

SELECT * 
FROM transactions
where from_user_id = :from_user_id;

UPDATE transactions
SET status = :status
WHERE id = :uuid;
