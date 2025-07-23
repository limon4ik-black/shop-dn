CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(50) UNIQUE NOT NULL,
  email VARCHAR(100) UNIQUE NOT NULL,
  password_hash TEXT NOT NULL
);

--
-- docker-compose up -d postgres
-- можно убрать postgres и тогда все сервисы запустятся
-- docker exec -it postgres-clear psql -U psql_limon4ik_user -d db
--
-- \dt (show all tables)
-- docker ps (show all running conteiner)