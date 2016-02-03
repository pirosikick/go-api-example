package main

import (
	"errors"
	"path/filepath"
)

var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません。")

type Avatar interface {
	AvatarURL(ChatUser) (string, error)
}

type TryAvatars []Avatar

func (a TryAvatars) AvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.AvatarURL(u); err == nil {
			return url, nil
		}
	}
	return "", ErrNoAvatarURL
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

func (AuthAvatar) AvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if url != "" {
		return url, nil
	}
	return "", ErrNoAvatarURL
}

type GravatarAvatar struct{}

// UseGravatar ... Avatar which use gravatar
var UseGravatar GravatarAvatar

// AvatarURL ... return avatar url of the user
func (GravatarAvatar) AvatarURL(u ChatUser) (string, error) {
	return "//www.gravatar.com/avatar/" + u.UniqueID(), nil
}

type FileSystemAvatar struct{}

var UseFileSystemAvatar FileSystemAvatar

func (FileSystemAvatar) AvatarURL(u ChatUser) (string, error) {
	matches, err := filepath.Glob(filepath.Join("avatars", u.UniqueID()+"*"))
	if err != nil || len(matches) == 0 {
		return "", ErrNoAvatarURL
	}
	return "/" + matches[0], nil
}
