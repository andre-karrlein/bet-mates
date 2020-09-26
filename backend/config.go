package main

import "os"

// Config return configuration for bet-mates
func Config() string {

	// TODO handling if empty environment variables
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	dbname := os.Getenv("DBNAME")

	return user + ":" + password + "@(" + host + ":" + port + ")/" + dbname + "?parseTime=true&charset=utf8"
}