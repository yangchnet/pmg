package pass

import (
	"container/list"
	"encoding/json"
	"time"
)

type Strength int

const (
	// Any indicates that any password can be used, such as a,b,3
	Any Strength = iota + 1

	// Easy allow passwords that are less secure, even guessable by trial and error
	// such as password, admin
	Easy

	// Normal required password is difficult to guess, but can be cracked by automated attacks
	// such as pqlrtmxr、wefourkings
	Normal

	// Hard require strong passwords that are difficult for even users to remember.
	// such as teKz54^JSfyTxpX6, f!tqx9DvYMqN7oPA
	Hard
)

// default strength is hard
var DefaultStrength = Hard

// // Passwd contains the strength of the password and the password itself
// type Passwd struct {
// 	// S is the strength of the password
// 	S Strength

// 	// PwdString is password
// 	PwdString string
// }

// Metadata is the metadata for token
type Metadata struct {
	// Site indicates which site this token for
	Site string `json:"site"`

	// Description is the description
	Description string `json:"description"`

	// CreateTime is the creation time
	CreateTime time.Time `json:"createTime"`

	// UpdateTime is the update time
	UpdateTime *time.Time `json:"updateTime"`
}

// Token contains token metadata, login name, and passwd
type Token struct {
	// Metadata description token
	Metadata *Metadata `json:"metadata"`

	// LoginName is the user name of one site
	LoginName string `json:"loginname"`

	// Passwd is a passwd list, which first elem is current password
	Passwd *list.List `json:"passwd"`
}

func (t *Token) String() string {
	// site, username, passwd, des
	// return fmt.Sprintf("%-20s%-20s%-20s%-20s", t.Metadata.Site, t.LoginName, t.Passwd.Front().Value, t.Metadata.Description)

	pl := []string{}
	for p := t.Passwd.Front(); p != nil; p = p.Next() {
		pl = append(pl, p.Value.(string))
	}

	token := struct {
		Metadata   *Metadata `json:"metadata"`
		LoginName  string    `json:"loginname"`
		PasswdList []string  `json:"passwdlist"`
	}{
		t.Metadata,
		t.LoginName,
		pl,
	}

	tokenByte, err := json.Marshal(token)
	if err != nil {
		return ""
	}

	return string(tokenByte)
}

func (t *Token) SetToken(passwd string, values ...any) *Token {
	if t == nil {
		metadata := &Metadata{
			CreateTime: time.Now(),
		}
		ll := list.New()
		ll.PushBack(passwd)
		t := &Token{}
		t.Metadata = metadata
		t.Passwd = ll

		for i := 0; i < len(values); i++ {
			WithFns[i](values[i].(string))(t)
		}

		return t
	}

	t.Passwd.PushFront(passwd)
	tm := time.Now()
	t.Metadata.UpdateTime = &tm

	for i := 0; i < len(values); i++ {
		WithFns[i](values[i].(string))(t)
	}

	return t
}

// NewToken creates a new token with password, optionally with username, site, description
func NewToken(passwd string, values ...any) *Token {
	metadata := &Metadata{
		CreateTime: time.Now(),
	}
	ll := list.New()
	ll.PushBack(passwd)
	token := &Token{
		Metadata: metadata,
		Passwd:   ll,
	}

	for i := 0; i < len(values); i++ {
		WithFns[i](values[i].(string))(token)
	}

	return token
}

type Option func(*Token)
type WithFn func(string) Option

var WithFns []WithFn = []WithFn{
	WithUserName,
	WithSite,
	WithDescription,
}

func WithUserName(username string) Option {
	return func(t *Token) {
		t.LoginName = username
	}
}

func WithSite(site string) Option {
	return func(t *Token) {
		t.Metadata.Site = site
	}
}

func WithDescription(description string) Option {
	return func(t *Token) {
		t.Metadata.Description = description
	}
}
