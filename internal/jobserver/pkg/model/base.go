package model

type Base struct {
	ID        uint64 `json:"id,omitempty" gorm:"primary_key,AUTO_INCREMENT;column:id"`
	CreatedAt uint64 `json:"created_at,omitempty"`
}
