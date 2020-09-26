package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

func getTeams(w http.ResponseWriter, r *http.Request) {
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

	userA := User{
		id,
		"akarrlein",
		"testPASSWORD",
		"email@test.com",
		0.0,
		Team{
			"1",
			"RB Leipzig",
		},
	}

	usersJSON, err := json.Marshal(userA)

	if err != nil {
		log.Fatal(err)
	}

	w.Write(usersJSON)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	userA := User{
		"1",
		"akarrlein",
		"testPASSWORD",
		"email@test.com",
		0.0,
		Team{
			"1",
			"RB Leipzig",
		},
	}
	userB := User{
		"2",
		"pweber",
		"testPASSWORD",
		"p@test2.de",
		0.0,
		Team{
			"2",
			"FC Bayern München",
		},
	}

	users := []User{
		userA,
		userB,
	}
	usersJSON, err := json.Marshal(users)

	if err != nil {
		log.Fatal(err)
	}

	w.Write(usersJSON)
}
