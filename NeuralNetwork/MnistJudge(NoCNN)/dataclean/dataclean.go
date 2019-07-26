package dataclean

import (
	"../matrix"
	"math/rand"
	"fmt"
	"time"
)

func Onehotencoding(data []int) [][]float64{
	ans := matrix.MakeMatrix(len(data),10)
	for i:=0;i<len(data);i++{
		ans[i][data[i]] = 1.0
	}
	return ans
}
//データ数*28*28のスライスをデータ数*784にする
func ReduceDimention(data [][][]int) [][]float64{
	ans := matrix.MakeMatrix(len(data),28*28)
	for i:=0;i<len(data);i++{
		for j:=0;j<28;j++{
			for k:=0;k<28;k++{
				ans[i][j*28+k] = float64(data[i][j][k]) / 255.0
			}
		}
	}
	return ans
}
//0からn-1までのインデックスをランダムに並べる
func Shuffle(n int) []int{
	rand.Seed(time.Now().UnixNano())
	ls := make([]int,n)
	for i:=0;i<n;i++{
		ls[i] = i
	}
	for i := len(ls)-1;i>=0;i--{
		j := rand.Intn(i + 1)
		ls[i], ls[j] = ls[j], ls[i]
	}
	return ls
}
//バッチサイズで分割 nはデータ数を割り切れる数にしてほしい
func SplitData(n int,data [][]float64,label [][]float64) ([][][]float64 ,[][][]float64){
	ans := make([][][]float64,len(data)/n)
	anslabel := make([][][]float64,len(data)/n)
	for i:=0;i<len(ans);i++{
		ans[i] = make([][]float64,n)
		anslabel[i] = make([][]float64,n)
		for j:=0;j<len(ans[0]);j++{
			ans[i][j] = make([]float64,len(data[0]))
			anslabel[i][j] = make([]float64,len(label[0]))
		}
	}
	indexes := Shuffle(len(data))
	count := 0
	for i:=0;i<len(ans);i++{
		for j:=0;j<len(ans[0]);j++{
			for k:=0;k<len(data[0]);k++{
				ans[i][j][k] = data[indexes[count]][k]
				//fmt.Print(data[indexes[count]][k])
				if i == 0 && j == 0 {
					//fmt.Print(data[indexes[count]][k], " ")
				}
			}
			for k:=0;k<len(anslabel[0][0]);k++{
				anslabel[i][j][k] = label[indexes[count]][k]
			}
			count ++
		}
	}
	return ans,anslabel
}
