package main

import (
	"net/http"
	"html/template"
	"fmt"
)

//templateで使用するためのデータ型を作成する
type JankenHTML struct {
	Uchite string
	Shouhai string
}

//HTMLをパースするための関数
func htmlHandler(w http.ResponseWriter, r *http.Request) {
	jankenHTML := JankenHTML{Uchite:"グー", Shouhai:"勝ち!"}

	//index.htmlをパースする。
	templ, err := template.ParseFiles("../template/index.html")

	if err != nil {
		fmt.Println(err)//エラー処理
	}

	/*template.Executeは、エラーがなければ、パースしたtemplateを指定したデータオブジェクトにして、
	w http.ResponseWriterに出力してくれる。エラーが起きれば、戻り値の値が入る
	 */
	err = templ.Execute(w, jankenHTML)

	if err != nil {
		fmt.Println(err)
	}

}



//formを受け取り、anken_model.goに処理をさせるメソッド
func ProcessJudge(r *http.Request) string{
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
	ClientUchite := int(r.Form)

	//クライアントの打ち手を引数に渡して、Janken.judgeGameに勝敗を決めさせる
	result := Janken.JudgeGame(*ClientUchite)
}




func main() {
	//"/"に対して、リクエストハンドラであるhtmlHandlerをバンドル
	http.HandleFunc("/", htmlHandler)

	//8080ポートをリッスンするようにする=>クライアントからの8080への接続を待ち受けている
	http.ListenAndServe(":8080", nil)


	//form入力後に表示させる"/Result"に対して、リクエストハンドラであるhtmlHandlerをバンドル
	http.HandleFunc("/Result", ProcessJudge)


}

