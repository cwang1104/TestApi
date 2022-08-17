package db

import (
	pbLogin "TestApi/proto/login"
	"fmt"
	"time"
)

type AccountRegisterReq struct {
	Phone int
}

type ResultCount struct {
	Count int
}

func FindUserByPhone(pb *pbLogin.AccountRegisterReq) (*User, error) {
	var user User
	err := DB.Table("user").Where("phone = ?", pb.Phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func IsExistUserByPhone(phone string) bool {
	var result ResultCount
	sql := fmt.Sprintf("select count(*) as count from user where phone = '%s'", phone)
	DB.Raw(sql).Scan(&result)
	if result.Count == 0 {
		return false
	}
	return true
}

func UserRegister(pb *pbLogin.AccountRegisterReq) error {
	addUser := User{
		Phone:      pb.Phone,
		NickName:   "",
		HeadURL:    "",
		Gender:     0,
		Position:   "",
		Birthday:   "",
		Email:      "",
		Introduce:  "",
		Love:       0,
		CreateTime: time.Now(),
	}

	err := DB.Table("user").Create(&addUser).Error
	if err != nil {
		fmt.Println("userRegister failed err ", err)
		return err
	}

	return nil
}
