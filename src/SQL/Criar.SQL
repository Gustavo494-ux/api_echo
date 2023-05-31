drop database  Teste_Database;

create database if not exists Teste_Database;
use Teste_Database;

Create table if not exists Tabela_Teste(
	Id int auto_increment primary key,
    Nome varchar(30)
)ENGINE=InnoDB;

CREATE TABLE IF NOT exists Usuarios(
	Id int primary key auto_increment,
    Nome varchar(100), 
    Email varchar(100),
    Senha varchar(100), 
    CriadoEm timestamp default current_timestamp
)ENGINE=InnoDB;