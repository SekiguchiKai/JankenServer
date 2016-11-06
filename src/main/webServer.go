package main

import (
	"net/http"
	"html/template"
	"fmt"
)


func HTMLResponser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Server running at http://localhost:8080/")

	// ParseFilesは、引数で渡したファイル名のファイルをパースして、関連付けたstructと読み込んだファイルのtemplateを関連付ける
	tmpl, err := template.ParseFiles("./src/main/index.html") // gbを使っているのでアプリのルートディレクトリからの指定
	if err != nil {
		panic(err)
	}


	// パースしたtemplateを第2引数で指定されたデータに適応して、wに結果を書く
	err = tmpl.Execute(w,nil)
	if err != nil {
		panic(err)
	}


}


func main() {
	// URLとハンドラのバンドル
	http.HandleFunc("/", HTMLResponser)
	// 指定んTCPのネットワークアドレスでlistenしている
	http.ListenAndServe(":8080", nil)
}
