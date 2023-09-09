DROP TABLE IF EXISTS user;
CREATE TABLE IF NOT EXISTS user(
    id INTEGER PRIMARY KEY AUTOINCREMENT,  -- '自增ID'
    tenant_id INTEGER NOT NULL, --  '租户ID'
    role_id INTEGER DEFAULT NULL, -- '角色id'
    passport  varchar(45) NOT NULL unique, --  'User Passport'
    password  varchar(45) NOT NULL, --  'User Password'
    nickname  varchar(45) NOT NULL, --  'User Nickname'
    create_at datetime(0) DEFAULT NULL, --  'Created Time'
    update_at datetime(0) DEFAULT NULL --  'Updated Time'
);

REPLACE INTO user
(id, tenant_id,role_id, passport, password, nickname, create_at, update_at) 
VALUES 
(1, 1, 1, 'admin', '8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918', '管理员', datetime('now'), datetime('now')),
(2, 1, 2, 'guest', '84983c60f7daadc1cb8698621f802c0d9f9a3c3c295c810748fb048115c186ec', '租户管理员', datetime('now'), datetime('now')),
(3, 1, null, 'test', '84983c60f7daadc1cb8698621f802c0d9f9a3c3c295c810748fb048115c186ec', '张三', datetime('now'), datetime('now'));

DROP TABLE IF EXISTS role;
CREATE TABLE IF NOT EXISTS role(
    id INTEGER PRIMARY KEY AUTOINCREMENT,  -- '自增ID'
    tenant_id INTEGER NOT NULL, --  '租户ID'
    name varchar(45) NOT NULL, --  '角色名称'
    create_at datetime(0) DEFAULT NULL, --  'Created Time'
    update_at datetime(0) DEFAULT NULL --  'Updated Time'
);

REPLACE INTO role
(id, tenant_id, name, create_at, update_at) 
VALUES 
(1, 1, '超级管理员', datetime('now'), datetime('now')),
(2, 1, '租户管理员', datetime('now'), datetime('now')),
(3, 1, '普通员工', datetime('now'), datetime('now'));

DROP TABLE IF EXISTS role_permission;
CREATE TABLE IF NOT EXISTS role_permission(
    role_id INTEGER PRIMARY KEY, --  '角色ID'
    tenant_id INTEGER NOT NULL, --  '租户ID'
    permissions text NOT NULL, --  '权限列表'
    create_at datetime(0) DEFAULT NULL, --  'Created Time'
    update_at datetime(0) DEFAULT NULL --  'Updated Time'
);

REPLACE INTO role_permission
(role_id, tenant_id, permissions, create_at, update_at) 
VALUES 
(1,1, 'ChargeStationView,ChargeStationAdd,ChargeStationUpdate,ChargeStationDel,ChargePileView,ChargePileAdd,ChargePileUpdate,ChargePileDel,TenantView,TenantAdd,TenantUpdate,TenantDel,PriceView,PriceAdd,PriceDel,PriceUpdate,ChargeOrderView,AddPrice,AtuditLogView,PermissionView,PermissionViewBind,RoleView,RoleAdd,RoleDel,RoleUpdate,UserView,UserAdd,UserDel,UserUpdate,ScreenView', datetime('now'), datetime('now')),
(2,1, 'ChargeStationView,ChargeStationAdd,ChargeStationUpdate,ChargeStationDel,ChargePileView,ChargePileAdd,ChargePileUpdate,ChargePileDel,PriceView,PriceAdd,PriceDel,PriceUpdate,ChargeOrderView,AddPrice,AtuditLogView,PermissionView,PermissionViewBind,RoleView,RoleAdd,RoleDel,RoleUpdate,UserView,UserAdd,UserDel,UserUpdate,ScreenView', datetime('now'), datetime('now'));

DROP TABLE IF EXISTS tenant;
CREATE TABLE tenant (
    id INTEGER PRIMARY KEY AUTOINCREMENT, -- 租户ID，自增主键
    name varchar(255) NOT NULL, -- 租户名称
    domain varchar(255) UNIQUE, -- 租户域名（可选，如果您的应用支持自定义域名）
    contact_name varchar(255), -- 联系人姓名
    contact_email varchar(255), -- 联系人邮箱
    contact_phone varchar(20), -- 联系人电话
    status text NOT NULL DEFAULT 'active', -- 租户状态（如：激活、未激活、暂停）
    create_at datetime(0) DEFAULT NULL, --  '创建时间'
    update_at datetime(0) DEFAULT NULL --  '修改时间'
    CHECK (status IN ('active', 'inactive', 'suspended'))
);

