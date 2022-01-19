package scripts

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Location struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	State    string `json:"state"`
	Postcode int    `json:"postcode"`
}

type User struct {
	Email     string   `json:"email"`
	Gender    string   `json:"gender"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Location  Location `json:"location"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Picture   string   `json:"picture"`
}

// %v    只输出所有的值
// %+v 先输出字段名字，再输出该字段的值
// %#v 先输出结构体名字值，再输出结构体（字段名字+字段的值）
func (u *User) String() string {
	return fmt.Sprintf("%#v", *u)
}

//func NewUser(firstName string, lastName string) *User {
//	return &User{FirstName: firstName, LastName: lastName}
//}

func testPtr() {
	user1 := User{}
	Info(fmt.Sprintf("%v\t%T", user1, user1))
	user1.FirstName = "Jim"
	Info(fmt.Sprintf("%v\t%T", user1, user1))

	user2 := &User{}
	Info(fmt.Sprintf("%v\t%T", user2, user2))
	user2.FirstName = "Jim"
	Info(fmt.Sprintf("%v\t%T", user2, user2))
	Info(fmt.Sprintf("%v\t%T", *user2, user2))
	Info(fmt.Sprintf("%v\t%T", &user2, user2))
}

func changeUser1(u User) User {
	u.LastName = "LL"
	Info(fmt.Sprintf("inner-user: %v\t%T", u, u))
	return u
}

func changeUser2(u *User) User {
	u.LastName = "LL"
	Info(fmt.Sprintf("inner-user: %v\t%T", u, u))
	return *u
}

func testPtr1() {
	user := User{FirstName: "LiLei"}
	Info(fmt.Sprintf("root-user: %v\t%T", user, user))
	user1 := changeUser1(user)
	Info(fmt.Sprintf("root-user1: %v\t%T", user1, user1))
}

func testPtr2() {
	user := User{FirstName: "LiLei"}
	Info(fmt.Sprintf("root-user: %v\t%T", user, user))
	user1 := changeUser2(&user)
	Info(fmt.Sprintf("root-user1: %v\t%T", user1, user1))
}

func testPtr3() {
	user := &User{FirstName: "LiLei"}
	Info(fmt.Sprintf("root-user: %v\t%T", user, user))
	user1 := changeUser2(user)
	Info(fmt.Sprintf("root-user1: %v\t%T", user1, user1))
	Info(fmt.Sprintf("root-user-last: %v\t%T", user, user))
}

func ToJson(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(data)
}

func FromJson(str string, v interface{}) error{
	return json.Unmarshal([]byte(str), v)
}

func testJson() {
	user := &User{FirstName: "Lei", LastName: "Li", Location: Location{City: "Beijing", Postcode: 100010}}
	jsonStr := ToJson(user)
	Info(jsonStr)

	user2 := &User{}
	err := FromJson(jsonStr, user2)
	if err != nil {
		Error(err)
	} else {
		Info(user2)
	}

	Info(reflect.DeepEqual(user2, user))
}


func JsonMain() {
	testJson()
}