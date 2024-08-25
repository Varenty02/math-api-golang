package appctx

import (
	"gorm.io/gorm"
)

type appContext struct {
	db          *gorm.DB
	secretS3Key string
}
type AppContext interface {
	GetMainConnection() *gorm.DB
	GetSecretKey() string
}

func NewAppContext(db *gorm.DB, secretS3Key string) *appContext {
	return &appContext{
		db:          db,
		secretS3Key: secretS3Key,
	}
}
func (appCtx *appContext) GetMainConnection() *gorm.DB { return appCtx.db }
func (appCtx *appContext) GetSecretKey() string        { return appCtx.secretS3Key }