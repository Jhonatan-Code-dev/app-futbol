package schemas

type Pago struct {
	IdPago       uint    `gorm:"primaryKey;autoIncrement" json:"id_pago"`
	IDAsistencia uint    `gorm:"not null" json:"id_asistencia"`
	IDUsuario    uint    `gorm:"not null" json:"id_usuario"`
	IDTipoPago   uint    `gorm:"not null" json:"id_tipo_pago"`
	Pago         float64 `gorm:"not null" json:"pago"`

	// Relaciones
	Asistencia Asistencia `gorm:"foreignKey:IDAsistencia;references:IdAsistencia"`
	Usuario    Usuario    `gorm:"foreignKey:IDUsuario;references:IdUsuario"`
	TipoPago   TipoPago   `gorm:"foreignKey:IDTipoPago;references:IdTipoPago"`
}
