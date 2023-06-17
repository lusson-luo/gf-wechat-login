-- DROP TABLE IF EXISTS user;
CREATE TABLE IF NOT EXISTS user(
    id INTEGER PRIMARY KEY AUTOINCREMENT,  -- '自增ID'
    passport  varchar(45) NOT NULL unique, --  'User Passport'
    password  varchar(45) NOT NULL, --  'User Password'
    nickname  varchar(45) NOT NULL, --  'User Nickname'
    create_at datetime(0) DEFAULT NULL, --  'Created Time'
    update_at datetime(0) DEFAULT NULL --  'Updated Time'
);

REPLACE INTO user
(id, passport, password, nickname, create_at, update_at) 
VALUES 
(1, 'admin', '8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918', '管理员', datetime('now'), datetime('now')),
(2, 'guest', '84983c60f7daadc1cb8698621f802c0d9f9a3c3c295c810748fb048115c186ec', '访客', datetime('now'), datetime('now'));

-- DROP TABLE IF EXISTS wx_user;
CREATE TABLE IF NOT EXISTS wx_user(
    user_id INTEGER PRIMARY KEY,  -- 'userID'
    open_id  varchar(45) NOT NULL, --  '余额'
    phone_no  varchar(45) DEFAULT NULL, --  '手机号'
    avatar_url  varchar(45) DEFAULT NULL, --  '头像地址'
    nickname  varchar(45) DEFAULT NULL, --  '昵称'
    gender  varchar(45) DEFAULT NULL, --  '性别'
    create_at datetime(0) DEFAULT NULL, --  'Created Time'
    update_at datetime(0) DEFAULT NULL --  'Updated Time'
);
