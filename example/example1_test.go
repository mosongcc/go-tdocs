package example

import (
	"github.com/mosongcc/go-tdocs"
	"log"
	"net/http"
	"testing"
)

type User struct {
	Id   int64  `bson:"id" comment:"编号"`
	Name string `bson:"name" comment:"名称"`
}

func (User) TableName() string {
	return "user"
}

func (User) TableTitle() string {
	return "用户信息表"
}

func (User) TableDescription() string {
	return "描述"
}

func TestExample1(t *testing.T) {

	tdocs.Register(tdocs.Bson, User{})

	http.HandleFunc("/tdocs.html", tdocs.HandleFunc)
	log.Print("启动服务 http://localhost")
	log.Fatal(http.ListenAndServe(":80", nil))

}
