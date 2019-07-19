package main

import
(
	"math/rand"
	"fmt"
	"reflect"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
	"os"
)

//指数
func pow(n int,m int) int{
	if m == 1{
		return n
	}else if m == 0 {
		return 1
	}else if m % 2 == 0{
		k := pow(n,m/2)
		return  k*k
	}else {
		return n*pow(n,m-1)
	}
}

func hash(str []byte)  [32]byte {
	return [32]byte(sha256.Sum256(str))
}
//[32]byteを[]byteに変換して16進数文字列に変換する
func convhex(b [32]byte) string{
	by := make([]byte,32)
	for i:=0;i<32;i++{
		by[i] = b[i]
	}
	return hex.EncodeToString(by)
}
//256bitのバイト列の比較 a < bのときtrue
func comper(a [32]byte,b []byte) int{
	check := make([]byte,32)
	for i:=0;i<32;i++{
		check[i] = a[i]
	}
	return bytes.Compare(b,check)
}
//2^nの256bitのバイト列を生成
func makebyte(n int) []byte{
	ans := make([]byte,32)
	if n >= 256{
		ans[0] = byte(255)
		return ans
	}
	amari := n % 8
	syo := n / 8
	check := pow(2,amari)
	ans[31-syo] = byte(check)
	return ans
}

//ハッシュ値を
//ブロックチェーンの中身の構造体
type block struct {
	Prevhash [32]byte
	Nonce int
	Difficulty int
	Name string
}

//ブロックを作成
func makeBlock(hash [32]byte,nonce int,difficulty int,name string) block{
	var a block
	a.Prevhash = hash
	a.Nonce = nonce
	a.Difficulty = difficulty
	a.Name = name
	return a
}

//ブロックをjson文字列に変換する
func convjson(b block) string{
	var buf bytes.Buffer
	byt, _ := json.Marshal(b)
	buf.Write(byt)
	return buf.String()
}


//ブロックチェーン
//var chain = []block{makeBlock([32]byte{},0)}
//チェーンのリスト
var list = make([][]block,4)

//初期化
func initList(){
	for i:=0;i<5;i++{
		list[i] = []block{makeBlock([32]byte{},0,256,"")}
	}
}


//チェーンの整合性をチェック
func check(ls []block) bool{
	//前後のブロックとハッシュ値の関係に矛盾がないかどうか
	for i:=0;i<len(ls) - 1;i++{
		//前のブロックのハッシュ値
		before := hash([]byte(convjson(ls[i])))
		//次のブロックに書いてある前のブロックのハッシュ値
		after := ls[i+1].Prevhash
		if ! reflect.DeepEqual(after,before){
			return false
		}
	}
	//ハッシュ値が2^difficultyよりも小さいかどうか
	for i:=0;i<len(ls);i++{
		//制限のバイト配列
		limit := makebyte(ls[i].Difficulty)
		//現在のブロックのハッシュ値
		nowBlock_hash := hash([]byte(convjson(ls[i])))
		if comper(nowBlock_hash,limit) != 1{
			fmt.Println("yei")
			return false
		}
	}
	return true
}

//全てのチェーンを調べて一番長くて整合性のあるチェーンを返す
func consensys() []block{
	//チェーンの長さを入れる。整合性のないチェーンの場合は-1
	chainLength := make([]int,len(list))
	for i:=0;i<len(list);i++{
		if check(list[i]){
			chainLength[i] = len(list[i])
		} else {
			chainLength[i] = -1
		}
	}
	maxLength := 0
	maxIndex := 0
	for i:=0;i<len(chainLength);i++{
		if chainLength[i] > maxLength{
			maxLength = chainLength[i]
			maxIndex = i
		}
	}
	ansChain := make([]block,maxLength)
	for i:=0;i<maxLength;i++{
		ansChain[i] = list[maxIndex][i]
	}
	return ansChain
}

//マイニング diffが小さいほど制約が大きい indexはチェーンのリストの何番目か
func mine(diff int,index int,name string) string{
	lastBlock := list[index][len(list[index]) - 1]
	lastBlock_json := convjson(lastBlock)
	lastBlock_hash := hash([]byte(lastBlock_json))
	limit := makebyte(diff)
	var n block
	for {
		n = makeBlock(lastBlock_hash,rand.Int(),diff,name)
		nowBlock_json := convjson(n)
		nowBlock_hash := hash([]byte(nowBlock_json))
		if comper(nowBlock_hash,limit) == 1{
			break
		}
	}
	list[index] = append(list[index],n)
	list[index] = consensys()
	return convjson(n)
}

//n番目のチェーンの様子を見る
func show(n int) {
	for i:=0;i<len(list[n]);i++{
		fmt.Println(convhex(list[n][i].Prevhash))
	}
}

func main(){
	initList()
	//Aさん
	go func() {
		for {
			mine(235,0,"A")
		}
	}()
	//Bさん
	go func() {
		for {
			mine(235,1,"B")
		}
	}()
	//Cさん
	go func() {
		for {
			mine(235,2,"C")
		}
	}()
	//Dさん
	go func() {
		for {
			mine(235,3,"D")
		}
	}()

	go func() {
		for {
			time.Sleep(5*time.Second)
			ls := consensys()
			for i:=0;i<len(ls);i++{
				fmt.Println(ls[i])
			}
		}
	}()

	time.Sleep(300*time.Second)

	file, err := os.Create(`C:\\pythontest\\block_chain.txt`)
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()

	ls := consensys()
	for i:=0;i<len(ls);i++{
		output := convjson(ls[i]) + "\r\n"
		file.Write(([]byte)(output))
	}
}
