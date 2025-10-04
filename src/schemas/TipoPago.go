package schemas

type TipoPago struct {
	IdTipoPago uint   `gorm:"primaryKey;autoIncrement"`
	TipoPago   string `gorm:"type:varchar(30);not null"`
}
