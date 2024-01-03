package main

type Command struct {
	Birthday string
	Start    string
}

type Status struct {
	Pending int
}

type Error struct {
	NO_USER_STATE  string
	INTERNAL_ERROR string
}

var Statuses = Status{
	Pending: 0,
}

var Commands = Command{
	Birthday: "birthday",
	Start:    "start",
}

var Errors = Error{
	NO_USER_STATE:  "NO USER STATE",
	INTERNAL_ERROR: "Internal error. Try next time.",
}
