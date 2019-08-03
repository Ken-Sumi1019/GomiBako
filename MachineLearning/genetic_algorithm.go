package main

import (
	"math/rand"
	"fmt"
)

/*
遺伝的アルゴリズムを実装していみる
ルーレット選択、二点交叉を用いる
エリート保存も行う(未実装)
*/
//遺伝子の長さ
var genom_length = 100
//遺伝子の数
var genom_size = 100
//交叉するペアの数
var cross_pairs = 30
//交叉する確率　小数点第二位まで
var cross_prob = 0.8
//繰り返す世代数
var generation = 500
//突然変異確率 小数点第三位まで指定しておｋ
var mutation_prob = 0.01
//遺伝子配列
var ls = [][]int{}
//遺伝子の配列を作ります（初期値）
func init() {
	ls = make([][]int,genom_size)
	for i:=0;i<genom_size;i++{
		ls[i] = make([]int,genom_length)
		for j:=0;j<genom_length;j++{
			ls[i][j] = rand.Intn(2)
		}
	}
}
//遺伝子の優秀さを評価する
func accuracy(ls []int) int {
	c := 0
	for i:=0;i<len(ls);i++{
		if ls[i] == 1{c ++}
	}
	return c
}

//cross_pars個のペアを生成する関数 インデックスを返す
//ようは重複を許さずに乱数でわちゃわちゃするということ
func makePairs() [][]int {
	check_ls := make([]int,genom_size)
	for i:=0;i<len(check_ls);i++{
		check_ls[i] = i
	}
	//おなじみのフィッシャーなんたらのアルゴリズムでシャッフル
	for i:=len(check_ls)-1;i>0;i--{
		n := rand.Intn(i)
		check_ls[i],check_ls[n] = check_ls[n],check_ls[i]
	}
	//頭から二つずつ入れてく
	ans := make([][]int,cross_pairs)
	for i:=0;i<cross_pairs*2;i++{
		if i%2 == 0{ans[i/2] = make([]int,2)}
		ans[i/2][i%2] = check_ls[i]
	}
	return ans
}
//二点交叉する関数。
//引数に渡した配列をそのまま変える
func crossing() {
	indexes := makePairs()
	for i:=0;i<len(indexes);i++{
		prob := rand.Intn(100)
		if int(cross_prob*100) - 1 < prob{continue}
		a := rand.Intn(genom_length)
		b := rand.Intn(genom_length)
		if a > b{a,b = b,a}
		for j:=a;j<b;j++{
			ls[indexes[i][0]][j],ls[indexes[i][1]][j] = ls[indexes[i][1]][j],ls[indexes[i][0]][j]
		}
	}
}

//突然変異を起こす
func mutation() {
	for i:=0;i<len(ls);i++{
		n := rand.Intn(1000)
		if int(mutation_prob*1000) - 1 < n{continue}
		n = rand.Intn(genom_length)
		if ls[i][n] == 0{
			ls[i][n] = 1
		} else {
			ls[i][n] = 0
		}
	}
}

//スライスを二分探索して指定した数が入りそうなインデックスを返す
//スライス[2,5,7,9,10] n = 6のとき2を返す
func binary_search(list []int,n int) int {
	if list[0] >= n {
		return 0
	}
	front := 0
	back := len(list) - 1
	for ; ;  {
		i := (front + back) / 2
		if list[i] >= n{
			back = i
		} else {
			front = i
		}
		if back - front == 1{
			return back
		}
	}
}

//ルーレット選択を行います。
func selectGenetic() [][]int {
	//次の世代
	next_generation := make([][]int,genom_size)
	//ルーレット選択用の重みを格納したスライス
	select_list := make([]int,genom_size)
	for i:=0;i<genom_size;i++{
		select_list[i] = accuracy(ls[i])
		if i != 0{select_list[i] += select_list[i-1]}
	}
	//エリート選択を実装　一番よい個体を強制的に後世に残す
	maxIndex := 0
	maxAccuracy := 0
	for i:=0;i<len(ls);i++{
		cc := accuracy(ls[i])
		if maxAccuracy < cc{
			maxIndex = i
			maxAccuracy = cc
		}
	}
	next_generation[0] = make([]int,genom_length)
	//怖いからコピー走らせとく
	for j:=0;j<genom_length;j++{
		next_generation[0][j] = ls[maxIndex][j]
	}
	//指定回数分だけランダムに取り出す
	for i:=1;i<genom_size;i++{
		next_generation[i] = make([]int,genom_length)
		n := rand.Intn(select_list[len(select_list)-1]+1)
		index := binary_search(select_list,n)
		//怖いからコピー走らせとく
		for j:=0;j<genom_length;j++{
			next_generation[i][j] = ls[index][j]
		}
	}
	//置き換え
	ls = next_generation
	return next_generation
}

func main(){
	//世代数だけループ
	for i:=0;i<generation;i++{
		selectGenetic()
		crossing()
		mutation()
	}
	//一番良い個体を出力
	maxIndex := 0
	maxAccuracy := 0
	for i:=0;i<len(ls);i++{
		if maxAccuracy < accuracy(ls[i]){
			maxIndex = i
			maxAccuracy = accuracy(ls[i])
		}
	}
	fmt.Println(ls[maxIndex])
	fmt.Println(maxAccuracy)
	fmt.Print("のん！")
}
