package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

// AvatarインスタンスがアバターのURLを返すことができない婆に発生するエラー
var ErrNoAvatrURL = errors.New("Chat: アバターのURLを返すことができない")

// ユーザのプロフィール画像を表す型
type Avatar interface {
	// 指定されたクライアントのアバターURLを返す
	// 問題が発生した場合にはエラーを返す。特にURLを取得できない場合には、
	// ErrNoAvatarURLを返す。
	GetAvatarURL(ChatUser) (string, error)
}

type TryAvatars []Avatar

func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url, nil
		}
	}
	return "", ErrNoAvatrURL
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

func (_ AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if url != "" {
		return url, nil
	}
	return "", ErrNoAvatrURL
}

type GravatarAvatar struct{}

var UseGravatar GravatarAvatar

func (_ GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	return "//www.gravatar.com/avatar/" + u.UniqueID(), nil
}

type FileSystemAvatar struct{}

var UseFileSystemAvatar FileSystemAvatar

func (_ FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
	if files, err := ioutil.ReadDir("avatars"); err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if match, _ := filepath.Match(u.UniqueID()+"*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}
	/*
	matches, err := filepath.Glob(filepath.Join("avatars", u.UniqueID()+"*"))
	if err != nil || len(matches) {
		return "", ErrNoAvatrURL
	}
	return "/" + matches[0], nil
	*/
	return "", ErrNoAvatrURL
}
