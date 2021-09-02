package main

/*
- phppg admin
- pg admin
*/

/*
- sudo apt-get install postgresql-10
- sudo systemctl start postgresql == start as daemon
- sudo -u postgres psql
or
- sudo docker network create postgres
- docker run --network postgres --name dpostgres -e POSTGRES_PASSWORD=password -d postgres
- docker run --network postgres -it --rm postgres psql -h dpostgres -U postgres
- create user myuser password mypass
- CREATE ROLE
- create database mydb owner myuser
- psql 'host=localhost user=myuser password=mypass dbname=mydb' < 001.sql == DDL (data definition language)
- DML (data manipulation language)
- DQL (data query language)
- DSN (Data Source Name)
- pull of connections == active, idle, life of one connection, etc.
- PreparedStatements
- Transactions
- sqlx lib
*/

func main() {

}
