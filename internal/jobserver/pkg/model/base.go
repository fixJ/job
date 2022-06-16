package model

type Base struct {
	ID        uint64 `json:"id,omitempty" gorm:"primary_key,AUTO_INCREMENT;column:id"`
	CreatedAt int64  `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt int64  `json:"updated_at,omitempty" gorm:"column:updated_at"`
}
