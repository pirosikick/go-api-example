package main

import (
	"testing"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	client := new(client)
	url, err := authAvatar.AvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("値が存在しない場合、AuthAvatar.AvatarURLは" +
			"ErrorNoAvatarURLを返すべきです")
	}
	testUrl := "http://url-to-avatar/"
	client.userData = map[string]interface{}{"avatar_url": testUrl}
	url, err = authAvatar.AvatarURL(client)
	if err != nil {
		t.Error("値が存在する場合、AuthAvatar.AvatarURLは" +
			"エラーを返すべきではありません")
	} else {
		if url != testUrl {
			t.Error("AuthAvatar.AvatarURLは正しいURLを返すべきです")
		}
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	client := new(client)
	client.userData = map[string]interface{}{
		"userid": "0bc83cb571cd1c50ba6f3e8a78ef1346",
	}
	url, err := gravatarAvatar.AvatarURL(client)
	if err != nil {
		t.Error("GravatarAvatar.AvatarURLはエラーを返すべきではありません.")
	}
	if url != "//www.gravatar.com/avatar/0bc83cb571cd1c50ba6f3e8a78ef1346" {
		t.Errorf("GravatarAvatar.AvatarURLが%sという誤った値を返しました", url)
	}
}
