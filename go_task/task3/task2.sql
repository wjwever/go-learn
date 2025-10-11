use task3;

drop table if exists transactions;
drop table if exists accounts;
DROP PROCEDURE IF EXISTS transfer_money;

-- 创建 accounts 表
CREATE TABLE accounts (
    id INT PRIMARY KEY AUTO_INCREMENT,
    balance DECIMAL(10, 2) NOT NULL DEFAULT 0.00
);

-- 创建 transactions 表
CREATE TABLE transactions (
    id INT PRIMARY KEY AUTO_INCREMENT,
    from_account_id INT,
    to_account_id INT,
    amount DECIMAL(10, 2) NOT NULL,
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (from_account_id) REFERENCES accounts(id),
    FOREIGN KEY (to_account_id) REFERENCES accounts(id)
);

-- 插入示例数据
INSERT INTO accounts (id, balance) VALUES (1, 500.00); -- 账户 A (ID=1) 有 500 元
INSERT INTO accounts (id, balance) VALUES (2, 200.00); -- 账户 B (ID=2) 有 200 元

DELIMITER $$

CREATE PROCEDURE transfer_money(
    IN from_acc_id INT,
    IN to_acc_id INT,
    IN transfer_amount DECIMAL(10, 2)
)
BEGIN
    -- 声明一个变量来存储转出账户的余额
    DECLARE from_balance DECIMAL(10, 2);

    -- 开始事务
    START TRANSACTION;

    -- 锁定要操作的行以防止并发问题，并获取转出账户的余额
    SELECT balance INTO from_balance FROM accounts WHERE id = from_acc_id FOR UPDATE;

    -- 检查余额是否充足
    IF from_balance >= transfer_amount THEN
        -- 从账户 A 扣款
        UPDATE accounts SET balance = balance - transfer_amount WHERE id = from_acc_id;

        -- 向账户 B 存款
        UPDATE accounts SET balance = balance + transfer_amount WHERE id = to_acc_id;

        -- 在 transactions 表中记录该笔转账
        INSERT INTO transactions (from_account_id, to_account_id, amount)
        VALUES (from_acc_id, to_acc_id, transfer_amount);

        -- 如果一切顺利，提交事务
        COMMIT;
        SELECT '转账成功！' AS status;
    ELSE
        -- 如果余额不足，则回滚事务
        ROLLBACK;
        SELECT '余额不足，转账失败！' AS status;
    END IF;

END$$

DELIMITER ;

CALL transfer_money(1, 2, 200.00);

