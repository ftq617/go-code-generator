package service

import (
	"code/gen/auto"
	"code/gen/mapper"
	"code/gen/util/autocode"
	"code/gen/util/common"
	"code/gen/util/conf"
	"code/gen/util/logger"
	"code/gen/util/strcase"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

var deletedFiles = map[string]struct{}{
	"id":{},
	"created_at":{},
	"updated_at":{},
	"deleted_at":{},
	"user_id":{},
	"account_id":{},
	"organization_id":{},
	"project_id":{},
}

func Autocode(tableList []auto.TableInfo) error {

	autoData := make([]auto.AutoCodeStruct, 0, len(tableList))

	for _, v := range tableList {
		v.SupStructName = strcase.UpperSnakeCase(v.TableName)
		v.LowStructName = strcase.SnakeCase(v.TableName)
		data,_ := mapper.GetColumns(conf.Database.DatabaseName, v.TableName)
		// 处理 model的继承
		v.BaseInfo = "base"
		var fields []auto.TableColumnInfo
		for _, j := range data {
			j.FieldName = strcase.UpperSnakeCase(j.ColumnName)
			if j.ColumnName == "id" {
				j.FieldName = "ID"
			}
			if j.ColumnName == "account_id" {
				v.BaseInfo = ""
				v.AccountInfo = "account"
			} else if j.ColumnName == "project_id" {
				v.BaseInfo = ""
				v.AccountInfo = ""
				v.ProjectInfo = "project"
			}
			if _,ok := deletedFiles[j.ColumnName];ok {
				continue
			}
			j.FieldJson = j.ColumnName
			j.FieldType = autocode.GetDbType(j.DataType)
			if j.FieldType == "time.Time"{
				v.HasTime = true
			}
			fields = append(fields,j)
		}
		autoData = append(autoData, auto.AutoCodeStruct{ *conf.Project,v, fields, ""})
	}

	var allTempFile []string
	pathName := "resource/temp"
	files, err := ioutil.ReadDir(pathName) // 找出所有模板文件
	if err != nil {
		logger.Log.WithFields(logrus.Fields{"data": err}).Error("代码生成出错")
		return err
	}
	for _, v := range files {
		if strings.HasSuffix(v.Name(), ".tpl") {
			allTempFile = append(allTempFile, pathName+"/"+v.Name())
		}
	}

	for _, tv := range autoData { //数据列表
		for _, fv := range allTempFile { // 文件列表
			if err := autocodeFile(tv, fv); err != nil {
				return err
			}
		}
	}

	return nil
}

func autocodeFile(tv auto.AutoCodeStruct, fv string) error {
	// 开始生成 代码
	autoPath := tv.Path + "\\" + tv.Abbr + "\\"
	if strings.Index(fv, "user tag") >= 0 {
		// TODO 这里根据自己的需求 自己加逻辑
		autoPath += "router/"
	} else {
		autoPath += fv[14:strings.Index(fv, ".")] + "\\"
	}
	if err := DirExistAndMake(autoPath); err != nil {
		return err
	}

	files, err := template.ParseFiles(fv)

	if err != nil {
		logger.Log.WithFields(logrus.Fields{"data": err}).Error("代码生成出错")
		return err
	}
	structName := tv.TableName
	tv.StructName = structName

	file, _ := os.OpenFile(autoPath+structName+".go", os.O_CREATE|os.O_WRONLY, 0755)
	err = files.Execute(file, tv)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{"data": err}).Error("代码生成失败")
		return err
	}
	logger.Log.WithFields(logrus.Fields{"data": autoPath + structName + ".go"}).Info("代码生成文件")
	return nil
}

func DirExistAndMake(autoPath string) error {
	if !common.Exists(autoPath) { // 检查 文件夹是否存在
		if err := os.MkdirAll(autoPath, os.ModePerm); err != nil {
			logger.Log.WithFields(logrus.Fields{"data": err}).Warn("文件夹不存在，创建文件夹出错")
			return err
		}
	}
	return nil
}
