package getmnist

import (
	"fmt"
	"os"
)

func GetTestLabel() []int{
	// ファイルをOpenする
	f, err := os.Open("C:\\hoge\\hoge\\t10k-labels-idx1-ubyte")
	// 読み取り時の例外処理
	if err != nil{
		os.Exit(1)
	}
	// 関数が終了した際に確実に閉じるようにする
	defer f.Close()

	// バイト型スライスの作成
	buf := make([]byte, 10008)
	// nはバイト数を示す
	n, err := f.Read(buf)
	// バイト数が0になることは、読み取り終了を示す
	if n == 0{
		os.Exit(1)
	}
	if err != nil{
		os.Exit(1)
	}
	label := make([]int,10000)
	// バイト型を数字に変換して出力
	for i:=8;i<len(buf);i++{
		label[i-8] = int(buf[i])
	}
	return label
}

func GetTrainLabel() []int{
	// ファイルをOpenする
	f, err := os.Open("C:\\hoge\\hoge\\train-labels-idx1-ubyte")
	// 読み取り時の例外処理
	if err != nil{
		fmt.Println("error")
	}
	// 関数が終了した際に確実に閉じるようにする
	defer f.Close()

	// バイト型スライスの作成
	buf := make([]byte, 60008)
	// nはバイト数を示す
	n, err := f.Read(buf)
	// バイト数が0になることは、読み取り終了を示す
	if n == 0{
		os.Exit(1)
	}
	if err != nil{
		os.Exit(1)
	}
	//データを格納するスライス
	label := make([]int,60000)
	// バイト型を数字に変換してスライスに入れる
	for i:=8;i<len(buf);i++{
		label[i-8] = int(buf[i])
	}
	return label
}

func GetTestImg() [][][]int{
	// ファイルをOpenする
	f, err := os.Open("C:\\hoge\\hoge\\t10k-images-idx3-ubyte")
	// 読み取り時の例外処理
	if err != nil{
		fmt.Println("error")
	}
	// 関数が終了した際に確実に閉じるようにする
	defer f.Close()

	// バイト型スライスの作成
	buf := make([]byte, 10000 * 28 * 28 + 16)
	// nはバイト数を示す
	n, err := f.Read(buf)
	// バイト数が0になることは、読み取り終了を示す
	if n == 0{
		os.Exit(1)
	}
	if err != nil{
		os.Exit(1)
	}
	imgs := make([][][]int,10000)
	counter := 16
	for i:=0;i<10000;i++{
		imgs[i] = make([][]int,28)
		for j:=0;j<28;j++{
			imgs[i][j] = make([]int,28)
			for k:=0;k<28;k++{
				imgs[i][j][k] = int(buf[counter])
				counter ++
			}
		}
	}
	return imgs
}
func GetTrainImg()  [][][]int{
	// ファイルをOpenする
	f, err := os.Open("C:\\hoge\\hoge\\train-images-idx3-ubyte")
	// 読み取り時の例外処理
	if err != nil{
		fmt.Println("error")
	}
	// 関数が終了した際に確実に閉じるようにする
	defer f.Close()

	// バイト型スライスの作成
	buf := make([]byte, 60000 * 28 * 28 + 16)
	// nはバイト数を示す
	n, err := f.Read(buf)
	// バイト数が0になることは、読み取り終了を示す
	if n == 0{
		os.Exit(1)
	}
	if err != nil{
		os.Exit(1)
	}
	imgs := make([][][]int,60000)
	counter := 16
	for i:=0;i<60000;i++{
		imgs[i] = make([][]int,28)
		for j:=0;j<28;j++{
			imgs[i][j] = make([]int,28)
			for k:=0;k<28;k++{
				imgs[i][j][k] = int(buf[counter])
				counter ++
			}
		}
	}
	return imgs
}
