package conf

import (
	"code/gen/util/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Version = "v1.0.0"

type DatabaseConf struct {
	UrlName      string `json:"urlName"`
	IP           string `json:"ip"`
	DatabaseName string `json:"databaseName"`
	Port         string `json:"port"`
	UserName     string `json:"userName"`
	Password     string `json:"password"`
}

type ProjectConf struct {
	Name       string `json:"name"`       //项目名称
	Abbr       string `json:"abbr"`       // 代码目录
	ModName    string `json:"modName"`    // mod 名称
	RouterName string `json:"routerName"` //路由前缀
	Path string `json:"path"` // 代码存放路径
}

var Database = &DatabaseConf{}
var Project = &ProjectConf{}
var DB *gorm.DB



func ResetData() error {
	Database.UrlName = viper.GetString("database.url")
	Database.IP = viper.GetString("database.ip")
	Database.Port = viper.GetString("database.port")
	Database.DatabaseName = viper.GetString("database.database")
	Database.UserName = viper.GetString("database.username")
	Database.Password = viper.GetString("database.password")

	Project.Name = viper.GetString("project.name")
	Project.Abbr = viper.GetString("project.abbr")
	Project.ModName = viper.GetString("project.mod")
	Project.RouterName = viper.GetString("project.router")
	Project.Path = viper.GetString("project.path")
	logger.Log.WithFields(logrus.Fields{"data": Database, "data2": Project}).Info("读取本地配置")
	return Database.GetDB()
}

func (d *DatabaseConf) Save() {
	logger.Log.WithFields(logrus.Fields{"data": d}).Info("保存数据库配置到文件")

	viper.Set("database.url", d.UrlName)
	viper.Set("database.ip", d.IP)
	viper.Set("database.port", d.Port)
	viper.Set("database.username", d.UserName)
	viper.Set("database.password", d.Password)
	viper.Set("database.database", d.DatabaseName)
	err := viper.WriteConfigAs("resource/conf.yaml")
	if err != nil {
		logger.Log.WithFields(logrus.Fields{"data": d, "err": err}).Error("保存数据库配置到文件失败")
	}
}

func (d *DatabaseConf) GetDB() (err error) {
	mysqlUrl := d.UserName + ":" + d.Password + "@tcp(" + d.IP + ":" + d.Port + ")/" + d.DatabaseName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(mysqlUrl), &gorm.Config{})
	if err == nil {
		logger.Log.WithFields(logrus.Fields{"data": mysqlUrl}).Info("数据库连接地址")
		DB = db
	}
	return
}
