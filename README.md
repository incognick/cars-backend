# cars-backend

- Create a user defined bridge network
`docker network create car-net`

- Start a Postgres db, exposing port 5432
`docker run --net car-net --name cars-db -p 5432:5432 -e POSTGRES_PASSWORD=car1 -d postgres`

- Login 
`psql -U postgres -h 127.0.0.1`

- Setup db and user
```sql
create database cars;
create user carsuser with encrypted password '1234';
grant all privileges on database cars to carsuser;
```

- (optional) Install `go-migrate` to create migrations
`brew install golang-migrate`

- Build cars app
`docker build . -tag cars:latest`

- Run app 
`docker run --name cars --net car-net -e DB_HOST=cars-db -e DB_NAME=cars -e DB_USER=carsuser -e DB_PASS=1234 -p 8080:8080 -d cars`

