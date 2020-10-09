package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Printf(Config())
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/user", getUsers).Methods(http.MethodGet)
	api.HandleFunc("/user", createUser).Methods(http.MethodPost)
	api.HandleFunc("/user/{id}", getUser).Methods(http.MethodGet)
	api.HandleFunc("/team", getTeams).Methods(http.MethodGet)
	api.HandleFunc("/team/{id}", getTeam).Methods(http.MethodGet)
	api.HandleFunc("/bet", getBets).Methods(http.MethodGet)
	api.HandleFunc("/bet/{id}", getBet).Methods(http.MethodGet)
	api.HandleFunc("/match", getMatches).Methods(http.MethodGet)
	api.HandleFunc("/match/{id}", getMatch).Methods(http.MethodGet)

	api.HandleFunc("", notImplemented)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	password, err := HashPassword(r.FormValue("password"))
	if err != nil {
		log.Fatal(err)
	}
	err = register(
		r.FormValue("username"),
		password,
		r.FormValue("email"),
	)

	if err != nil {
		fmt.Printf("Error registering user")
	}
}

func getTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]

	connection := Config()

	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT id, name FROM teams WHERE id = ?`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var teams []Team
	for rows.Next() {
		var e Team
		err = rows.Scan(&e.ID, &e.Name)
		if err != nil {
			log.Fatal(err)
		}
		teams = append(teams, e)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	teamJSON, err := json.Marshal(teams[0])

	if err != nil {
		log.Fatal(err)
	}

	w.Write(teamJSON)
}

func getTeams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	connection := Config()

	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT id, name FROM teams`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var teams []Team
	for rows.Next() {
		var e Team
		err = rows.Scan(&e.ID, &e.Name)
		if err != nil {
			log.Fatal(err)
		}
		teams = append(teams, e)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	teamJSON, err := json.Marshal(teams)

	if err != nil {
		log.Fatal(err)
	}

	w.Write(teamJSON)
}
func getBet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]

	team := Team{
		id,
		"RB Leipzig",
	}

	teamJSON, err := json.Marshal(team)

	if err != nil {
		log.Fatal(err)
	}

	w.Write(teamJSON)
}

func getBets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	teamA := Team{
		"1",
		"RB Leipzig",
	}
	teamB := Team{
		"2",
		"FC Bayern München",
	}

	teams := []Team{
		teamA,
		teamB,
	}

	teamJSON, err := json.Marshal(teams)

	if err != nil {
		log.Fatal(err)
	}

	w.Write(teamJSON)
}
func getMatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]

	team := Team{
		id,
		"RB Leipzig",
	}

	teamJSON, err := json.Marshal(team)

	if err != nil {
		log.Fatal(err)
	}

	w.Write(teamJSON)
}

func getMatches(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	teamA := Team{
		"1",
		"RB Leipzig",
	}
	teamB := Team{
		"2",
		"FC Bayern München",
	}

	teams := []Team{
		teamA,
		teamB,
	}

	teamJSON, err := json.Marshal(teams)

	if err != nil {
		log.Fatal(err)
	}

	w.Write(teamJSON)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]

	connection := Config()

	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT id, username, email, score, t.id, t.name FROM users JOIN teams t ON user.favorite_team = teams.id WHERE user.id = ?`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var e User
		err = rows.Scan(&e.ID, &e.Username, &e.Email, &e.Score, &e.Favorite.ID, &e.Favorite.Name)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, e)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	usersJSON, err := json.Marshal(users[0])

	if err != nil {
		log.Fatal(err)
	}

	w.Write(usersJSON)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	connection := Config()

	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT id, username, email, score, t.id, t.name FROM users JOIN teams t ON user.favorite_team = teams.id`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var e User
		err = rows.Scan(&e.ID, &e.Username, &e.Email, &e.Score, &e.Favorite.ID, &e.Favorite.Name)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, e)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	usersJSON, err := json.Marshal(users)

	if err != nil {
		log.Fatal(err)
	}

	w.Write(usersJSON)
}
