package structs

import (
	"time"
)

// Users users table structure
type Users struct {
	ID        uint      `structs:"id,omitnested" json:"id,omitempty" gorm:"primary_key"`
	Username  string    `structs:"username,omitnested" json:"username,omitempty" gorm:"unique; not null"`
	Email     string    `structs:"email,omitnested" json:"email,omitempty" gorm:"unique; not null"`
	Password  string    `structs:"-" json:"password,omitempty" gorm:"not null"`
	CreatedAt time.Time `structs:"created_at,omitnested" json:"created_at,omitempty" gorm:"type:datetime; default:current_timestamp; not null"`
	UpdatedAt time.Time `structs:"updated_at,omitnested" json:"updated_at,omitempty" gorm:"type:datetime; default:current_timestamp; not null"`
}

