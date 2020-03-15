package structs

import "time"

type EditPermohonan struct {
	ID        uint       `json:"-" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`

	ID_Permohonan			string
	ID_user_pemohon 		string
	Nip 					string
	Status 					string
	Nama_server 			string
	Detail 					string
	Jenis_server 			string
	Os 						string
	Ram 					string
	Storage					string
	Hostname				string
	Internet				string
	Internet_status			string
	Open_port				string
	Lokasi					string
	Id_kontainment			string
	Rak						string
	Id_petugas_approval		string
}

func (EditPermohonan) EditPermohonan()string{
	return "editpermohonan"
}