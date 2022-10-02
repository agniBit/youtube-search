package utils

import (
	"io/ioutil"
	"strings"

	"gorm.io/gorm"
)

func RunSQLFromFile(gdb *gorm.DB, filePath string) {
	sqlStr, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	sqlCoursesString := string(sqlStr)
	commands := strings.Split(sqlCoursesString, "/*delimiter*/")
	
	for _, command := range commands {
		tx := gdb.Exec(command)
		if tx.Error != nil {
			panic(tx.Error)
		}
	}
}
