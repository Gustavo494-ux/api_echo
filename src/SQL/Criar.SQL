drop database  Teste_Database;
create database if not exists Teste_Database;
use Teste_Database;
Create table if not exists Tabela_Teste(
	Id int auto_increment primary key,
    Nome varchar(30)
)ENGINE=InnoDB;
