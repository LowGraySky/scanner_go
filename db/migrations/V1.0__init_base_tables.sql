create schema if not exists dca;

create table if not exists dca.jupiter_token ( address varchar(50) unique, symbol varchar(15));

create table if not exists dca.token_info (symbol varchar(15) unique, is_exists_on_mexc bool null , is_exists_on_bitget bool, is_exists_on_gate bool);
