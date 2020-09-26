package main

import "fmt"

type bet struct {
	ID string `json:"id"`
	Home int `json:"home"`
	Away int `json:"away"`
	Minute int `json:"minute"`
	Match match `json:"match"`
	User User `json:"user"`
}

type match struct {
	ID string `json:"id"`
	Matchday string `json:"matchday"`
	Home Team `json:"home"`
	Away Team `json:"away"`
}

// Team is the model definition of a soccer team
type Team struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

// User is the model definiton of a bet-mate user
type User struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email string `json:"email"`
	Score float64 `json:"score"`
	Favorite Team `json:"favoriteTeam"`
}

func (u *User) changeEmail(email string) error {
	u.Email = email

	return nil
}

func (u *User) changeFavorite(favorite Team) error {
	u.Favorite = favorite
	
	return nil
}

func (u *User) changeScore(points float64) error {
	u.Score = u.Score + points

	return nil
}

func register(username, password, email string) error {
	user := User{
		ID: "1234",
		Username: username,
		Password: password,
		Email: email,
		Score: 0.0,
		Favorite: Team{"0", ""},
	}
	
	fmt.Printf(user)

	return nil
}