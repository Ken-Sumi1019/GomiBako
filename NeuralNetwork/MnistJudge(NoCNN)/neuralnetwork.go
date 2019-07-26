package main

import (
	"os"
	"math"
	"fmt"
	"strconv"
	"./matrix"
	"./getmnist"
	"./dataclean"
)

var w1_1,w1_0,w2_1,w2_0,w3_1,w3_0,layer_z1,layer_a1,layer_z2,layer_a2,layer_z3,layer_a3,dw1,db1,dw2,db2,dw3,db3 [][]float64
var costList []float64
//重みの初期化
func initValue(){
	//データの次元数
	dim := 28*28
	//中間層1のノードの数
	midNode := 64
	//中間層2のノード数
	midNode2:= 64
	//出力層のノード数
	finNode := 10
	//入力層から中間層1への重み
	w1_1 = matrix.MakeWight(dim,midNode)
	//入力層から中間層1へのかけ合わせないやつ
	w1_0 = matrix.MakeWight(1,midNode)
	//中間層1から中間層2への重み
	w2_1 = matrix.MakeWight(midNode,midNode2)
	//中間層1から中間層2へのかけ合わせないやつ
	w2_0 = matrix.MakeWight(1,midNode2)
	//中間層から出力層への重み
	w3_1 = matrix.MakeWight(midNode2,finNode)
	//中間層から出力層へのかけ合わせないやつ
	w3_0 = matrix.MakeWight(1,finNode)
}
//順伝搬
func forward(data [][]float64) [][]float64{
	layer_z1 = matrix.AddVector(matrix.Multi(data,w1_1),w1_0)
	layer_a1 = matrix.Sigmoid(layer_z1)
	layer_z2 = matrix.AddVector(matrix.Multi(layer_a1,w2_1), w2_0)
	layer_a2 = matrix.Sigmoid(layer_z2)
	layer_z3 = matrix.AddVector(matrix.Multi(layer_a2,w3_1), w3_0)
	layer_a3 = matrix.Softmax(layer_z3)
	return layer_a3
}

//コスト
func cost(x [][]float64,y [][]float64) float64{
	return matrix.Square(x,y)
}

//逆伝搬
func back(x [][]float64,y [][]float64){
	//中間層2-出力層の重みでの微分を求める
	output_delta := matrix.Add(layer_a3,matrix.ConstMult(-1,y))
	dw3 = matrix.Multi(matrix.Trans(layer_a2),output_delta)
	db3 = matrix.VecMul(output_delta)

	//中間層1-中間層2の重みで微分を求める
	//fmt.Println(len(output_delta),len(output_delta[0]),len(w2_1),len(w2_1[0]))
	mid2_delta := matrix.AdaMul(matrix.Multi(output_delta,matrix.Trans(w3_1)),matrix.DiffSigmoid(layer_a2))
	dw2 = matrix.Multi(matrix.Trans(layer_a1),mid2_delta)
	db2 = matrix.VecMul(mid2_delta)

	//入力層-中間層の重みでの微分を求める
	mid_delta := matrix.AdaMul(matrix.Multi(mid2_delta,matrix.Trans(w2_1)),matrix.DiffSigmoid(layer_a1))
	dw1 = matrix.Multi(matrix.Trans(x),mid_delta)
	db1 = matrix.VecMul(mid_delta)
}

//重みの更新
func update(alpha float64){
	w1_1 = matrix.Add(w1_1,matrix.ConstMult(-alpha,dw1))
	w1_0 = matrix.Add(w1_0,matrix.ConstMult(-alpha,db1))
	w2_1 = matrix.Add(w2_1,matrix.ConstMult(-alpha,dw2))
	w2_0 = matrix.Add(w2_0,matrix.ConstMult(-alpha,db2))
	w3_1 = matrix.Add(w3_1,matrix.ConstMult(-alpha,dw3))
	w3_0 = matrix.Add(w3_0,matrix.ConstMult(-alpha,db3))
}

func train(x [][]float64,y [][]float64,alpha float64,epoc int){
	//initValue()
	for i:=0;i<epoc;i++{
		forward(x)
		back(x,y)
		update(alpha)
		if epoc % 10 == 0{
			costList = append(costList,cost(forward(x),y))
		}
	}
}

func accuracy(data [][]float64,label []int) float64{
	acc := 0
	for i:=0;i<len(data);i++{
		max := 0.0
		num := 0
		for j:=0;j<10;j++{
			if data[i][j] > max{
				max = data[i][j]
				num = j
			}
		}
		if label[i] == num{
			acc ++
		}
	}
	return float64(acc)/float64(len(label))
}

func main(){
	initValue()
	data := dataclean.ReduceDimention(getmnist.GetTrainImg())
	label := dataclean.Onehotencoding(getmnist.GetTrainLabel())
	epoc := 8
	//確率的勾配法
	for i:=0;i<epoc;i++{
		t,l := dataclean.SplitData(100,data,label)
		count := 0
		costsum := 0.0
		for j:=0;j<len(t);j++{
			x := t[j]
			y := l[j]
			train(x,y,0.1,1)
			count ++
			a := cost(layer_a3,y)
			costsum += a
			if a == math.NaN(){
				i = epoc
				break
			}
			if count % 10 == 0{
				costList = append(costList,costsum)
				fmt.Println(costList[len(costList)-1])
				costsum = 0.0
				count = 0
			}
		}
	}


	output := forward(dataclean.ReduceDimention(getmnist.GetTestImg()))
	la := getmnist.GetTestLabel()
	fmt.Println(accuracy(output,la))

	file, _ := os.Create(`C:\\hoge\\grade.txt`)
	defer file.Close()

	for i:=0;i<len(costList);i++{
		output := []byte(strconv.FormatFloat(costList[i], 'f', 15, 64) + " ")
		file.Write(([]byte)(output))
	}
}
