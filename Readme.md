drop database socialmedia;
create database socialmedia;
GRANT ALL ON DATABASE socialmedia TO abdillah;
ALTER DATABASE socialmedia OWNER TO abdillah;

migrate -path db/migrations -database "postgresql://abdillah:pass@localhost:5432/socialmedia?sslmode=disable" up
