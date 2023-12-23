-- +goose Up
create table assets (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    contract_number INT NOT NULL,
    on_the_road INT NOT NULL,
    admin_fee INT NOT NULL,
    installment INT NOT NULL,
    interest FLOAT NOT NULL,
    asset_name VARCHAR(20) NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- +goose Down
drop table assets;