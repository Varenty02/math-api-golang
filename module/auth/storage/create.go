package authstorage

import (
	common "code-math/common"
	usermodel "code-math/module/auth/model"
	"context"
)
func(s *sqlStore) CreateUser(ctx context.Context, data usermodel.UserCreate) (error){
	db:=s.db.Begin()
	if err:=db.Table(data.TableName()).Create(data).Error;err!=nil{
		db.Rollback();
		return common.ErrorDB(err)
	}
	if err:=db.Commit().Error;err!=nil{
		return common.ErrorDB(err)
	}
	return nil
}