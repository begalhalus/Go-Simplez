CREATE TABLE com_role (
	role_id serial4 NOT NULL,
	role_nm varchar(50) NULL,
	role_desc varchar(50) NULL,
	mdb timestamp DEFAULT current_timestamp
);