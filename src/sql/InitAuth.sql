use `auth`;

-- insert into dbConnection(connectionString,createdOn,createdBy) values
-- 	    ('WebClient@tcp(localhost:3306)/commish', CURRENT_TIMESTAMP, 'kyl:initLocal')
-- 	  , ('WebClient@tcp(localhost:3306)/ava', CURRENT_TIMESTAMP, 'kyl:initLocal');

-- insert into principal(dbConnectionId, createdOn, createdBy) values
--     (1, CURRENT_TIMESTAMP, 'kyl:initLocal')
--   , (2, CURRENT_TIMESTAMP, 'kyl:initLocal');

-- insert into principalLogin(principalId, identifier, contactTypeId, contactIdentifier, passwordHash, hashAlgorithm, createdOn, createdBy) values
--     (1, 'commish', 1, 'foo@bar.com', '$2a$10$NSudUrpR47hLfuuabwKXJOY9qXsy5nOITs6vqhWDS3/P6wJjDfUse', 'bcrypt', CURRENT_TIMESTAMP, 'kyl:test')
--   , (2, 'ava', 1, 'foo@bar.com', '$2a$10$LSmFQ0DDgMevM05VIRdyoOl9yKak6aZn3YmmpEJk7.4BhkEBAplfy', 'bcrypt', CURRENT_TIMESTAMP, 'kyl:test');
