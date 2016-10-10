package main

import (
	"net/http"
	"html/template"
	"fmt"
)

//templateで使用するためのデータ型を作成する
type JankenHTML struct {
	ClientUchite string //クライアントの打ち手
	ServeUchite string //サーバの打ち手
	Result string //結果
}

//HTMLをパースするための関数
func htmlHandler(w http.ResponseWriter, r *http.Request) {

	//index.htmlをパースする。
	templ, err := template.ParseFiles("./template/index.html")

	if err != nil {
		fmt.Println(err)//エラー処理
	}

	/*template.Executeは、エラーがなければ、パースしたtemplateを指定したデータオブジェクトにして、
	w http.ResponseWriterに出力してくれる。エラーが起きれば、戻り値の値が入る
	 */
	err = templ.Execute(w, nil)

	if err != nil {
		fmt.Println(err)
	}

}



//formを受け取り、anken_model.goに処理をさせるメソッド

func (jH *JankenHTML) ProcessJudge(w http.ResponseWriter, r *http.Request) {
	//POSTのときはリクエストのボディをフォームとして解析する
	//r.Form（http.Request.Formに値が格納される）
	//戻り値は、err
	//エラー確認
	err := r.ParseForm();
	if err !=nil {
		fmt.Println(err)
	}
	//clientFormの構造体のclientUchiteに代入
	//型変換
	clientUchite := int(r.Form)

	//クライアントの打ち手を引数に渡して、Janken.judgeGameに勝敗を決めさせる
	serveUchite, result := Janken.JudgeGame(clientUchite)

	//HTMLに表示するためにstructの値代入
	jH.Result = result

	//0=グー、1-=チョキ、2=パー
	//数字=>string
	switch clientUchite {
	case 0:
		jH.ServeUchite = "グー"
	case 1:
		jH.ServeUchite = "チョキ"
	case 2:
		jH.ServeUchite = "パー"

	}

	//0=グー、1-=チョキ、2=パー
	//数字=>string
	switch serveUchite {
	case 0:
		jH.ServeUchite = "グー"
	case 1:
		jH.ServeUchite = "チョキ"
	case 2:
		jH.ServeUchite = "パー"

	}


	//result.htmlへリダイレクト
	http.Redirect(w, r, "./result.html", 301)
}


func main() {
	//"/"に対して、リクエストハンドラであるhtmlHandlerをバンドル
	http.HandleFunc("/", htmlHandler)

	//8080ポートをリッスンするようにする=>クライアントからの8080への接続を待ち受けている
	http.ListenAndServe(":8080", nil)


	//form入力後に表示させる"/Result"に対して、リクエストハンドラであるhtmlHandlerをバンドル
	http.HandleFunc("/Result", JankenHTML.ProcessJudge)



}

