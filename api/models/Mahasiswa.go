package models

import "github.com/jinzhu/gorm"

type Mahasiswa struct {
	NPM         uint32  `gorm:"primary_key" json:"npm"`
	Nama        string  `gorm:"type:varchar(30);not null" json:"nama"`
	Waktu_Studi uint    `gorm:"column:waktu_studi;type:numeric" json:"waktuStudi"`
	IPKS1       float64 `gorm:"type:numeric(3,2)" json:"ipks1"`
	IPKS2       float64 `gorm:"type:numeric(3,2)" json:"ipks2"`
	IPKS3       float64 `gorm:"type:numeric(3,2)" json:"ipks3"`
	IPKS4       float64 `gorm:"type:numeric(3,2)" json:"ipks4"`
	NoPhone     uint64  `gorm:"type:numeric(13)" json:"noPhone"`
}

func FindMahasiswaByNPM(db *gorm.DB, npm uint32) (*Mahasiswa, error) {
	var mahasiswa Mahasiswa
	err := db.Debug().Model(&Mahasiswa{}).Where("npm = ?", npm).Take(&mahasiswa).Error
	if err != nil {
		return nil, err
	}
	return &mahasiswa, nil
}
