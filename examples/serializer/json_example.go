package serializer

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

// ref: docs/core/serializer/Go-json-overview.md

type User struct {
	Name      string
	IsAdmin   bool
	Followers uint
	Password  string          `json:"-"`                 // 忽略该字段
	Age       int             `json:",omitempty"`        // 如果为空就忽略这个字段，如 nil false 0 len(v)==0
	Address   string          `json:"address,omitempty"` // 别名
	Auth      json.RawMessage `json:",omitempty"`
}

func TestJsonMarshal(t *testing.T) {
	user := &User{
		Name:      "shniu",
		IsAdmin:   true,
		Followers: 10,
	}

	// Marshal 返回 []byte
	data, _ := json.Marshal(user)
	s := string(data)
	fmt.Println(s)
	assert.Equal(t, []byte(`{"Name":"shniu","IsAdmin":true,"Followers":10}`), data)
}

func TestJsonUnmarshal(t *testing.T) {
	data := []byte(`{"Name":"shniu","IsAdmin":true,"Followers":10}`)
	u := &User{}
	err := json.Unmarshal(data, u)
	assert.Nil(t, err)

	assert.Equal(t, "shniu", u.Name)
	assert.Equal(t, true, u.IsAdmin)
	assert.Equal(t, uint(10), u.Followers)
}

func TestJsonTag(t *testing.T) {
	user := &User{
		Name:      "shniu",
		IsAdmin:   true,
		Followers: 12,
		Address:   "bj",
	}

	data, _ := json.Marshal(user)
	assert.Equal(t, []byte(`{"Name":"shniu","IsAdmin":true,"Followers":12,"address":"bj"}`), data)
}

func TestJsonInterface(t *testing.T) {
	data := []byte(`{"Name":"shniu","IsAdmin":true,"Followers":12,"address":"bj"}`)

	var u interface{}

	err := json.Unmarshal(data, &u) // here is &u, not u
	assert.Nil(t, err)
	fmt.Printf("%v\n", u)

	// 更简便的方式
	var u2 map[string]interface{}
	err = json.Unmarshal(data, &u2)
	assert.Nil(t, err)
	fmt.Printf("%v\n", u2)

	name := u2["Name"].(string)
	fmt.Println(name)

	// 解析过来的数字类型默认是 float64, 需要自己手动转成对应的类型
	followers := int(u2["Followers"].(float64))
	fmt.Println(followers)
}

func TestJsonStream(t *testing.T) {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)

	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}

		for k := range v {
			if k != "Name" {
				delete(v, k)
			}
		}

		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}
