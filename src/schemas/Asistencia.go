package schemas

type Asistencia struct {
	IdAsistencia uint `gorm:"primaryKey;autoIncrement"`
	IDUsuario    uint `gorm:"not null"`
	IDFecha      uint `gorm:"not null"`

	// Relaciones
	Usuario Usuario `gorm:"foreignKey:IDUsuario;references:IdUsuario"`
	Fecha   Fecha   `gorm:"foreignKey:IDFecha;references:IdFecha"`
}
