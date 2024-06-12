package main

var Run = make(map[string]bool)

type Config struct {
	MaxInt int      `json:"maxint"`
	Module []string `json:"module"`
	Token  string   `json:"token"`
}
