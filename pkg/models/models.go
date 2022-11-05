package models

import "time"

type ExternalField struct {
	Name  string
	Value string
}

type PasswordItems struct {
	ID             int64
	SiteName       string
	Username       string
	Password       string
	Tags           []string
	ExternalFields []ExternalField
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
