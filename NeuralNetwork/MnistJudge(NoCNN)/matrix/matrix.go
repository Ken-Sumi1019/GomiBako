package matrix

import
(
	"os"
	"math"
	"math/rand"
	"time"
)

//大きい値を出力
func Max(a float64,b float64) float64{
	if a > b{return a} else{return b}
}

//行列を生成
func MakeMatrix(i int,j int) [][]float64{
	ans := make([][]float64,i)
	for k:=0;k<i;k++{
		ans[k] = make([]float64,j)
	}
	return ans
}
//行列を合計
func Add(a [][]float64,b [][]float64) [][]float64{
	if (len(a) != len(b)) || (len(a[0]) != len(b[0])){
		str := "ふええ～おにいちゃーん、大きさが違う行列は加算できないよぉ～～"
		println(str)
		os.Exit(1)
	}
	ans := MakeMatrix(len(a),len(a[0]))
	for i:=0;i<len(a);i++{
		for j:=0;j<len(a[0]);j++{
			ans[i][j] = a[i][j] + b[i][j]
		}
	}
	return ans
}
//アダマール積
func AdaMul(a [][]float64,b [][]float64) [][]float64{
	if (len(a) != len(b)) || (len(a[0]) != len(b[0])){
		str := "ふええ～おにいちゃーん、大きさが違う行列は計算できないよぉ～～"
		println(str)
		os.Exit(1)
	}
	ans := MakeMatrix(len(a),len(a[0]))
	for i:=0;i<len(a);i++{
		for j:=0;j<len(a[0]);j++{
			ans[i][j] = a[i][j] * b[i][j]
		}
	}
	return ans
}
//行列の積
func Multi(a [][]float64,b [][]float64) [][]float64{
	if len(a[0]) != len(b){
		println("ふええ～おにいちゃーん、内積が計算できないよぉ～～")
		os.Exit(1)
	}
	ans := MakeMatrix(len(a),len(b[0]))
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
func ConstAdd(n float64,mat [][]float64) [][]float64{
	ans := MakeMatrix(len(mat),len(mat[0]))
	for i:=0;i<len(mat);i++{
		for j:=0;j<len(mat[0]);j++{
			ans[i][j] = mat[i][j] + n
		}
	}
	return ans
}
//行列を定数倍
func ConstMult(n float64,matrix [][]float64) [][]float64{
	ans := MakeMatrix(len(matrix),len(matrix[0]))
	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix[0]);j++{
			ans[i][j] = n * matrix[i][j]
		}
	}
	return ans
}
//行列の転置
func Trans(a [][]float64) [][]float64{
	ans := MakeMatrix(len(a[0]),len(a))
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
func AddVector(mat [][]float64,vec [][]float64) [][]float64{
	if len(vec) != 1 {
		println("ふええ～おにいちゃーん、ベクトルじゃないないよ～～")
		os.Exit(1)
	} else if len(mat[0]) != len(vec[0]){
		println("ふええ～おにいちゃーん、列の数が違うよ～～")
		os.Exit(1)
	}
	ans := MakeMatrix(len(mat),len(mat[0]))
	for i:=0;i<len(mat);i++{
		for j:=0;j<len(mat[0]);j++{
			ans[i][j] = mat[i][j] + vec[0][j]
		}
	}
	return ans
}
//要素1、大きさ1行n列のベクトルとの内積つまり、列方向の合計
func VecMul(mat [][]float64) [][]float64{
	ans := MakeMatrix(1,len(mat[0]))
	for i:=0;i<len(mat);i++{
		for j:=0;j<len(mat[0]);j++{
			ans[0][j] += mat[i][j]
		}
	}
	return ans
}

//シグモイド関数
func Sigmoid(x [][]float64) [][]float64{
	ans := MakeMatrix(len(x),len(x[0]))
	for i:=0;i<len(x);i++{
		for j:=0;j<len(x[0]);j++{
			ans[i][j] = 1 / (1 + math.Exp(-x[i][j]))
		}
	}
	return ans
}
//シグモイド関数の微分(出力された値を引数に入れる)
func DiffSigmoid(x [][]float64) [][]float64{
	ans := MakeMatrix(len(x),len(x[0]))
	for i:=0;i<len(x);i++{
		for j:=0;j<len(x[0]);j++{
			ans[i][j] = x[i][j] * (1.0 - x[i][j])
		}
	}
	return ans
}

//ソフトマックス関数
func Softmax(x [][]float64) [][]float64{
	expMatrix := MakeMatrix(len(x),1)
	for i:=0;i<len(x);i++{
		for j:=0;j<len(x[0]);j++{
			expMatrix[i][0] += math.Exp(x[i][j])
		}
	}
	ans := MakeMatrix(len(x),len(x[0]))
	for i:=0;i<len(x);i++{
		for j:=0;j<len(x[0]);j++{
			ans[i][j] = math.Exp(x[i][j]) / expMatrix[i][0]
		}
	}
	return ans
}

//二乗誤差
func 	Square(y [][]float64,t [][]float64) float64 {
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
func MakeWight(i int,j int) [][]float64{
	ans := MakeMatrix(i,j)
	for ii:=0;ii<i;ii++{
		for jj:=0;jj<j;jj++{
			ans[ii][jj] = rand.NormFloat64()
		}
	}
	return ans
}
