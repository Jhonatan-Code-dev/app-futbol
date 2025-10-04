package schemas

type TipoPago struct {
	IdTipoPago uint   `gorm:"primaryKey;autoIncrement"`
	TipoPago   string `gorm:"size:30;not null"`
}
