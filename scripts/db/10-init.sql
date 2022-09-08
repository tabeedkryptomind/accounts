
CREATE DATABASE interview_accountapi;

\c interview_accountapi

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE  AccountAttributes (
    AccountClassification TEXT NOT NULL,    
    AccountMatchingOptOut BOOLEAN NOT NULL,
    AccountNumber TEXT NOT NULL, 
    AlternativeNames varchar[],
    BankID TEXT NOT NULL, 
    BankIDCode TEXT NOT NULL, 
    BaseCurrency TEXT NOT NULL, 
    Bic TEXT NOT NULL, 
    Country TEXT NOT NULL, 
    Iban TEXT NOT NULL, 
    JointAccount BOOLEAN NOT NULL,
    Name varchar[],
    SecondaryIdentification TEXT NOT NULL, 
    Status TEXT NOT NULL, 
    Switched BOOLEAN NOT NULL,

    PRIMARY KEY (Iban)
);

CREATE TABLE  AccountData    (
    ID  uuid DEFAULT uuid_generate_v4 (),
    organisationID  uuid DEFAULT uuid_generate_v4 (),
    type TEXT NOT NULL,
    version TEXT NOT NULL,
	Iban   TEXT NOT NULL, 

    PRIMARY KEY (ID),
    FOREIGN KEY (Iban) REFERENCES AccountAttributes(Iban) ON DELETE CASCADE
);

