CREATE TABLE com_user (
	user_id serial4 NOT NULL,
	role_id int4 NOT NULL,
	name varchar(50) NULL,
	email varchar(50) NULL,
	username varchar(100) NULL,
	"password" varchar(100) NULL,
	"file" text NULL,
	mdb timestamp NULL,
	"token" varchar(255) NULL
);