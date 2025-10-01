package schemas

type Asistencia struct {
	IdAsistencia uint `gorm:"primaryKey;autoIncrement" json:"id_asistencia"`
	IDUsuario    uint `gorm:"not null" json:"id_usuario"`
	IDFecha      uint `gorm:"not null" json:"id_fecha"`

	// Relaciones
	Usuario Usuario `gorm:"foreignKey:IDUsuario;references:IdUsuario"`
	Fecha   Fecha   `gorm:"foreignKey:IDFecha;references:IdFecha"`
}
