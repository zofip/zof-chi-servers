package models

type DomainServers struct {
	ServersChanged   string	`json:"servers_changed"`
	SslGrade         string `json:"ssl_grade"`
	PreviousSslGrade string	`json:"previous_ssl_grade"`
	Logo             string	`json:"logo"`
	Title            string `json:"title"`
	IsDown           bool	`json:"is_down"`
	Servers          []Server `json:"servers"`
}
