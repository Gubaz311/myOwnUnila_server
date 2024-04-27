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
	TglMasuk     time.Time  `gorm:"column:tglmasuk" json:"tglMasuk"`
	TglKeluar    *time.Time `gorm:"column:tglkeluar" json:"tglKeluar,omitempty"`
	GpaSemester1 float64    `gorm:"column:gpaSemester1" json:"gpaSemester1"`
	GpaSemester2 float64    `gorm:"column:gpaSemester2" json:"gpaSemester2"`
	GpaSemester3 float64    `gorm:"column:gpaSemester3" json:"gpaSemester3"`
	GpaSemester4 float64    `gorm:"column:gpaSemester4" json:"gpaSemester4"`
	FakultasID   int        `gorm:"column:fakultasid" json:"fakultasID"`
	Fakultas     Fakultas   `gorm:"foreignKey:FakultasID" json:"fakultas"`
}

func (Mahasiswa) TableName() string {
	return "mahasiswa"
}

func (m *Mahasiswa) GetDataByRange(db *gorm.DB, tahunAwal, tahunAkhir, fakultasID uint64) (*[]Mahasiswa, error) {
	// return &mahasiswa, nil
	var mahasiswa []Mahasiswa
	// data := db.Debug().Model(&Mahasiswa{}).Where("EXTRACT(YEAR FROM tglmasuk) BETWEEN ? AND ?", tahunAwal, tahunAkhir)
	data := db.Debug().Preload("Fakultas").Model(&Mahasiswa{}).Where("EXTRACT(YEAR FROM tglmasuk) BETWEEN ? AND ?", tahunAwal, tahunAkhir)

	if fakultasID != 0 {
		data = data.Where("fakultasID = ?", fakultasID)
	}

	if err := data.Find(&mahasiswa).Error; err != nil {
		return nil, err
	}

	return &mahasiswa, nil
}

// Mengambil data yang lulus only
func (m *Mahasiswa) GetDataLulus(db *gorm.DB, tahunAwal, tahunAkhir, fakultasID uint64) (*[]Mahasiswa, error) {
	// return &mahasiswa, nil
	var mahasiswa []Mahasiswa
	// data := db.Debug().Model(&Mahasiswa{}).Where("EXTRACT(YEAR FROM tglmasuk) BETWEEN ? AND ? AND tglkeluar IS NOT NULL", tahunAwal, tahunAkhir)
	data := db.Debug().Preload("Fakultas").Model(&Mahasiswa{}).
		Where("EXTRACT(YEAR FROM tglmasuk) BETWEEN ? AND ? ", tahunAwal, tahunAkhir).
		Where("tglKeluar IS NOT NULL")

	if fakultasID != 0 {
		data = data.Where("fakultasID = ?", fakultasID)
	}

	if err := data.Find(&mahasiswa).Error; err != nil {
		return nil, err
	}

	return &mahasiswa, nil
}
