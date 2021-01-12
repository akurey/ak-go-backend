package models

import "time"

//Posts posts table structure
type Posts struct {
	ID        uint      `structs:"id,omitnested" json:"id,omitempty" gorm:"primary_key"`
	Title     string    `structs:"title,omitnested" json:"title,omitempty" gorm:"type:text; not null"`
	Content   string    `structs:"content,omitnested" json:"content,omitempty" gorm:"type:longtext; not null"`
	CreatedAt time.Time `structs:"created_at,omitnested" json:"created_at,omitempty" gorm:"type:datetime; default:current_timestamp; not null"`
	UpdatedAt time.Time `structs:"updated_at,omitnested" json:"updated_at,omitempty" gorm:"type:datetime; default:current_timestamp; not null"`
}
