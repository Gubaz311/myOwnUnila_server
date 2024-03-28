package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Mahasiswa struct {
	NPM          uint32     `gorm:"primaryKey" json:"npm"`
	Nama         string     `gorm:"column:nama" json:"nama"`
	AsalSekolah  string     `gorm:"column:asalSekolah" json:"asalSekolah"`
	NoHape       uint64     `gorm:"column:noHape" json:"noHape"`
	TglMasuk     time.Time  `gorm:"column:tglMasuk" json:"tglMasuk"`
	TglKeluar    *time.Time `gorm:"column:tglKeluar" json:"tglKeluar,omitempty"`
	GpaSemester1 float64    `gorm:"column:gpaSemester1" json:"gpaSemester1"`
	GpaSemester2 float64    `gorm:"column:gpaSemester2" json:"gpaSemester2"`
	GpaSemester3 float64    `gorm:"column:gpaSemester3" json:"gpaSemester3"`
	GpaSemester4 float64    `gorm:"column:gpaSemester4" json:"gpaSemester4"`
}

func (m *Mahasiswa) GetAllData(db *gorm.DB, tahunAwal, tahunAkhir int) (*[]Mahasiswa, error) {
	data := []Mahasiswa{}
	err := db.Where("YEAR(tglMasuk) BETWEEN ? AND ?", tahunAwal, tahunAkhir).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
