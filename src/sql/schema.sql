drop schema if exists `commish`;

CREATE SCHEMA if not exists `commish` ;

use `commish`;

create table if not exists genderType(
	id int primary key not null auto_increment,
    name nvarchar(32) not null,
    createdOn datetime not null default CURRENT_TIMESTAMP
);

insert into genderType(name) values
	('female')
  , ('male')
  , ('co-ed')
  , ('other')
  , ('unspecified');
  
create table if not exists contactType(
	id int primary key not null auto_increment,
    name nvarchar(32) not null,
    createdOn datetime not null default CURRENT_TIMESTAMP
);

insert into contactType(name) values
	('email');
    
create table if not exists gameStateType(
	id int primary key not null auto_increment,
    name nvarchar(32) not null,
    createdOn datetime not null default CURRENT_TIMESTAMP
);

insert into gameStateType(name) values
	('pending')
  , ('inProgress')
  , ('completed');
    
create table if not exists tournamentFormatType(
	id int primary key not null auto_increment,
    name nvarchar(32) not null,
    createdOn datetime not null default CURRENT_TIMESTAMP
);

insert into tournamentFormatType(name) values
	('singleElimination')
  , ('doubleElimination')
  , ('doubleEliminationWithCrossover');
  
create table if not exists leagueTournamentRelationType(
	id int primary key not null auto_increment,
    name nvarchar(32) not null,
    createdOn datetime not null default CURRENT_TIMESTAMP
);

insert into leagueTournamentRelationType(name) values
	('championshipTournament');

