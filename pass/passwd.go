package pass

import (
	"container/list"
	"fmt"
	"time"
)

type Strength int

const (
	// Any indicates that any password can be used, such as a,b,3
	Any Strength = iota + 1

	// Easy allow passwords that are less secure, even guessable by trial and error
	// such as password, admin, admin123, lightferret
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

// Passwd contains the strength of the password and the password itself
type Passwd struct {
	// S is the strength of the password
	S Strength

	// PwdString is password
	PwdString string
}

// Metadata is the metadata for token
type Metadata struct {
	// Site indicates which site this token for
	Site string

	// Description is the description
	Description string

	// CreateTime is the creation time
	CreateTime time.Time

	// UpdateTime is the update time
	UpdateTime *time.Time

	// History is the history of the token, which organized as a linked list
	History list.List
}

// Token contains token metadata, login name, and passwd
type Token struct {
	// Metadata description token
	Metadata Metadata

	// LoginName is the user name of one site
	LoginName string

	// Passwd is a passwd
	Passwd Passwd
}

func (t *Token) String() string {
	// site, username, passwd, des
	return fmt.Sprintf("%-20s%-20s%-20s%-20s", t.Metadata.Site, t.LoginName, t.Passwd.PwdString, t.Metadata.Description)
}

// // GetSite return site for the token
// func (t *Token) GetSite() string {
// 	return t.Metadata.Site
// }
