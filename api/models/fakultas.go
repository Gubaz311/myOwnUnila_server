package models

type Fakultas struct {
	ID           int    `gorm:"primaryKey;column:id" json:"id"`
	NamaFakultas string `gorm:"column:namaFakultas" json:"namaFakultas"`
}

func (Fakultas) TableName() string {
	return "fakultas"
}
