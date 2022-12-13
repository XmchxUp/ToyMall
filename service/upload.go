package service

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"xm-mall/conf"
)

func UploadAvatarToLocaleStatic(file multipart.File, uId uint, userName string) (filePath string, err error) {
	aId := strconv.Itoa(int(uId))
	basePath := "." + conf.AvatarPath + "user" + aId + "/"
	if !ExistDir(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + userName + ".jpg"
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return "", err

	}
	return "user" + aId + "/" + userName + ".jpg", err
}

func ExistDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		return false
	}
	return s.IsDir()
}

func CreateDir(path string) bool {
	err := os.Mkdir(path, 0755)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