create table if not exists principal(
	id int primary key not null auto_increment,
    publicId binary(16) unique,
    active bit not null default 1,
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

create table if not exists role(
	id int primary key not null auto_increment,
    name nvarchar(32) not null unique,
    description nvarchar(64) not null,
    active bit not null default 1,
	createdOn datetime not null,
    createdBy nvarchar(32) not null,
    modifiedOn datetime null,
    modifiedBy nvarchar(32) null
);

create table if not exists claim(
	id int primary key not null auto_increment,
    name nvarchar(32) not null,
    value nvarchar(32) not null,
    description nvarchar(64) not null,
    active bit not null default 1,
	createdOn datetime not null,
    createdBy nvarchar(32) not null,
    modifiedOn datetime null,
    modifiedBy nvarchar(32) null
);

create table if not exists rolePrincipal(
	id int primary key not null auto_increment,
    roleId int not null, foreign key(roleId) references role(id),
    principalId int not null, foreign key(principalId) references principal(id),
    expiresOn datetime null,
	createdOn datetime not null,
    createdBy nvarchar(32) not null,
    modifiedOn datetime null,
    modifiedBy nvarchar(32) null
);

create table if not exists roleClaim(
	id int primary key not null auto_increment,
    roleId int not null, foreign key(roleId) references role(id),
    claimId int not null, foreign key(claimId) references claim(id),
	expiresOn datetime null,
	createdOn datetime not null,
    createdBy nvarchar(32) not null,
    modifiedOn datetime null,
    modifiedBy nvarchar(32) null
);
    
create table if not exists player(
	id int primary key not null auto_increment,
    publicId binary(16) unique,
    principalId int null, foreign key(principalId) references principal(id),
    name nvarchar(128) not null,
    age int null,
    genderId int not null, foreign key(genderId) references genderType(id),
    weight int null,
    height int null,
    createdOn datetime not null,
    createdBy nvarchar(32) not null,
    modifiedOn datetime not null,
    modifiedBy nvarchar(32) null
);

create trigger before_insert_player 
	before insert on player
    for each row
    set new.publicId = UNHEX(REPLACE(UUID(), '-', ''));
    
create table if not exists team(
	id int primary key not null auto_increment,
    publicId binary(16) unique,
    name nvarchar(64) not null,
	createdOn datetime not null,
    createdBy nvarchar(32) not null,
    modifiedOn datetime not null,
    modifiedBy nvarchar(32) null
);

create trigger before_insert_team
	before insert on team
    for each row
    set new.publicId = UNHEX(REPLACE(UUID(), '-', ''));
    
create table if not exists playerTeam(
	id int primary key not null auto_increment,
    playerId int not null, foreign key(playerId) references player(id),
    teamId int not null, foreign key(teamId) references team(id),
	createdOn datetime not null,
    createdBy nvarchar(32) not null,
    modifiedOn datetime not null,
    modifiedBy nvarchar(32) null
);

create table if not exists division(
	id int primary key not null auto_increment,
    publicId binary(16) unique,
    name nvarchar(64) not null,
	createdOn datetime not null,
    createdBy nvarchar(32) not null,
    modifiedOn datetime not null,
    modifiedBy nvarchar(32) null
);

create trigger before_insert_division
	before insert on division
    for each row
    set new.publicId = UNHEX(REPLACE(UUID(), '-', ''));

create table if not exists tournament(
	id int primary key not null auto_increment,
    publicId binary(16) unique,
    name nvarchar(64) not null,
    location nvarchar(64) null,
    description nvarchar(128) null,
    divisionId int not null, foreign key(divisionId) references division(id),
    genderId int not null, foreign key(genderId) references genderType(id),
    startDate datetime not null,
    endDate datetime not null,
	createdOn datetime not null,
    createdBy nvarchar(32) not null,
    modifiedOn datetime not null,
    modifiedBy nvarchar(32) null
);

create trigger before_insert_tournament
	before insert on tournament
    for each row
    set new.publicId = UNHEX(REPLACE(UUID(), '-', ''));
    
create table if not exists league(
	id int primary key not null auto_increment,
    publicId binary(16) unique,
    name nvarchar(64) not null,
    location nvarchar(64) null,
    description nvarchar(128) null,
    divisionId int not null, foreign key(divisionId) references division(id),
    genderId int not null, foreign key(genderId) references genderType(id),
    startDate datetime not null,
    endDate datetime not null,
	createdOn datetime not null,
    createdBy nvarchar(32) not null,
    modifiedOn datetime not null,
    modifiedBy nvarchar(32) null
);

create trigger before_insert_league
	before insert on league
    for each row
    set new.publicId = UNHEX(REPLACE(UUID(), '-', ''));
    
create table if not exists leagueTournament(
	id int primary key not null auto_increment,
    leagueId int not null, foreign key(leagueId) references league(id),
    tournamentId int not null, foreign key(tournamentId) references tournament(id),
	createdOn datetime not null,
    createdBy nvarchar(32) not null,
    modifiedOn datetime not null,
    modifiedBy nvarchar(32) null
);

create table if not exists game(
	id int primary key not null auto_increment,
    publicId binary(16) unique,
    homeTeamId int not null, foreign key(homeTeamId) references team(id),
    awayTeamId int not null, foreign key(awayTeamId) references team(id),
    stateId int not null, foreign key(stateId) references gameStateType(id),
    homeTeamScore int not null default 0,
    awayTeamScore int not null default 0,
    leagueId int null, foreign key(leagueId) references league(id),
    tournamentId int null, foreign key(tournamentId) references tournament(id),
	createdOn datetime not null,
    createdBy nvarchar(32) not null,
    modifiedOn datetime not null,
    modifiedBy nvarchar(32) null
);

create trigger before_insert_game
	before insert on game
    for each row
    set new.publicId = UNHEX(REPLACE(UUID(), '-', ''));
    
create table if not exists statisticCalculator(
	id int primary key not null auto_increment,
    publicId binary(16) unique,
    name nvarchar(32) not null,
    description nvarchar(128) null,
    createdOn datetime not null,
    createdBy nvarchar(32) not null,
    modifiedOn datetime not null,
    modifiedBy nvarchar(32) null
);

create trigger before_insert_statisticCalculator
	before insert on statisticCalculator
    for each row
    set new.publicId = UNHEX(REPLACE(UUID(), '-', ''));
    
-- Immutable, de-normalized table for general time-series stats
create table if not exists statistic(
	id int primary key not null auto_increment,
    statisticCalculatorId int not null, foreign key(statisticCalculatorId) references statisticCalculator(id),
    value decimal not null,
    timeReported datetime not null,
    dataHash binary(16) not null,
    playerId int null, foreign key(playerId) references player(id),
    teamId int null, foreign key(teamId) references team(id),
    gameId int null, foreign key(gameId) references game(id),
    tournamentId int null, foreign key(tournamentId) references tournament(id),
    leagueId int null, foreign key(leagueId) references league(id),
    createdOn datetime not null,
    createdBy nvarchar(32) not null
);
 
-- Mutable, normalized table for human-entered stats
create table if not exists mutableStatistic(
	id int primary key not null auto_increment,
    statisticCalculatorId int not null, foreign key(statisticCalculatorId) references statisticCalculator(id),
    value decimal not null,
    timeReported datetime not null,
    playerId int null, foreign key(playerId) references player(id),
    teamId int null, foreign key(teamId) references team(id),
    gameId int null, foreign key(gameId) references game(id),
    tournamentId int null, foreign key(tournamentId) references tournament(id),
    leagueId int null, foreign key(leagueId) references league(id),
    createdOn datetime not null,
    createdBy nvarchar(32) not null,
    modifiedOn datetime null,
    modifiedBy nvarchar(32) not null
);

