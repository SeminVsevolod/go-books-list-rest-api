# go-books-list-rest-api
Simple Rest API project on Golang.

## To create database connection

1.	Register on elephantsql https://api.elephantsql.com/. Then select free plan "Tiny" and create instance with name "books-store".

2.	Open "browser" tab and execute following SQL commands in series:
``` SQL
create table books (id serial, title varchar, author varchar, year varchar);
insert into books (title, author, year) values('The Go Programming Language',	'Alan A. A. Donovan, Brian W. Kernighan', '2015');
insert into books (title, author, year) values('Concurrency in Go: Tools and Techniques for Developers', 'Cox-Buday, Katherine', '2017');
insert into books (title, author, year) values('Go in Action', 'William Kennedy, Brian Ketelsen, Erik St. Martin', '2015');
insert into books (title, author, year) values('An Introduction to Programming in Go', 'Caleb Doxsey', '2012');
insert into books (title, author, year) values('Introducing Go: Build Reliable, Scalable Programs', 'Caleb Doxsey', '2016');
```
Then you can check data if you execute following SQL command:
``` SQL
select * from books
```
And you should see your inserted data.

3.	Open "details" tab and copy URL from field "URL". Create file ".env" in the root directory of this project. Paste copied URL to variable ELEPHANTSQL_URL in ".env" file. For example:
``` env
ELEPHANTSQL_URL: "postgres://UsernameFromDetailsTab:PassFromDetailsTab@ruby.db.elephantsql.com:5432/DatabaseNameFromDetailsTab"
```
## To install dependencies:
``` bash
go get
```

## To run program
``` bash
go run main.go
```
