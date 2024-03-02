CREATE TABLE Blogs(id UUID DEFAULT uuid_generate_v4(),userid UUID,title varchar(255) NOT NULL,content varchar(255) NOT NULL,category varchar(255) NOT NULL,created_at timestamp NOT NULL,updated_at timestamp NOT NULL, PRIMARY KEY(id),CONSTRAINT fk_userid FOREIGN KEY (userid) REFERENCES Users(id));

CREATE TABLE Users (id UUID DEFAULT uuid_generate_v4() ,username VARCHAR(255) NOT NULL,password VARCHAR(255) NOT NULL,access_token VARCHAR(255) NOT NULL,refresh_token VARCHAR(255) NOT NULL, PRIMARY KEY(id));

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

