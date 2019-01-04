package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func uploaderHandler(w http.ResponseWriter, req *http.Request) {
	// ファイル取り出し
	file, header, err := req.FormFile("avatarFile")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	defer file.Close()

	// ファイル読みこみ
	data, err := ioutil.ReadAll(file)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	// ファイル保存
	filename := filepath.Join("avatars", req.FormValue("userid")+filepath.Ext(header.Filename))
	err = ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	io.WriteString(w, "成功")
}
