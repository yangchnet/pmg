package pass

import (
	"fmt"
	"testing"
	"time"
)

func Test_Token(t *testing.T) {
	token := &Token{
		Metadata: Metadata{
			Site:        "www.github.com",
			Description: "kong",
			CreateTime:  time.Now(),
		},
		LoginName: "yangchnet",
		Passwd: Passwd{
			S:         DefaultStrength,
			PwdString: "f!tqx9DvYMqN7oPA",
		},
	}

	fmt.Printf("%-20s%-20s%-20s%-20s\n", "site", "username", "password", "description")
	fmt.Printf("%-20s%-20s%-20s%-20s\n", "---", "---", "---", "---")
	fmt.Println(token)
}
