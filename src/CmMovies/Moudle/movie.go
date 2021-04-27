package Moudle

import "time"

type Movie struct {
	Id              string
	TableName       string
	Name            string
	Status          string
	ProtagonistList []Info
	Type            string
	Update          string
	DirectorList    []Info
	Year            time.Time
	Area            string
	Language        string
	Description     string
	Resource        []Info
}
type Info struct {
	Name string
	Url  string
}
