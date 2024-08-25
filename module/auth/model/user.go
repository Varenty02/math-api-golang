package usermodel

import (
	"time"

	"github.com/google/uuid"
)

//không dùng json tag thì tự động mapping theo default convention
type UserDetail struct {
	Id        uuid.UUID        `json:"id" gorm:"type:uuid;default:uuid_generate_v4();column:id;primary_key"`
	Email     string     `json:"email" gorm:"column:email"`
	LastName  string     `json:"last_name" gorm:"column:last_name"`
	FirstName string     `json:"first_name" gorm:"column:first_name"`
	Phone     string     `json:"phone,omitempty" gorm:"column:phone"`
	Role      string     `json:"role,omitempty" gorm:"column:role"`
	Status    int        `json:"status,omitempty" gorm:"column:status"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
	
}
type UserList struct{
	Id        uuid.UUID        `json:"id" gorm:"type:uuid;default:uuid_generate_v4();column:id;primary_key"`
	Email     string     `json:"email" gorm:"column:email"`
	LastName  string     `json:"last_name" gorm:"column:last_name"`
	FirstName string     `json:"first_name" gorm:"column:first_name"`
	Phone     string     `json:"phone,omitempty" gorm:"column:phone"`
	Role      string     `json:"role,omitempty" gorm:"column:role"`
	Status    int        `json:"status,omitempty" gorm:"column:status"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}
type UserCreate struct{
	Id        uuid.UUID        `json:"id" gorm:"type:uuid;default:uuid_generate_v4();column:id;primary_key"`
	Email     string     `json:"email" gorm:"column:email"`
	Password  string     `json:"password" gorm:"column:password"`
	Salt      string     `json:"-" gorm:"column:salt;"`
	LastName  string     `json:"last_name" gorm:"column:last_name;"`
	FirstName string     `json:"first_name" gorm:"column:first_name;"`
	Phone     string     `json:"phone" gorm:"column:phone;"`
	Role      string     `json:"role,omitempty" gorm:"column:role;"`
	Status    int        `json:"status" gorm:"column:status;default:1;"`
	CreatedAt *time.Time `json:"-" gorm:"column:created_at;"`
	CreatedBy int `json:"-" gorm:"column:created_by;"`
}
type UserUpdate struct {
	Email     *string `json:"email" gorm:"column:email;"`
	Password  *string `json:"password" gorm:"column:password;"`
	Salt      *string `json:"-" gorm:"column:salt;"`
	LastName  *string `json:"last_name" gorm:"column:last_name;"`
	FirstName *string `json:"first_name" gorm:"column:first_name;"`
	Phone     *string `json:"phone" gorm:"column:phone;"`
	Role      *string `json:"role" gorm:"column:role;"`
	Status    *int    `json:"status" gorm:"column:status"`
	UpdatedAt *time.Time `json:"-" gorm:"column:updated_at;"`
	UpdatedBy int `json:"-" gorm:"column:updated_at;"`
	DeletedAt *time.Time `json:"-" gorm:"column:deleted_at;"`
	DeletedBy int `json:"-" gorm:"column:deleted_by;"`
}
type UserLogin struct {
	Email    string `json:"email" gorm:"column:email;"`
	Password string `json:"password" gorm:"column:password;"`
}
func (UserDetail) TableName() string { return "users" }
func (UserList) TableName() string { return "users" }
func (UserUpdate) TableName() string { return "users" }
func (UserCreate) TableName() string { return "users" }
func (UserLogin) TableName() string  { return "users" }
