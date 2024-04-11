CREATE TABLE Blogs(id UUID DEFAULT uuid_generate_v4(),userid UUID,title varchar(255) NOT NULL,content TEXT[],category varchar(255) NOT NULL,created_at timestamp NOT NULL,updated_at timestamp NOT NULL, PRIMARY KEY(id),CONSTRAINT fk_userid FOREIGN KEY (userid) REFERENCES Users(id));

CREATE TABLE Users (id UUID DEFAULT uuid_generate_v4() ,name VARCHAR(255) NOT NULL,username VARCHAR(255) NOT NULL UNIQUE,password VARCHAR(255) NOT NULL,refresh_token VARCHAR(255),category TEXT[], PRIMARY KEY(id));

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE Categories (id UUID DEFAULT uuid_generate_v4(), name VARCHAR(255) UNIQUE, frequency INT, PRIMARY KEY (id));

ALTER TABLE users ADD COLUMN name VARCHAR(255);

