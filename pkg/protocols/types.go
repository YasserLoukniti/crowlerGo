package protocols

import "time"

type Dir struct {
	Id       int64     `json:"id"`
	Hostname string    `json:"hostname"`
	Domain   string    `json:"domain"`
	Lastseen time.Time `json:"lastseen"`
}

type File struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	Url      string    `json:"url"`
	Page     string    `json:"page"`
	Lastseen time.Time `json:"lastseen"`
}
type Database struct {
	Dirs  []Dir
	Files []File
}
