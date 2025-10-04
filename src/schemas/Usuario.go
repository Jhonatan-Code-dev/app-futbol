package schemas

import "time"

type Usuario struct {
	IdUsuario uint   `gorm:"primaryKey;autoIncrement"`
	IDRol     uint   `gorm:"not null"`
	Nombre    string `gorm:"size:100;not null"`
	Apellido  string `gorm:"size:100;not null"`
	Correo    string `gorm:"size:255;not null;uniqueIndex"`
	Pass      string `gorm:"size:255;not null"`
	Estado    bool   `gorm:"not null"`

	FechaAceptacion time.Time `gorm:"type:datetime"`
	FechaSolicitud  time.Time `gorm:"type:datetime;not null"`

	Rol Rol `gorm:"foreignKey:IDRol;references:IdRol"`
}
