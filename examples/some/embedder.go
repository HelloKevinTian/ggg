package some

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func (u *user) notice() {
	fmt.Printf("name[%s] email[%s]\n", u.name, u.email)
}

type admin struct {
	user
	level string
}

// TestEmbed 测试结构嵌入，admin会继承user的所有属性和方法
func TestEmbed() {
	ad := admin{
		user: user{
			name:  "joe",
			email: "joe@qq.com",
		},
		level: "top",
	}
	ad.user.notice()
	ad.notice()
}
