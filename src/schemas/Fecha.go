package schemas

import "time"

type Fecha struct {
	IdFecha       uint      `gorm:"primaryKey;autoIncrement" json:"id_fecha"`
	FechaApertura time.Time `gorm:"type:datetime;not null" json:"fecha_apertura"`
	FechaCierre   time.Time `gorm:"type:datetime;not null" json:"fecha_cierre"`
}
