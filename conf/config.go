package conf

import (
	"fmt"
	"strings"
	"xm-mall/dao"

	logging "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string

	PhotoHost        string
	ProductPhotoPath string
	AvatarPath       string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径：", err)
	}

	LoadServerConfig(file)
	LoadMysqlConfig(file)
	LoadEmailConfig(file)
	LoadPhotoPathConfig(file)

	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		logging.Info(err) //日志内容
		panic(err)
	}

	mysqlReadPath := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	mysqlWritePath := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	dao.Database(mysqlReadPath, mysqlWritePath)
}

func LoadServerConfig(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysqlConfig(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadEmailConfig(file *ini.File) {
	ValidEmail = file.Section("email").Key("ValidEmail").String()
	SmtpHost = file.Section("email").Key("SmtpHost").String()
	SmtpEmail = file.Section("email").Key("SmtpEmail").String()
	SmtpPass = file.Section("email").Key("SmtpPass").String()
}

func LoadPhotoPathConfig(file *ini.File) {
	PhotoHost = file.Section("path").Key("Host").String()
	ProductPhotoPath = file.Section("path").Key("ProductPath").String()
	AvatarPath = file.Section("path").Key("AvatarPath").String()
}
