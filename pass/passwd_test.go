package pass

import (
	"container/list"
	"fmt"
	"testing"
	"time"
)

func Test_Token(t *testing.T) {
	ll := list.New()
	ll.PushBack("abc")

	token := &Token{
		Metadata: &Metadata{
			Site:        "www.github.com",
			Description: "kong",
			CreateTime:  time.Now(),
		},
		LoginName: "yangchnet",
		Passwd:    ll,
	}

	fmt.Printf("%-20s%-20s%-20s%-20s\n", "site", "username", "password", "description")
	fmt.Printf("%-20s%-20s%-20s%-20s\n", "---", "---", "---", "---")
	fmt.Println(token)
}

func Test_NewToken(t *testing.T) {
	// token := NewToken("password", "username", "site", "des")
	// t.Log(token)

	token := (*Token)(nil)
	token.SetToken("passwd")
}
