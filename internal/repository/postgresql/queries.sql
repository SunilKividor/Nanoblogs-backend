CREATE TABLE Blogs(id UUID DEFAULT uuid_generate_v4(),userid UUID,title varchar(255) NOT NULL,content varchar(255) NOT NULL,category varchar(255) NOT NULL,created_at timestamp NOT NULL,updated_at timestamp NOT NULL, PRIMARY KEY(id),CONSTRAINT fk_userid FOREIGN KEY (userid) REFERENCES Users(id));

CREATE TABLE Users (id UUID DEFAULT uuid_generate_v4() ,name VARCHAR(255) NOT NULL,username VARCHAR(255) NOT NULL,password VARCHAR(255) NOT NULL,refresh_token VARCHAR(255), PRIMARY KEY(id));

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE Categories (id UUID DEFAULT uuid_generate_v4(), name VARCHAR(255), frequency INT, PRIMARY KEY (id));

CREATE TABLE user_category (id UUID DEFAULT uuid_generate_v4(), user_id UUID, category_id UUID, frequency INT, PRIMARY KEY(id), CONSTRAINT fk_userid FOREIGN KEY(user_id) REFERENCES Users(id) ON DELETE CASCADE, CONSTRAINT fk_categoryid FOREIGN KEY(category_id) REFERENCES Categories(id) ON DELETE CASCADE);

CREATE TABLE blog_category (id UUID DEFAULT uuid_generate_v4(), blog_id UUID, category_id UUID, PRIMARY KEY(id), CONSTRAINT fk_blogid FOREIGN KEY(blog_id) REFERENCES blogs(id) ON  DELETE CASCADE, CONSTRAINT fk_categoryid FOREIGN KEY(category_id) REFERENCES Categories(id) ON DELETE CASCADE);

ALTER TABLE users ADD COLUMN name VARCHAR(255);

