CREATE SEQUENCE blog_tag_seq;
CREATE TABLE blog_tag (
  id int check (id > 0) NOT NULL DEFAULT NEXTVAL ('blog_tag_seq'),
  name varchar(100) DEFAULT '' ,
  created_on int check (created_on > 0) DEFAULT '0' ,
  created_by varchar(100) DEFAULT '' ,
  modified_on int check (modified_on > 0) DEFAULT '0' ,
  modified_by varchar(100) DEFAULT '' ,
  state smallint check (state > 0) DEFAULT '1' ,
  PRIMARY KEY (id)
);

CREATE SEQUENCE blog_article_seq;
CREATE TABLE blog_article (
  id int check (id > 0) NOT NULL DEFAULT NEXTVAL ('blog_article_seq'),
  tag_id int check (tag_id > 0) DEFAULT '0' ,
  title varchar(100) DEFAULT '' ,
  description varchar(255) DEFAULT '',
  content text,
  created_on int DEFAULT NULL,
  created_by varchar(100) DEFAULT '' ,
  modified_on int check (modified_on > 0) DEFAULT '0' ,
  modified_by varchar(255) DEFAULT '' ,
  state smallint check (state > 0) DEFAULT '1' ,
  PRIMARY KEY (id)
);

CREATE SEQUENCE blog_auth_seq;
CREATE TABLE blog_auth (
  id int check (id > 0) NOT NULL DEFAULT NEXTVAL ('blog_auth_seq'),
  username varchar(50) DEFAULT '' ,
  password varchar(50) DEFAULT '' ,
  PRIMARY KEY (id)
);

INSERT INTO blog_auth (username, password) VALUES ('test', 'test123456');