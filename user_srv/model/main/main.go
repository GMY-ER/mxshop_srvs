package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	_ "github.com/anaskhan96/go-password-encoder"
	"io"
	"strings"
)

func genMd5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}
func main() {
	//dsn := "root:root@tcp(152.136.121.186:3306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
	//	logger.Config{
	//		SlowThreshold:             time.Second, // 慢 SQL 阈值
	//		LogLevel:                  logger.Info, // 日志级别
	//		IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
	//		Colorful:                  true,        // 禁用彩色打印
	//	},
	//)
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	NamingStrategy: schema.NamingStrategy{
	//		SingularTable: true,
	//	},
	//	Logger: newLogger,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//_ = db.AutoMigrate(&model.User{})
	//fmt.Println(genMd5("xxxxx_123456"))
	salt, encodedPwd := password.Encode("generic password", nil)
	fmt.Println(salt, encodedPwd)
	check := password.Verify("generic password", salt, encodedPwd, nil)
	fmt.Println(check) //true
	// Using custom options
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd = password.Encode("generic password", options)
	newpassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	fmt.Println(newpassword)
	passwordInfo := strings.Split(newpassword, "$")
	fmt.Println(passwordInfo)
	check = password.Verify("generic password", passwordInfo[2], passwordInfo[3], options)
	fmt.Println(check) // true
	//要加盐值，随机值，否则可能会被暴力破解
	//e10adc3949ba59abbe56e057f20f883e
}
