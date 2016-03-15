use `auth`;

insert into dbConnection(connectionString,createdOn,createdBy) values
    ('WebClient@tcp(localhost:3306)/commish', CURRENT_TIMESTAMP, 'kyl:initLocal')
  , ('WebClient@tcp(localhost:3306)/ava', CURRENT_TIMESTAMP, 'kyl:initLocal');

insert into principal(dbConnectionId, createdOn, createdBy) values
    (1, CURRENT_TIMESTAMP, 'kyl:initLocal')
  , (2, CURRENT_TIMESTAMP, 'kyl:initLocal');

-- commish: commish
-- ava: ava
insert into principalLogin(principalId, identifier, contactTypeId, contactIdentifier, passwordHash, hashAlgorithm, createdOn, createdBy) values
    (1, 'commish', 1, 'foo@bar.com', '$2a$10$morCA8HeImT3CbIXpkbgvuq3KisU3ZwC5N2sMuuuR1813tVLZ5DaC', 'bcrypt', CURRENT_TIMESTAMP, 'kyl:test')
  , (2, 'ava', 1, 'foo@bar.com', '$2a$10$KDN/hBexBTN.DBJ1g7b.TeYzLlmMx1pL0Ytb2wn2e9foZc657hP3a', 'bcrypt', CURRENT_TIMESTAMP, 'kyl:test');
