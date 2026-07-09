package main

import (
	"dependency-injection/repo"
	"dependency-injection/service"
	"fmt"
)

func main() {
	//1. Create Database connection
	db := repo.NewDatabase() // ----> dependency(user service depends on it)

	//2. Create User Service
	//2.1 user service requires database as it's dependency.
	//2.2 user service doesn't implement the database logic and only focuses on business logic.
	//2.3 If database changes(MySQL -> Postgres), user service remains untouched and business logic in not affected.
	//2.4 If user service handles the database connection then each user service invocation would create a new connection pool.
	//2.5 unit tests would be complex as it would require db connection, setup, schema etc.
	userService := service.NewUserService(db)

	//3. use user service
	fmt.Println(userService.DummyFunc())
}
