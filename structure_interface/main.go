package main

import (
	"encoding/json"
	"fmt"
)

// 構造体の定義 (構造体は値型の一種。各フィールドは型の初期値で初期化される)
type Point struct {
	X int
	Y int
}

// メソッド: 任意の型に特化した関数を定義する仕組み
// 下記はPoint型のメソッドRender
// 関数とは異なり、funcとメソッド名の間にレシーバの型とその変数名が必要。
// このレシーバの定義が必要な以外は通常の関数定義と同じ
// レシーバが異なれば同名のメソッドを定義することは可能
// メソッドは構造体以外も定義可能
func (p *Point) Render() {
	fmt.Printf("<%d, %d>\n", p.X, p.Y)
}

// 構造体に定義するメソッドのレシーバはポインタ型」にするのが基本原則みたい
// 型のコンストラクタ
// 下記のように先頭大文字で外部パッケージに公開している状態
func NewPoint(x, y int) *Point {
	p := new(Point)
	p.X = x
	p.Y = y
	return p
}

// errorインターフェースのメソッドの実装
type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}
func RaiseError() error {
	return &MyError{Message: "エラーが発生", ErrCode: 1234}
}

type Points []*Point

func (ps Points) ToString() string {
	str := ""
	for _, p := range ps {
		if str != "" {
			str += ","
		}
		if p == nil {
			str += "<nil>"
		} else {
			str += fmt.Sprintf("[%d,%d]", p.X, p.Y)
		}
	}
	return str
}

func main() {
	// ポインタ: Cで躓くあれです
	// 下記はint型のポイントですね、あぁそうですそうです
	// var p *int
	var i int
	p := &i
	fmt.Printf("%T\n", p)
	pp := &p
	fmt.Printf("%T\n", pp)
	// デリファレンス: ポインタ型が保持するメモリ上のアドレスを経由してデータ本体を参照するための仕組み
	i = 5
	fmt.Println("i =", i)
	*p = 10
	fmt.Println("i =", i)
	p_ary := &[3]int{1, 2, 3}
	pow(p_ary)
	fmt.Println(p_ary)
	// 実際にGolangでは、デリファレンスは下記のように簡易的にかけるおー
	a := [3]string{"Apple", "Banana", "Cherry"}
	p_string := &a
	fmt.Println(a[1])
	fmt.Println(p_string[1]) // デリファレンスしたものが表示される
	p_string[2] = "Grape"
	fmt.Println(a[2])
	fmt.Println(p_string[2]) // デリファレンスしたものが表示される
	// スライス、len、capなども配列へのポインタ型であればデリファレンスを省略可能
	fmt.Printf("type=%T, address=%p\n", p, p)
	// 構造体: オブジェクト指向でクラスやオブジェクトを定義するのと同じくらいGolangでは重要
	// type: aliasに使える予約語
	type (
		IntPair [2]int
		Strings []string
		AreaMap map[string][2]float64
	)
	pair := IntPair{1, 2}
	amap := AreaMap{"Tokyo": {35.689488, 139.691706}}
	fmt.Println(pair, amap)
	// type Callback func(i int) int のようにcallbackを定義するのに使うなどもあり
	// 注意点: 同じ型から派生した場合であっても、エイリアスの間には互換性が成り立たない
	var point Point
	// 参照や代入はドットでつなぐんだおー
	fmt.Printf("X=%d, Y=%d\n", point.X, point.Y)
	point = Point{X: 1, Y: 2}
	fmt.Printf("X=%d, Y=%d\n", point.X, point.Y)
	// 構造体の中に構造体を定義できる
	type Feed struct {
		Name   string
		Amount uint
	}
	type Animal struct {
		Name string
		Feed /*Feed*/ // ここは型のみ指定するとフィールド名が型になるので型名だけ書くのもあり
	}
	animal := Animal{
		Name: "Monkey",
		Feed: Feed{
			Name:   "Banana",
			Amount: 10,
		},
	}
	fmt.Printf("Name:%s, Feed name: %s, Feed amount: %d\n", animal.Name, animal.Feed.Name, animal.Feed.Amount)
	// フィールド名を省略して埋め込まれた構造体のフィールド名が一意に定まる場合に限り、中間のフィールド名を省略可能
	// 上記でいうと、「animal.Amount」など
	// これは、異なる構造体型に、共通の性質を持たせる」場合などに有用みたいよ
	// 再帰的な構造体の定義は禁止 (自身の構造体に自身のフィールドを持つ、互いに参照し合うなど)
	// 本来は、構造体型のポインタを直接生成する方が頻出するパターン
	swap(&point)
	fmt.Printf("After swap, X=%d, Y=%d\n", point.X, point.Y)
	point.Render()
	fmt.Println(NewPoint(10, 10))
	// new: 指定した型のポインタ型を生成するための組み込み関数
	type Person struct {
		Id   int
		Name string
		Area string
	}
	person := new(Person)
	fmt.Println(person.Id, person.Name, person.Area)
	// これは「&Person」とするのと同じなので、状況に応じて使い分けるで良いみたい
	// スライスと構造体の組み合わせはよく出る
	ps := make([]Point, 5)
	for _, p := range ps {
		fmt.Println(p.X, p.Y)
	}
	pss := Points{}
	pss = append(pss, &Point{X: 2, Y: 4})
	pss = append(pss, nil)
	pss = append(pss, &Point{X: 8, Y: 10})
	fmt.Println(pss.ToString())
	// mapと構造体の組み合わせ
	// リテラル内の構造体の型名を省略できる
	m1 := map[Person]string{
		{Id: 1, Name: "Taro", Area: "Tokyo"}:    "Soccer",
		{Id: 2, Name: "Hanako", Area: "Nagoya"}: "Tennis",
	}
	m2 := map[int]Person{
		1: {Id: 1, Name: "Taro", Area: "Tokyo"},
		2: {Id: 2, Name: "Hanako", Area: "Nagoya"},
	}
	fmt.Println(m1, m2)
	// Tag: フィールドにメタ情報を付与する機能
	// タグが文字列リテラルか、RAQ文字列リテラルのどちらかを使用可能
	type User struct {
		Id   int    `json:"user_id"`
		Name string `json:"user_name"`
		Age  uint   `json:"user_age"`
	}
	u := User{Id: 1, Name: "Taro", Age: 32}
	bs, _ := json.Marshal(u)
	fmt.Println(string(bs))
	// インターフェース
	// 任意の型がどのようなメソッドを実装するべきかを規定するための枠組み
	// errorは下記のように定義されている
	// type error interface {
	//   Error() string
	// }
	// interface { メソッドのシグネチャの列挙 }
	err := RaiseError()
	err.Error()
	// 型アサーションで本来の型を取得することも可能
	// e, ok := err.(*MyError)
	// インターフェースのメリットは異なる型を共通するインターフェース型にまとめることができる
	type Stringify interface {
		ToString() string
	}
	// このインターフェースをこれまでのUser型、Person型が実装していれば
	// vs := []Stringify{ &Persona{〜}, &User{〜} }のようにStringify型のデータとしてまとめることができる
}

// ポインタを利用したやーつ
func pow(p *[3]int) {
	i := 0
	for i < 3 {
		(*p)[i] = (*p)[i] * (*p)[i]
		i++
	}
}

// 構造体とポインタ
func swap(p *Point) {
	x, y := p.Y, p.X
	p.X = x
	p.Y = y
}
