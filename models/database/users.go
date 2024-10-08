package database

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID 			uuid.UUID 	`gorm:"primaryKey;type:uuid;default:uuid_generate_v4(); not null"`
	Name 		string 		`gorm:"type:varchar(255);not null"`
	Email 		string 		`gorm:"type:varchar(255);not null;unique"`
	Password 	string 		`gorm:"type:varchar(255);not null"`
	Address		string 		`gorm:"type:text;"`
	PhoneNumber string 		`gorm:"type:varchar(255);"`
	IsArtisan 	bool 		`gorm:"type:boolean;default:false"`
	IsActive 	bool 		`gorm:"type:boolean;default:true"`
	CreatedBy 	string 		`gorm:"type:varchar(50);not null; default:'system'"`
	UpdatedBy 	string 		`gorm:"type:varchar(50);not null; default:'system'"`
	CreatedAt   time.Time 	`gorm:"autoCreateTime;not null;default:now()"`
	UpdatedAt   time.Time 	`gorm:"autoUpdateTime;not null;default:now()"`

	// Start of References
	Artisan 	Artisans 	`gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Orders 		[]Orders 	`gorm:"foreignKey:UserID"`
	// Carts 		[]Carts 	`gorm:"foreignKey:UserID"`
	// Reviews 	[]Reviews 	`gorm:"foreignKey:UserID"`
	// End of References
}
