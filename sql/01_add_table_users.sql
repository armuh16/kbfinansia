-- +goose Up
create table users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name      VARCHAR(255) NOT NULL,
    password  VARCHAR(255) NOT NULL,
    email     VARCHAR(255) UNIQUE NOT NULL,
    role      INT NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

insert into users(name, password, email, role) values ('Admin','$2a$12$yT.dJTZnu4FRJq9zXw0mBOA/xmZHJPVi5ni13Zk9Pn6E0QmwKkZTu','admin@mail.com',1);
insert into users(name, password, email, role) values ('Budi','$2a$12$yT.dJTZnu4FRJq9zXw0mBOA/xmZHJPVi5ni13Zk9Pn6E0QmwKkZTu','user1@mail.com',2);
insert into users(name, password, email, role) values ('Annisa','$2a$12$yT.dJTZnu4FRJq9zXw0mBOA/xmZHJPVi5ni13Zk9Pn6E0QmwKkZTu','user2@mail.com',2);

-- +goose Down
drop table users;