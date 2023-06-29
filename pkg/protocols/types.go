package protocols

import "time"

type Site struct {
	Id       int64     `json:"id"`
	HostIp   string    `json:"hostip"`
	Domain   string    `json:"domain"`
	Lastseen time.Time `json:"lastseen"`
}

type File struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	Url      string    `json:"url"`
	SiteId   int64     `json:"siteId"`
	Lastseen time.Time `json:"lastseen"`
}
type Database struct {
	Sites []Site
	Files []File
}
