package schemas

import "time"

type Fecha struct {
	IdFecha       uint      `gorm:"primaryKey;autoIncrement"`
	FechaApertura time.Time `gorm:"type:datetime;not null"`
	FechaCierre   time.Time `gorm:"type:datetime;not null"`
}
