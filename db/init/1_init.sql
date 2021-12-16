CREATE USER contentman_admin PASSWORD '9TknsRD3';
-- GRANT ALL PRIVILEGES ON DATABASE projectdb TO projectUser;
CREATE DATABASE contentman_db 
    WITH
    OWNER contentman_admin
    ENCODING = 'UTF8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = 50;