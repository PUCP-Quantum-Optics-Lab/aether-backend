-- roles
CREATE ROLE aether WITH
    NOSUPERUSER
    NOCREATEDB
    NOLOGIN;
CREATE USER aether_adm WITH
    NOSUPERUSER
    NOCREATEDB
    PASSWORD 'not_the_password'
    IN ROLE aether;
CREATE USER aether_api WITH
    NOSUPERUSER
    NOCREATEDB
    PASSWORD 'not_the_password'
    IN ROLE aether;
CREATE USER aether_read WITH
    NOSUPERUSER
    NOCREATEDB
    PASSWORD 'not_the_password'
    IN ROLE aether;

-- db
CREATE DATABASE aether WITH OWNER aether_adm;
GRANT CONNECT ON DATABASE aether TO aether;
