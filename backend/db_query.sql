create role anangs
alter role anangs login encrypted password 'faridanangspolio1123'
select current_user;
set role postgres
-- DROP ROLE IF EXISTS anangs;
-- REVOKE ALL PRIVILEGES ON DATABASE portfolio FROM anangs;

create schema portfolio;
show search_path;
set search_path to portfolio;
-- DROP SCHEMA IF EXISTS portfolio;
RESET search_path;

SELECT table_name 
FROM information_schema.tables 
WHERE table_schema = 'public' AND table_type = 'BASE TABLE';

grant all on schema portfolio to anangs; 

grant select, insert, delete
	on users to anangs;
	
grant select, insert, delete, update
	on skills,projects to anangs;

grant all privileges on database portfolio to anangs;
grant all privileges on portfolio.skills_id_seq to anangs


create table users
(
	id text not null,
	username varchar(50) not null,
	email varchar(150) not null,
	password text not null,
	created_at bigint not null,
	primary key(id),
	constraint users_email_unique unique(email)
);

create table skills
(
	id serial not null,
	image varchar(250) not null,
	name varchar(20) not null,
	created_at bigint not null,
	primary key(id)
);
alter table portfolio.skills
	add column updated_at bigint not null
create table projects
(
	id text not null,
	image varchar(250) not null,
	title varchar(100) not null,
	description text not null,
	tech varchar(300) not null,
	created_at bigint not null,
	primary key(id)
);

alter table portfolio.projects
	add column public_id_image varchar(150)

drop table portfolio.users
drop table portfolio.projects
drop table portfolio.skills


delete from portfolio.projects where title in('welcome to my blog', 'ini adalah hello worrldfdfdfefe', 'ini adalah hello worrld')

select * from portfolio.users;
select * from portfolio.projects;
select * from portfolio.skills;


insert into users(id, username, email, password, created_at)
	values('1', 'anangs', 'anangs@gmail.com', 'anangs', 100)