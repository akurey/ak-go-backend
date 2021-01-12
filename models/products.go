package models

import "time"

// Products products table structure
type Products struct {
	ID        uint      `structs:"id,omitnested" json:"id,omitempty" gorm:"primary_key"`
	Price    float32    `structs:"price,omitnested" json:"price,omitempty" gorm:"not null"`
	CreatedAt time.Time `structs:"created_at,omitnested" json:"created_at,omitempty" gorm:"type:datetime; default:current_timestamp; not null"`
	UpdatedAt time.Time `structs:"updated_at,omitnested" json:"updated_at,omitempty" gorm:"type:datetime; default:current_timestamp; not null"`
}
