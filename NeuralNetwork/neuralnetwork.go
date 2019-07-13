package main

import (
	"os"
	"fmt"
	"math"
	"math/rand"
	"time"
	"strconv"
)

//大きい値を出力
func max(a float64,b float64) float64{
	if a > b{return a} else{return b}
}

//行列を生成
func makeMatrix(i int,j int) [][]float64{
	ans := make([][]float64,i)
	for k:=0;k<i;k++{
		ans[k] = make([]float64,j)
	}
	return ans
}
//ベクトルの最大値を求める
func vecMax(x [][]float64) float64{
	if len(x[0]) != 1 && len(x) != 1{
		println("ふええ～おにいちゃーん、これベクトルじゃないよぉ～～")
		os.Exit(1)
	}
	ans := 0.0
	for i:=0;i<len(x);i++{
		for j:=0;j<len(x[0]);j++{
			ans = max(ans,x[i][j])
		}
	}
	return ans
}
//ベクトルのexpの合計
func vecExp(x [][]float64) float64{
	ans := 0.0
	for i:=0;i<len(x);i++{
		for j:=0;j<len(x[0]);j++{
			ans = math.Exp(x[i][j])
		}
	}
	return ans
}
//行列を合計
func add(a [][]float64,b [][]float64) [][]float64{
	if (len(a) != len(b)) || (len(a[0]) != len(b[0])){
		str := "ふええ～おにいちゃーん、大きさが違う行列は加算できないよぉ～～"
		println(str)
		os.Exit(1)
	}
	ans := makeMatrix(len(a),len(a[0]))
	for i:=0;i<len(a);i++{
		for j:=0;j<len(a[0]);j++{
			ans[i][j] = a[i][j] + b[i][j]
		}
	}
	return ans
}
//アダマール積
func adaMul(a [][]float64,b [][]float64) [][]float64{
	if (len(a) != len(b)) || (len(a[0]) != len(b[0])){
		str := "ふええ～おにいちゃーん、大きさが違う行列は計算できないよぉ～～"
		println(str)
		os.Exit(1)
	}
	ans := makeMatrix(len(a),len(a[0]))
	for i:=0;i<len(a);i++{
		for j:=0;j<len(a[0]);j++{
			ans[i][j] = a[i][j] * b[i][j]
		}
	}
	return ans
}
//行列の積
func multi(a [][]float64,b [][]float64) [][]float64{
	if len(a[0]) != len(b){
		println("ふええ～おにいちゃーん、内積が計算できないよぉ～～")
		os.Exit(1)
	}
	ans := makeMatrix(len(a),len(b[0]))
	for i:=0;i<len(a);i++{
		for j:=0;j<len(b[0]);j++{
			for k:=0;k<len(b);k++{
				ans[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return ans
}
//行列に定数を加算
func constAdd(n float64,mat [][]float64) [][]float64{
	ans := makeMatrix(len(mat),len(mat[0]))
	for i:=0;i<len(mat);i++{
		for j:=0;j<len(mat[0]);j++{
			ans[i][j] = mat[i][j] + n
		}
	}
	return ans
}
//行列を定数倍
func constMult(n float64,matrix [][]float64) [][]float64{
	ans := makeMatrix(len(matrix),len(matrix[0]))
	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix[0]);j++{
			ans[i][j] = n * matrix[i][j]
		}
	}
	return ans
}
//行列の転置
func trans(a [][]float64) [][]float64{
	ans := makeMatrix(len(a[0]),len(a))
	for i:=0;i<len(a[0]);i++{
		for j:=0;j<len(a);j++{
			ans[i][j] = a[j][i]
		}
	}
	return ans
}
//行列にベクトルを足し算
/*
[[1,2,3],					[[1+1,1+2,1+3],
 [4,5,6],  +  [[1,1,4]]  =   [4+1,5+1,6+3],
 [7,8,9]]					 [7+1,8+1,9+3]]
*/
func addVector(mat [][]float64,vec [][]float64) [][]float64{
	if len(vec) != 1 {
		println("ふええ～おにいちゃーん、ベクトルじゃないないよ～～")
		os.Exit(1)
	} else if len(mat[0]) != len(vec[0]){
		println("ふええ～おにいちゃーん、列の数が違うよ～～")
		os.Exit(1)
	}
	ans := makeMatrix(len(mat),len(mat[0]))
	for i:=0;i<len(mat);i++{
		for j:=0;j<len(mat[0]);j++{
			ans[i][j] = mat[i][j] + vec[0][j]
		}
	}
	return ans
}
//要素1、大きさ1行n列のベクトルとの内積つまり、列方向の合計
func vecMul(mat [][]float64) [][]float64{
	ans := makeMatrix(1,len(mat[0]))
	for i:=0;i<len(mat);i++{
		for j:=0;j<len(mat[0]);j++{
			ans[0][j] += mat[i][j]
		}
	}
	return ans
}
//行列に行列を追加（横方向）
/*
func addMat(a [][]float64,b [][]float64) [][]float64{
	if len(a) != len(b){
		println("ふええ～おにいちゃーん、行数が違うからつなげれないよ～～")
		os.Exit(1)
	}
	ans := makeMatrix(len(a),len(a[0]) + len(b[0]))
	for i:=0;i<len(a);i++{
		for j:=0;j<len(a[0]);j++{
			ans[i][j] = a[i][j]
		}
		for j:=0;j<len(b[0]);j++{
			ans[i][len(a[0]) + j] = b[i][j]
		}
	}
	return ans
}
 */
//relu
func relu(a [][]float64) [][]float64{
	ans := makeMatrix(len(a),len(a[0]))
	for i:=0;i<len(a);i++{
		for j:=0;j<len(a[0]);j++{
			ans[i][j] = max(0,a[i][j])
		}
	}
	return ans
}
//reluの微分
func diffRelu(a [][]float64) [][]float64{
	ans := makeMatrix(len(a),len(a[0]))
	for i:=0;i<len(a);i++{
		for j:=0;j<len(a[0]);j++{
			if a[i][j] > 0{
				ans[i][j] = 1
			} else {
				ans[i][j] = 0
			}
		}
	}
	return ans
}
//シグモイド関数
func sigmoid(x [][]float64) [][]float64{
	ans := makeMatrix(len(x),len(x[0]))
	for i:=0;i<len(x);i++{
		for j:=0;j<len(x[0]);j++{
			ans[i][j] = 1 / (1 + math.Exp(-x[i][j]))
		}
	}
	return ans
}
//シグモイド関数の微分(出力された値を引数に入れる)
func diffSigmoid(x [][]float64) [][]float64{
	ans := makeMatrix(len(x),len(x[0]))
	for i:=0;i<len(x);i++{
		for j:=0;j<len(x[0]);j++{
			ans[i][j] = x[i][j] * (1.0 - x[i][j])
		}
	}
	return ans
}
//ソフトマックス関数
func softmax(x [][]float64) [][]float64{
	expMatrix := makeMatrix(len(x),1)
	for i:=0;i<len(x);i++{
		for j:=0;j<len(x[0]);j++{
			expMatrix[i][0] += math.Exp(x[i][j])
		}
	}
	ans := makeMatrix(len(x),len(x[0]))
	for i:=0;i<len(x);i++{
		for j:=0;j<len(x[0]);j++{
			ans[i][j] = math.Exp(x[i][j]) / expMatrix[i][0]
		}
	}
	return ans
}
//クロスエントロピー誤差
func crossEntropy(y [][]float64,t [][]float64) float64{
	n := 0.0
	for i:=0;i<len(y);i++{
		for j:=0;j<len(y[0]);j++{
			n += t[i][j] * math.Log(y[i][j])// + (1 - t[i][j]) * math.Log(1 - y[i][j])
		}
	}
	return -n/float64(len(y))
}
//二乗誤差
func 	square(y [][]float64,t [][]float64) float64 {
	n := 0.0
	for i:=0;i<len(y);i++{
		for j:=0;j<len(y[0]);j++{
			n += (t[i][j] - y[i][j]) * (t[i][j] - y[i][j])
		}
	}
	return n/float64(len(y) * len(y[0]))
}
//乱数を発生
func random(min, max float64,seed int) float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()*(max-min) + min
}
//i*jの中身0～1の乱数の行列を生成
func makeWight(i int,j int) [][]float64{
	ans := makeMatrix(i,j)
	for ii:=0;ii<i;ii++{
		for jj:=0;jj<j;jj++{
			ans[ii][jj] = rand.NormFloat64()
		}
	}
	return ans
}

var w1_1,w1_0,w2_1,w2_0,layer_z1,layer_a1,layer_z2,layer_a2,dw1,db1,dw2,db2 [][]float64
var costList []float64
//重みの初期化
func initValue(){
	//中間層のノードの数
	midNode := 2
	//データの次元数
	dim := 2
	//出力層のノード数
	finNode := 2
	//入力層から中間層への重み
	w1_1 = makeWight(dim,midNode)
	//入力層から中間層へのかけ合わせないやつ
	w1_0 = makeWight(1,midNode)
	//中間層から出力層への重み
	w2_1 = makeWight(midNode,finNode)
	//中間層から出力層へのかけ合わせないやつ
	w2_0 = makeWight(1,finNode)
}
//順伝搬
func forward(data [][]float64) [][]float64{
	layer_z1 = addVector(multi(data,w1_1),w1_0)
	layer_a1 = sigmoid(layer_z1)
	layer_z2 = addVector(multi(layer_a1,w2_1), w2_0)
	layer_a2 = softmax(layer_z2)
	return layer_a2
}

//コスト
func cost(x [][]float64,y [][]float64) float64{
	return square(x,y)
}

//逆伝搬
func back(x [][]float64,y [][]float64){
	//中間層-出力層の重みでの微分を求める
	output_delta := adaMul(adaMul(add(layer_a2,constMult(-1,y)),layer_a2),constAdd(1.0,constMult(-1,layer_a2)))
	dw2 = multi(trans(layer_a1),output_delta)
	db2 = vecMul(output_delta)

	/*
	delayer2 := add(y,constMult(-1,layer_a2))
	dw2 = multi(trans(layer_a1),delayer2)
	db2 = vecMul(delayer2)
	*/
	//入力層-中間層の重みでの微分を求める
	//fmt.Println(layer_z1)
	//fmt.Println(w1_1)
	//fmt.Println(delayer2)
	mid_delta := adaMul(multi(output_delta,trans(w1_1)),diffSigmoid(layer_a1))
	dw1 = multi(trans(x),mid_delta)
	db1 = vecMul(mid_delta)
	/*
	fmt.Println("dw2",dw2)
	fmt.Println("db2",db2)
	fmt.Println("dw1",dw1)
	fmt.Println("db1",db1)
	 */
}

//重みの更新
func update(alpha float64){
	w1_1 = add(w1_1,constMult(-alpha,dw1))
	w1_0 = add(w1_0,constMult(-alpha,db1))
	w2_1 = add(w2_1,constMult(-alpha,dw2))
	w2_0 = add(w2_0,constMult(-alpha,db2))
}

func train(x [][]float64,y [][]float64,alpha float64,epoc int){
	initValue()
	for i:=0;i<epoc;i++{
		forward(x)
		back(x,y)
		update(alpha)
		if epoc % 10 == 0{
			costList = append(costList,cost(forward(x),y))
		}
	}
}

func main(){
	x := [][]float64{{1,1},{0,0},{0,1},{1,0}}
	y := [][]float64{{1,0},{1,0},{0,1},{0,1}}
	train(x,y,0.1,10000)
	//mat := multi(trans(x),y)
	fmt.Println(costList)
	fmt.Println(forward(x))
	//fmt.Println("----------------")
	//x = [][]float64{{0.9,0.1},{0.9,0.1},{0.1,0.9},{0.1,0.9}}
	//y = [][]float64{{1,0},{1,0},{0,1},{0,1}}
	//fmt.Println(relu(y))
  //コストの情報
	file, _ := os.Create(`C:\\pythontest\\grade.txt`)
	
	defer file.Close()

	for i:=0;i<len(costList);i++{
		output := []byte(strconv.FormatFloat(costList[i], 'f', 15, 64) + " ")
		file.Write(([]byte)(output))
	}


}
