package structs

import "time"

type ReviewPetugas struct {
	ID        uint       `json:"-" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`

	ID_petugas 				string
	ID_struktural_Approval 	string
	Status_review			string

}

func (ReviewPetugas) ReviewPetugas()string{
	return "reviewpetugas"
}

type Ip_Address struct {
	Id_review 	string
	Ip_address 	string
	Ip_public 	string
	Hostname	string
	Netmask		string
	Geteway		string
	Dns			string

	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

func (Ip_Address) Ip_Address()string{
	return "ip_address"
}

type Management struct {
	Id_review 	string
	Ip_address 	string
	Netmask		string
	Geteway		string
	Dns			string
	Username	string
	Password	string

	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

func (Management) Management()string{
	return "management"
}

type Spesifikasi struct {
	Id_review 		string
	Jenis_server	string
	Os				string
	Ram				string
	Storage			string

	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

func (Spesifikasi) Spesifikasi()string{
	return "spesifikasi"
}