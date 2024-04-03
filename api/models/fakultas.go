package models

type Fakultas struct {
	ID           int    `gorm:"column:id" json:"id"`
	NamaFakultas string `gorm:"column:namaFakultas" json:"namaFakultas"`
}

func (Fakultas) TableName() string {
	return "fakultas"
}
