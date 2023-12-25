-- +goose Up
create table tenors (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
--     admin_id INT NOT NULL,
    tenor INT NOT NULL,
    `limit` INT NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
--     FOREIGN KEY (admin_id) REFERENCES users(id)
);

-- +goose Down
drop table tenors;
