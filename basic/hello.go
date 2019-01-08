package main

// f "fmt"とかにすると、fmtは使えずfのみ使えるようになる(ようは上書き)
// . "math"とするとパッケージ名の指定は不要
// ファイルが分割されていても、定数や変数といった要素はpackageが同じであれば参照可能だが、
// import宣言のみ特立していることに注意
import (
	"./animals"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
	fmt.Println(AppName())
	// %vは様々なデータ型のデータを埋め込める
	fmt.Printf("数値=%v 文字列=%v 配列=%v\n", 5, "Golang", [...]int{1, 2, 3})
	// %#vはGoのリテラル表現でデータを埋め込む
	// %Tはデータの型情報を埋め込む
	// 変数の明示的な指定(型と変数名の両方が必要)
	var n int
	var (
		id   int
		name string
	)
	// 変数の暗黙的な指定(型がない。正確には下記にすることで型の定義と値の代入をまとめてやっている)
	i := 1
	// この暗黙的な指定は、一度しか許されない(:=は変数の定義ということを忘れずに)
	// 変数への再代入は=を使い、こちらに再代入の制限はない
	n = 2
	name = "Golang"
	// Goではグローバル変数はないが、パッケージ変数というのものは定義可能
	// Goでは使用していない変数などはチェックされる
	fmt.Printf("n=%d, i=%d, id=%d, name=%s\n", n, i, id, name)
	// Goは暗黙的な型変換を許容しないことに注意(intとint64の暗黙的型変換など)
	nu := uint(17)
	bu := byte(n)
	fmt.Printf("nu=%T, bu=%T\n", nu, bu)
	// 整数型はオーバーフローに注意すること(心配な時はMathパッケージの定数などを利用してチェックするのが良さげ)
	// rune型: Unicodeコードポイントを表す特殊な整数型(int32の別名で定義されてる。シングルクォートを使う)
	ary1 := [3]int{}
	ary2 := [...]int{1, 2, 3}
	fmt.Printf("ary1=%V, ary2=%V\n", ary1, ary2)
	// interface{}型: あらゆる型と互換性があるもの。故にあらゆる型の値を代入可能。
	// ただし、何らかの演算の対象としては利用不可
	var x interface{}
	fmt.Printf("%#v\n", x)
	// +記号は文字列の結合に使えるおー
	// 関数
	// 戻り値を破棄する場合は、「_」を使うこと(引数には_を使うことは可能)
	var div1, div2 = div(ary2[1], ary2[2])
	fmt.Printf("商=%d, 剰余=%d\n", div1, div2)
	fmt.Println(sample1())
	// 無名関数
	// 関数を返す関数の定義や、関数を引数に取るなどに使えるやーつ
	f := func(x, y int) int { return x * y }
	fmt.Printf("f = %d\n", f(2, 3))
	returnFunc()()
	callFunction(func() {
		fmt.Println("I'm a function")
	})
	flater := later()
	fmt.Println(flater("Golang"))
	fmt.Println(flater("is"))
	fmt.Println(flater("awesome!"))
	// 定数はconstで定義、()を利用して複数まとめて定義も可能。その際、値を省略すると直前に定義したものが割り当てられる
	// 定数にはすでに定義された定数を利用することや、計算式を使うことも可能。あとは型指定も可能。
	const (
		X = 1
		y
		z
		I64 = int64(-1)
	)
	// 列挙型に近い振る舞いをiotaを使えば実現可能
	// iotaは0から始まる
	const (
		A = iota
		B
		C
	)
	// スコープ
	// パッケージに定義された定数、変数、関数などが他パッケージから参照可能かどうかは、
	// 識別子の1文字目が大文字であるかどうかで決まる(大文字だと他パッケージから参照可能)
	// パッケージで定義された識別子の参照はfoo.MAXのようにドットでつなぐ
	// 制御構文
	// ループはforのみらしい
	// ifの条件式は論理値である必要あり
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("i=%d\n", i)
	}
	// if [簡易文] ; [条件式]という書き方もある
	// if _, err := doSomething(); err != nil {}
	fruits := [3]string{"Apple", "Banana", "Cherry"}
	for i, s := range fruits {
		fmt.Printf("No.%d fruit is %s\n", i, s)
	}
	// switch: fallthrouth文で次のcaseにつなぐことができる
	// 型アサーション: interface{}型で隠蔽された型をチェックできる。x.(T)のような形。
	// var i, isInt := x.(int) とかで1番目には変数そのものが、2番目には変換できたかとうか
	// switch x.(type)でcase boolなどと使うことも可能
	// defer: 関数の終了時に実行される式を登録できる。複数登録可能で、あとで登録されたものから実行される(その際は無名関数を使うと良い)
	// panicはランタイムエラーで終了させるもの。ただし、deferに登録したものは呼ばれる。
	// それを利用して、recoverを使うと復帰させることも可能。
	// パッケージの初期化は、func init()を置いておけばOK
	// go(ゴルーチンはまた別でやる)
}

// 関数
// 基本構造は他言語と同じ感じなので、特記事項だけまとめておく
// 引数の型指定をまとめることができる(同じ型の場合のみ)
// 複数の戻り値を指定可能
func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

// Go言語は例外機構がないので、関数のエラーは下記のようにすることが多い
// result, err := doSomething()
// if (err != nil) {}
// 戻り値に変数を割り当てると、初期化とreturnへの指定しているのと同じになる
func sample1() (a int) {
	return
}

func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func callFunction(f func()) {
	f()
}

// クロージャー
func later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}
