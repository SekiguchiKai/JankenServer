/*このファイルの機能は以下

①1=グー、2=チョキ、3=パーとして打ち手に数字を割り当てる

②乱数を発生させて、機械の打ち手を決める

③クライアントの打ち手を取ってくる

④クライアントの打ち手と機械の打ち手の勝敗判断

⑤Contlloerに送信して、HTMLを変えてもらう
*/

package main

import(
	"math/rand"
	"time"

)

type Janken struct{

	//機械の打ち手
	serverUchite int

	//勝敗の結果を代入するもの
	result string
}


//擬似乱数を発生させて、機械側の打ち手を決定するメソッド
func (j *Janken) decisionUchite() {
	//乱数を発生させる
	rand.Seed(time.Now().UnixNano())//毎回違う乱数を発生させるために、現在の時間をseedに与える

	//jに0~2の擬似乱数を代入
	j.serverUchite = rand.Intn(2)
}

//機械とクライアントの勝敗を決定するメソッド
//クライアントの打ち手を引数に勝敗をジャッジ
//cu = clientUchite
func (j *Janken) JudgeGame (cu *int) string{
	//じゃんけんアルゴリズム
	//機械の打ち手をjuに代入
	jsu := j.serverUchite
	//クライアントの打ち手を代入
	jcu := cu

	switch jcu, jsu{
	//クライアントが勝ちの場合
	case (jcu == 0 && jsu == 2) || (jcu == 1 && jsu == 0) || (jcu == 2 && jsu == 1):
		j.result = "あなたの勝ち"
	//サーバが勝ちの場合
	case (jcu == 2 && jsu == 0) || (jcu == 0 && jsu == 1) || (jcu == 1 && jsu == 2):
		j.result = "あなたの負け"
	//引き分けの場合
	case jcu == jsu:
		j.result = "引き分けね"
	}

	return j.result

}



