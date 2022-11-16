CREATE TABLE user_balance (
                              id int NOT NULL UNIQUE,
                              balance int
);
CREATE TABLE reserve_account (
                                 id int NOT NULL,
                                 service int NOT NULL,
                                 order_id int NOT NULL UNIQUE ,
                                 cost int
);
CREATE TABLE accounting_report (
                                   service int NOT NULL,
                                   cost int,
                                   order_date date NOT NULL
);