package itc_repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
)

func SetupTestDatabase() (testcontainers.Container, *sqlx.DB) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:13",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "0252",
			"POSTGRES_DB":       "ppo",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}
	dbContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Could not start container: %s", err)
	}

	host, err := dbContainer.Host(ctx)
	if err != nil {
		log.Fatalf("Could not get container host: %s", err)
	}

	port, err := dbContainer.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatalf("Could not get container port: %s", err)
	}

	dsn := fmt.Sprintf("user=postgres password=0252 dbname=ppo sslmode=disable host=%s port=%s", host, port.Port())
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	schema := `
	 CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	 
	 CREATE TABLE IF NOT EXISTS users (
	  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	  name TEXT,
	  surname TEXT,
	  email TEXT UNIQUE,
	  phone_number TEXT,
	  address TEXT,
	  password TEXT
	 );
	
	 CREATE TABLE IF NOT EXISTS workers (
	  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	  name TEXT,
	  surname TEXT,
	  email TEXT UNIQUE,
	  phone_number TEXT,
	  address TEXT,
	  password TEXT,
	  role INT
	 );
	
	 CREATE TABLE IF NOT EXISTS orders (
	  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	  worker_id UUID REFERENCES workers(id) ON DELETE SET NULL DEFAULT NULL,
	  user_id UUID REFERENCES users(id) ON DELETE SET NULL DEFAULT NULL,
	  status INT2 DEFAULT 0,
	  address TEXT,
	  deadline TIMESTAMP,
	  creation_date TIMESTAMP DEFAULT NOW(),
	  rate INT2 DEFAULT 0
	 );
	
	 CREATE TABLE IF NOT EXISTS tasks (
	  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	  name TEXT,
	  price_per_single FLOAT8,
	  category INT2
	 );
	
	 CREATE TABLE IF NOT EXISTS order_contains_tasks (
	  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	  order_id UUID REFERENCES orders(id),
	  task_id UUID REFERENCES tasks(id),
	  quantity INT2 DEFAULT 1
	 );
	                                                 
	 CREATE TABLE IF NOT EXISTS categories (
    	id SERIAL,
   		name VARCHAR
	 );`
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatalf("Could not create schema: %s", err)
	}

	return dbContainer, db
}
