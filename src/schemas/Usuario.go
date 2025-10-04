package schemas

import "time"

type Usuario struct {
	IdUsuario uint   `gorm:"primaryKey;autoIncrement" json:"id_usuario"`
	IDRol     uint   `gorm:"not null" json:"id_rol" validate:"required,gte=1"`
	Nombre    string `gorm:"type:varchar(100);not null" json:"nombre" validate:"required,alphaSpace,min=2,max=100"`
	Apellido  string `gorm:"type:varchar(100);not null" json:"apellido" validate:"required,alphaSpace,min=2,max=100"`
	Correo    string `gorm:"type:varchar(255);unique;not null" json:"correo" validate:"required,email,gmail"`
	Pass      string `gorm:"type:varchar(255);not null" json:"pass" validate:"required,min=6"`
	Estado    bool   `gorm:"not null" json:"estado"`

	FechaAceptacion time.Time `gorm:"type:datetime;not null" json:"fecha_aceptacion"`
	FechaSolicitud  time.Time `gorm:"type:datetime;not null" json:"fecha_solicitud"`

	Rol Rol `gorm:"foreignKey:IDRol;references:IdRol"`
}
