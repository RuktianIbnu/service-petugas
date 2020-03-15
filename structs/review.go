package structs

import "time"

type ReviewPetugas struct {
	ID        uint       `json:"-" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

func (ReviewPetugas) ReviewPetugas()string{
	return "reviewpetugas"
}