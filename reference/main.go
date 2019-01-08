package main

import (
	"fmt"
	"time"
)

func main() {
	// slice: いわゆる可変長配列
	s := make([]int, 5, 10) // 要素数5、容量10であるint型のスライス
	fmt.Println(s)
	s[0] = 100
	fmt.Printf("s[0]=%d, s[1]=%d\n", s[0], s[1])
	// スライスの要素数は関数len、容量は関数capが使える
	// 容量とはそこに入れれる大きさ。要素数はすでに入っている大きさと捉えてOKかも
	fmt.Printf("var s's length is %d, cap is %d\n", len(s), cap(s))
	// リテラル、つまり[]int{1, 2, 3}とかでも生成可能
	// 元のスライスから切り出してスライスを生成
	s1 := s[0:3]
	fmt.Println(s1)
	// 配列とスライスの大きな違いは拡張性
	s = append(s, 1000) // sの末尾に要素を追加(appendは= or := が伴う必要あり)
	fmt.Println(s)
	s2 := append(s, s1...) // スライス末尾に別のスライスの要素を追加
	fmt.Println(s2)
	// appendで元の容量より大きくなると、新たに確保されたメモリ領域を見るようになることに注意
	// copy(dist, src []T) int でコピー可能。返り値はコピーに成功した数。
	// 完全スライス式: a[low:hight:max], 0 <= low <= high <= max <= cap(a)
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(s...))
	a := []int{1, 2, 3}
	pow(a)
	fmt.Println(a)
	// 配列の一部からスライスを作成すると、そのスライスは元の配列と同じメモリを参照することにも注意
	// map: いわゆる連想配列
	// 関数型と参照型を除く、任意の型のキーと任意の型で作れるおー
	m := make(map[string]string)
	m["first_name"] = "Hogehoge"
	m["last_name"] = "Fugafuga"
	fmt.Println(m)
	// リテラル指定も可能
	m1 := map[int]string{
		1: "Taro",
		2: "Hanako",
		3: "Jiro", // 見通しをよくするために改行可能だが、最後のカンマを忘れずに
	}
	fmt.Println(m1)
	// マップの中にスライスを指定することも可能 map[int][]int{}
	// もちろんmapの中にmapを指定することも可能 map[int]map[int]stringなど。その際入れ子になった中では、
	// 1: {1: "hogehoge"}のように省略記法を使える
	// 下記はキーへの参照が存在するかどうかを2つ目の変数への代入で確認してる
	result, ok := m1[1]
	fmt.Printf("result=%s, ok=%V\n", result, ok)
	result, ok = m1[8]
	fmt.Printf("result=%s, ok=%V\n", result, ok)
	// if _,ok := m[9]; ok {} のような書き方は頻出するおー
	// mapとfor: 順序は不定みたい
	for k, v := range m1 {
		fmt.Printf("key=%d, value=%s\n", k, v)
	}
	fmt.Println(len(m1)) // 同じようにlenで要素数は観れるおー
	// キーを指定して要素を消せる(keyがなければ何もしない)
	delete(m1, 3)
	fmt.Println(m1)
	// mapもmake(map[int]string, 100)のように要素数を確保することもできるが、巨大なmapを使う際とかでOKみたい
	// チャネル: ゴルーチンとゴルーチンの間でデータの受け渡しを司るためのデータ構造
	// var ch chan int
	// <-chan: 受信専用チャネル
	// chan<-: 送信専用チャネル
	// 指定しなければ双方向のチャネル
	ch1 := make(chan int, 20) //2つ目の引数でバッファサイズを指定。指定しない場合はバッファサイズ0になる
	// チャネルは「キュー」の性質を備えるデータ構造(あくまでゴルーチン間)
	// バッファとはこのキューを格納する領域で、バッファサイズはキューを格納するサイズ。
	// 下記はチャネルに整数5を送信
	//ch1 <- 5
	// 下記はチャネルから整数値を受信
	//res := <-ch1
	// 送信で、チャネルのサイズに収まる限りゴルーチンは停止しない
	// 関数lenでチャネルのバッファに溜められているサイズを知れる
	// 関数capでチャネルのバッファサイズがわかる
	go receiver("1st goroutine", ch1)
	go receiver("2nd goroutine", ch1)
	go receiver("3rd goroutine", ch1)
	fmt.Println("ゴルーチンスタート")
	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}
	close(ch1)
	// goroutineの完了を3秒待つ
	time.Sleep(3 * time.Second)
	// select: 複数のチャネルに対する受信、送信処理ともにゴルーチンを停止させることなくコントロール可能
	// case説は全てチャネルへの処理である必要がある
	/*
		select {
			case e1 := <-ch1 // ch1からの受信が成功した場合の処理
			case e2,ok := <-ch2 // ch2からの受信が成功した場合の処理
			case ch3 <- e3:
			default: //
		}
	*/
}

// 下記は可変長引数をとる関数
// 可変長引数は引数の末尾に1つだけ定義可能
func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

// スライスの各要素を2乗する(参照型の例)
func pow(a []int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}

// ゴルーチン
func receiver(name string, ch <-chan int) {
	for {
		// i, ok := <-ch1 でチャネルがクローズしているかどうかわかる(正確にはバッファ内が空でかつクローズされた状態)
		i, ok := <-ch
		if ok == false {
			// 受信できなくなったら終了
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + " is done.")
}
