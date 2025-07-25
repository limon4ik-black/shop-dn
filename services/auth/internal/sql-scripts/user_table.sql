CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(50) UNIQUE NOT NULL,
  password_hash TEXT NOT NULL
);

--
-- docker-compose up -d postgres
-- можно убрать postgres и тогда все сервисы запустятся
-- docker exec -it postgres-clear psql -U psql_limon4ik_user -d db
--
-- \dt (show all tables)
-- docker ps (show all running conteiner)


CREATE OR REPLACE FUNCTION insert_user(p_username VARCHAR, p_password_hash TEXT)
RETURNS INTEGER AS $$
DECLARE
    new_user_id INTEGER;
BEGIN
    INSERT INTO users (username, password_hash)
    VALUES (p_username, p_password_hash)
    RETURNING id INTO new_user_id;

    RETURN new_user_id;
EXCEPTION
    WHEN unique_violation THEN
        RAISE EXCEPTION 'Пользователь с именем "%" уже существует.', p_username;
END;
$$ LANGUAGE plpgsql;

-- есть таблица \dt и функция \df