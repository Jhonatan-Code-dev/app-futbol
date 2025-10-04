package schemas

type Pago struct {
	IdPago       uint    `gorm:"primaryKey;autoIncrement"`
	IDAsistencia uint    `gorm:"not null"`
	IDUsuario    uint    `gorm:"not null"`
	IDTipoPago   uint    `gorm:"not null"`
	Pago         float64 `gorm:"not null"`

	// Relaciones
	Asistencia Asistencia `gorm:"foreignKey:IDAsistencia;references:IdAsistencia"`
	Usuario    Usuario    `gorm:"foreignKey:IDUsuario;references:IdUsuario"`
	TipoPago   TipoPago   `gorm:"foreignKey:IDTipoPago;references:IdTipoPago"`
}
