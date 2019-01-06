package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"path/filepath"
)

// ErrNoAvatarURL はAvatarインタンスがアバターのURLを返すことができない場合に発生するエラーです。
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません。")

// Avatar はユーザーのプロフィール画像を表す型です。
type Avatar interface {
	// GetAvatarURLは指定されたクライアントのアバターのURLを返します。
	// 問題が発生した場合にはエラーを返します。
	// 特に、URLを取得できなかった場合にはErrNoAvatarURLを返します。
	GetAvatarURL(c *client) (string, error)
}

// AuthAvatar はユーザーのプロフィール画像取得用メソッドを持った構造体
type AuthAvatar struct{}

// GetAvatarURL はユーザーのプロフィール画像URLを返します
func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

// UseAuthAvatar はAuthAvatarを外部から使用する時用の変数
var UseAuthAvatar AuthAvatar

type GravatarAvatar struct{}

func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			return "//www.gravatar.com/avatar/" + useridStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

var UseGravatar GravatarAvatar

type FileSystemAvatar struct{}

func (FileSystemAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			avatarsDir := "avatars"
			if flag.Lookup("test.v") != nil {
				avatarsDir = "../avatars"
			}
			if files, err := ioutil.ReadDir(avatarsDir); err == nil {
				for _, file := range files {
					println(file.Name())
					if file.IsDir() {
						continue
					}
					if isMatch, _ := filepath.Match(useridStr+"*", file.Name()); isMatch {
						return "/avatars/" + file.Name(), nil
					}
				}
			} else {
				log.Fatal(err)
			}
		}
	}
	return "", ErrNoAvatarURL
}

var UseFileSystemAvatar FileSystemAvatar
