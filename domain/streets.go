package domain

type Streets struct {
	ID        int64  `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"not null;column:name"`
	Kelurahan string `json:"kelurahan" gorm:"not null;column:kelurahan"`
	Kecamatan string `json:"kecamatan" gorm:"not null;column:kecamatan"`
	Kota      string `json:"kota" gorm:"not null;column:kota"`
}
