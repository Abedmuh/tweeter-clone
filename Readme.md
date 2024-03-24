drop database socialmedia;
create database socialmedia;
GRANT ALL ON DATABASE socialmedia TO abdillah;
ALTER DATABASE socialmedia OWNER TO abdillah;

mingw32-make migrate_up

testing on k6
$env:BASE_URL = 'http://localhost:8000'
k6 run --vus 1 --iterations 1 script.js