REPLACE INTO tenant
(id, name, domain, contact_name, contact_email, contact_phone, status, create_at, update_at) 
VALUES 
(1, '土拨鼠充电公司', 'greenman.xx.com', '浩克','hunk@gmail.com','13544223322','active', datetime('now'), datetime('now')),
(2, '钢铁侠充电公司', 'steelman.xx.com', '托尼·斯塔克','tony@gmail.com','8888888','active', datetime('now'), datetime('now'));

DROP TABLE IF EXISTS pay_record;
CREATE TABLE IF NOT EXISTS pay_record(
    id INTEGER PRIMARY KEY AUTOINCREMENT,  -- '自增ID'
    tenant_id INTEGER NOT NULL, --  '租户ID'
    pay_code  varchar(45) NOT NULL UNIQUE, --  '订单号'
    user_id INTEGER NOT NULL, --  '用户Id'
    state INTEGER NOT NULL, --  '0:正在充值,1:充值成功,2:取消充值'
    price NUMERIC DEFAULT 0, --  ''
    pay_at datetime(0) DEFAULT NULL, --  '充值时间'
    create_at datetime(0) DEFAULT NULL, --  'Created Time'
    update_at datetime(0) DEFAULT NULL --  'Updated Time'
);

REPLACE INTO pay_record(id, tenant_id, pay_code,user_id, state, price,pay_at,create_at,update_at) VALUES
(1, 1, '202305111249-8745',1, 1, 10.5, datetime('now'),datetime('now'),datetime('now')),
(2, 1, '202306111250-5442',1, 1, 33.0, datetime('now'),datetime('now'),datetime('now')),
(3, 1, '202306231041-2341',1, 1, 12.2, datetime('now'),datetime('now'),datetime('now')),
(4, 1, '202306241832-2864',1, 1, 20.0, datetime('now'),datetime('now'),datetime('now'))
;

DROP TABLE IF EXISTS user_wallet;
CREATE TABLE IF NOT EXISTS user_wallet(
    user_id INTEGER PRIMARY KEY,  -- '自增ID'
    tenant_id INTEGER NOT NULL, --  '租户ID'
    balance  NUMERIC NOT NULL, --  '余额'
    create_at datetime(0) DEFAULT NULL, --  'Created Time'
    update_at datetime(0) DEFAULT NULL --  'Updated Time'
);
REPLACE INTO user_wallet
(user_id, tenant_id, balance,create_at,update_at) 
VALUES 
(1, 1, 10.0, datetime('now'),datetime('now')),
(2, 1, 50.0, datetime('now'),datetime('now')),
(3, 1, 50.0, datetime('now'),datetime('now'));

DROP TABLE IF EXISTS wx_user;
CREATE TABLE IF NOT EXISTS wx_user(
    user_id INTEGER PRIMARY KEY,  -- 'userID'
    tenant_id INTEGER NOT NULL, --  '租户ID'
    open_id  varchar(45) NOT NULL, --  '余额'
    phone_no  varchar(45) DEFAULT NULL, --  '手机号'
    avatar_url  varchar(45) DEFAULT NULL, --  '头像地址'
    nickname  varchar(45) DEFAULT NULL, --  '昵称'
    gender  varchar(45) DEFAULT NULL, --  '性别'
    create_at datetime(0) DEFAULT NULL, --  'Created Time'
    update_at datetime(0) DEFAULT NULL --  'Updated Time'
);

DROP TABLE IF EXISTS audit_log;
CREATE TABLE audit_log (
    id INTEGER PRIMARY KEY AUTOINCREMENT,  -- '自增ID'
    tenant_id INTEGER NOT NULL, --  '租户ID'
    operation_type varchar(20) NOT NULL,
    table_name varchar(50) NOT NULL,
    old_value varchar(512) DEFAULT NULL,
    new_value varchar(512) DEFAULT NULL,
    operate_sql varchar(512) NOT NULL,
    changed_date datetime(0) NOT NULL,
    changed_user varchar(50) NOT NULL
);
REPLACE INTO audit_log
(id, tenant_id, operation_type, table_name, old_value, new_value, operate_sql, changed_date, changed_user) 
VALUES 
(1, 1, "add", "user", "","{name:'张三'}","insert into user VALUES('xxxxx')",datetime('now'), "admin"),
(2, 1, "delete", "user", "{name:'张三'}","","delete user where id=1" ,datetime('now'), "zhangsan");