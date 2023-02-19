CREATE TABLE EST_STATUS (
	id SMALLINT PRIMARY KEY NOT NULL,
	label VARCHAR(20) NOT NULL
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE EST_USERS (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  status SMALLINT NOT NULL,
  name  VARCHAR(200) NOT NULL,
  password  VARCHAR(200) NOT NULL,
  email VARCHAR(200) NOT NULL UNIQUE,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP,	
  CONSTRAINT fk_statususers
      FOREIGN KEY(status) 
	  REFERENCES EST_STATUS(id)	
);

insert into EST_STATUS values (1, 'normal');
insert into EST_STATUS values (2, 'admin');
insert into EST_STATUS values (3, 'bloqueado');

CREATE TABLE EST_DEVICES (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_agent  VARCHAR(200) NOT NULL,
  user_id uuid NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  last_used TIMESTAMP,	
  CONSTRAINT fk_devicesusers
      FOREIGN KEY(user_id) 
	  REFERENCES EST_USERS(id)	
);