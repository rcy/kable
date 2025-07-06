create role appuser with password 'appuser' login;
grant appuser to postgres;
create database kable_development with owner appuser;
