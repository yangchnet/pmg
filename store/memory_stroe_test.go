package store

import (
	"testing"
)

func Test_memoryStore(t *testing.T) {
	store := NewMemoryStore()
	store.Set("github", "passwdforgithub", "yangchnet", "www.github.com", "description")
	if t1, _ := store.Get("github"); t1.LoginName != "yangchnet" ||
		t1.Metadata.Site != "www.github.com" ||
		t1.Passwd.Front().Value != "passwdforgithub" {
		t.Fail()
	}

	store.Set("github", "passwdforgithub222", "yyyyangchnet", "www.github.com", "description")
	store.Set("gitee", "passwdforgitee", "yangchnet", "www.gitee.com", "description")

	if t2, _ := store.Get("github"); t2.LoginName != "yyyyangchnet" ||
		t2.Metadata.Site != "www.github.com" ||
		t2.Passwd.Front().Value != "passwdforgithub222" {
		t.Fail()
	}
}
