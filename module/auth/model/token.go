package usermodel

import (
	"time"

	"github.com/google/uuid"
)

//token thì không cần audit
// type UserToken struct {
// 	Id int `json:"-" gorm:"column:refresh_token"`
// 	UserId int `json:"-" gorm:"column:user_id"`
// 	RefreshToken  string  `json:"-" gorm:"column:refresh_token"`
// 	ExpiredTime *time.Time `json:"-" gorm:"column:expired_time"`
// 	CreatedAt *time.Time `json:"-" gorm:"column:created_at"`
// 	UpdatedAt *time.Time `json:"-" gorm:"column:updated_at"`
// 	CreatedBy int `json:"-" gorm:"column:created_by"`
// 	UpdatedBy int `json:"-" gorm:"column:updated_by"`
// }
type RefreshToken struct {
    ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
    UserID    uuid.UUID `gorm:"type:uuid"`
    Token     string    `gorm:"type:varchar(255);unique_index"`
    ExpiresAt time.Time
    CreatedAt time.Time
}