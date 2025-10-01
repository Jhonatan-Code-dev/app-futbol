package schemas

type TipoPago struct {
	IdTipoPago uint   `gorm:"primaryKey;autoIncrement" json:"id_tipo_pago"`
	TipoPago   string `gorm:"type:varchar(30);not null" json:"tipo_pago"`
}
