package models

type Users struct {
	Fname string
	Lname string
	Class int
}

type Auth struct {
	Email    string
	Password string
}

type Candidates struct {
	Position string
	Photo    []byte
}

type VoteCount struct {
	Number_of_votes int
}
