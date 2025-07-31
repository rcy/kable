-- prevent new connections to the database
update pg_database set datallowconn = false where datname = 'kable_development';

-- terminate existing connections to the database
select pg_terminate_backend(pg_stat_activity.pid)
from pg_stat_activity
where pg_stat_activity.datname = 'kable_development'
  and pid <> pg_backend_pid();

-- drop databases and roles
drop database if exists kable_development;
drop role if exists appuser;
