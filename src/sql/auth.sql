drop schema if exists `auth`;

CREATE SCHEMA if not exists `auth` ;

use `auth`;

create table if not exists contactType(
	id int primary key not null auto_increment,
    name nvarchar(32) not null,
    createdOn datetime not null default CURRENT_TIMESTAMP
);

insert into contactType(name) values
	('email');
    
create table if not exists dbConnection(
	id int primary key not null auto_increment,
    publicId binary(16) unique,
    active bit not null default 1,
    connectionString nvarchar(256) not null,
    createdOn datetime not null,
    createdBy nvarchar(32) not null,
    modifiedOn datetime null,
    modifiedBy nvarchar(32) null
);

create trigger before_insert_dbConnection
	before insert on dbConnection
    for each row
    set new.publicId = UNHEX(REPLACE(UUID(), '-', ''));
    

create table if not exists principal(
	id int primary key not null auto_increment,
    publicId binary(16) unique,
    active bit not null default 1,
    tokenValidForHours int not null default 24,
    dbConnectionId int not null, foreign key(dbConnectionId) references dbConnection(id),
    createdOn datetime not null,
    createdBy nvarchar(32) not null,
    modifiedOn datetime null,
    modifiedBy nvarchar(32) null
);

create trigger before_insert_principal
	before insert on principal
    for each row
    set new.publicId = UNHEX(REPLACE(UUID(), '-', ''));
    
create table if not exists principalLogin(
	id int primary key not null auto_increment,
    principalId int not null, foreign key(principalId) references principal(id),
    identifier nvarchar(128) not null,
	contactTypeId int not null, foreign key(contactTypeId) references contactType(id),
    contactIdentifier nvarchar(128) not null,
    passwordHash nvarchar(128) null,
    verifyTokenHash nvarchar(128) null,
    hashAlgorithm nvarchar(32) not null,
    active bit not null default 1,
    createdOn datetime not null,
    createdBy nvarchar(32) not null,
    modifiedOn datetime null,
    modifiedBy datetime null
);