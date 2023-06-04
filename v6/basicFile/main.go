package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("愛はあるんか?\n")
	err := ioutil.WriteFile("data1", data, 0644)
	if err != nil {
		panic(err)
	}

	read1, _ := ioutil.ReadFile("data1")
	fmt.Println(string(read1))

	file1, _ := os.Create("data2") // ファイルを作成する
	defer file1.Close()

	bytes, _ := file1.Write(data) // ファイルに書き込み
	fmt.Println("Wrote ", bytes)

	file2, _ := os.Open("data2") // ファイルの構造体を取得する
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2) // file2にデータbyte数だけ読み込みを実行する
	fmt.Println(string(read2))
}
