drop database if exists voting;
create database voting;
use voting;

create table auth (
	last_logged_in timestamp,
	login_count int default 1,
	email varchar(50),
	passwd varchar(100),
	account_created timestamp default current_timestamp,
	id int unique,
	primary key(email)
);

create table users (
	fname varchar(30),
	lname varchar(30),
	class int,
	id int,	
	primary key(id),
	foreign key(id) references auth(id)
);

create table candidates (
	id int,
	position varchar(30),
	photo blob,
	primary key(id),
	foreign key(id) references users(id)
);

create table vote_count (
	id int,
	number_of_votes int,
	primary key(id),
	foreign key(id) references candidates(id)
);

create table dummy (
  id int,
  description varchar(30),
  PRIMARY KEY(id)
);
