package main

import (
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	/** OS **/
	f, err := os.Open("foo")
	if err != nil {
		// エラーを出力しつつ、プログラムを終了する(os.Exit(1))
		//log.Fatal(err)
	}
	// コマンドラインの引数受け取り
	fmt.Printf("length=%d\n", len(os.Args))
	for _, v := range os.Args {
		fmt.Println(v)
	}
	//f, err = os.Open("test.txt")
	// 第2引数で指定したフラグで、第3引数で指定したパーミッションでファイルをオープン
	// 第2引数で指定するフラグは、複数の指定を「|」で繋いで指定したりもする
	f, err = os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	// 関数の終了時に確実にクローズするための処理
	defer f.Close()

	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fi, err := f.Stat()
	fmt.Printf("File name is %s, size is %d, bytes: %d\n", fi.Name(), fi.Size(), n)

	// ファイルの新規作成
	// os.Create("tmp.txt")
	fn, _ := os.OpenFile("tmp.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	fi, _ = fn.Stat()
	fmt.Printf("Created file name is %s, size is %d\n", fi.Name(), fi.Size())
	fn.Write([]byte("Hello, World!\n"))
	fn.Seek(0, os.SEEK_END) // ファイルの末尾にオフセットを移動
	fn.Write([]byte("Yeah!"))
	defer fn.Close()

	// ファイルの削除
	// 他のファイルやディレクトリ含む全てを削除する場合は、RemoveAll
	err = os.Remove("tmp.txt")

	// ファイル名の変更や移動: os.Rename
	// ディレクトリ操作
	dir, err := os.Getwd()
	fmt.Printf("current dir is %s\n", dir)
	// os.Chdirでカレントディレクトリ変更
	// os.Mkdirでディレクトリの作成
	// os.Symlinkでシンボリックリンクの操作、os.Readlinkでリンク先を読み込む
	host, err := os.Hostname()
	fmt.Printf("Hostname is %s\n", host)
	for _, v := range os.Environ() {
		fmt.Println(v)
	}
	// 環境変数の存在をチェックしつつ参照する場合
	if gopath, ok := os.LookupEnv("GOPATH"); ok {
		fmt.Printf("This env's go path is %s\n", gopath)
	}

	/** Time **/
	t := time.Date(2020, 7, 24, 9, 0, 0, 0, time.Local)
	fmt.Println(t)
	fmt.Println(t.Zone())
	fmt.Println(t.Weekday())
	tn := time.Now()
	tn = tn.Add(2*time.Minute + 15*time.Second)
	fmt.Println(tn)
	// オリンピックまでの時間
	fmt.Println(t.Sub(tn))
	// tの時間はtnの時間より前か。後かはAfterを使う。同じかはEqualを使う。
	fmt.Println(t.Before(tn))
	// 時間の増減
	fmt.Println(t.AddDate(5, -2, 1))
	// 文字列から時刻を生成
	ts, err := time.Parse(time.RFC3339, "2018-10-10T10:10:10+09:00")
	fmt.Println(ts)
	// 時刻から文字列
	fmt.Println(t.Format(time.RFC3339))
	// Unit time
	fmt.Println(t.Unix())
	// time.After()を使って、タイムアウト処理を実現もできるみたい

	/** math **/
	// Max&Min, Trunc&Floor(切り捨て)&Ceil(切り上げ)
	/** math/rand **/
	rand.Seed(time.Now().UnixNano())
	rand.Intn(100)

	/** flag **/
	// 引数やオプションなどの処理を効率的にやれるパッケージ
	var (
		max int
		msg string
		x bool
	)
	// コマンドラインオプションの定義
	flag.IntVar(&max, "n", 32, "処理の最大値")
	flag.StringVar(&msg, "m", "", "処理メッセージ")
	flag.BoolVar(&x, "x", false, "拡張オプション")
	flag.Parse()
	fmt.Println("処理の最大数 = ", max)
	fmt.Println("処理メッセージ = ", msg)
	fmt.Println("拡張オプション = ", x)

	/** fmt **/
	// よく使うやつ。
	str := fmt.Sprintf("[rand]=%d", rand.Intn(10))
	fmt.Println(str)
	// %vは配列やスライス、マップなどで有用。%+vでフィールドが出力される

	/** log **/
	log.SetPrefix("[LOG]")
	log.Println(str)

	/** strconv **/
	// 文字列に変換系のライブラリ
	_, err = strconv.ParseInt("57777", 10, 0)
	if err != nil {
		log.Fatalln(err)
	}
	/** strings **/
	// 文字列操作(検索、結合、置換処理など)がまとめられたもの
	fmt.Println(strings.Join([]string{"A", "B", "C"}, "|"))
	fmt.Println(strings.Index("hogehoge", "hoge"))
	fmt.Println(strings.HasPrefix("[INFO]hgoehoge", "["))
	fmt.Println(strings.Contains("hogehoge", "hoge"))
	fmt.Println(strings.Count("hogehoge", "hoge"))
	fmt.Println(strings.Replace("hogehoge", "hoge", "fuga", -1))
	fmt.Println(strings.Split("A,B,C,D,E", ","))
	fmt.Println(strings.TrimSpace("h o geh o g e")) // タブや改行コードも含むことに注意

	/** io/ioutil **/
	// ioutil.ReadAllとかでAPIのレスポンスを全て[]byteにして読み込むとかはあるみたい

	/** regexp **/
	// 正規表現。regexp.MustCompileとかでコンパイル時に正規表現をコンパイルして起きつつ、正規表現が正しいことを保証するのは良き
	re := regexp.MustCompile(`^[XYZ]+$`)
	fmt.Println(re.MatchString("hogehoge"))
	fmt.Println(re.FindString("XYZXYZ"))
	// ReplaceAllStringで文字列の置換
	re = regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	s := `
00-1111-2222
3333-44-5555
0120-114-114
`
	ms := re.FindAllStringSubmatch(s, -1)
	for _, v := range ms {
		fmt.Println(v)
	}
	fmt.Println(re.ReplaceAllString("Tel:000-111-222", "$3-$2-$1"))

	/** json **/
	u := new(User)
	u.Id = 1
	u.Name = "Alice"
	u.Email = "alice@example.com"
	u.Created = time.Now()
	bs, err = json.Marshal(u)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
	err = json.Unmarshal(bs, u)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", u)

	/** net/url **/
	urlTest, err := url.Parse("https://tree.taiga.io/project/beta-yumatsud-nekusutochiyarenzinixiang-kete/us/74?milestone=212187")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("scheme:%s, host:%s, Query:%+v\n", urlTest.Scheme, urlTest.Host, urlTest.Query())

	/** net/http **/
	res, err := http.Get("https://www.yahoo.co.jp/")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Status code: %d, Content-type: %v\n", res.StatusCode, res.Header["Content-Type"])
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))
	// url.Values{} -> vs.Add(key, value) -> vs.Encode()
	// http.PostForm(url, vs)
	// http.Post(url, "image/jpeg", io.Reader型)
	// より詳細には別本でやると思うので、ここでは一旦この程度で

	/** sync **/
	// 非同期処理などをまとめたパッケージ
	// sync.Mutex型で、排他制御が可能
	// sync.WaitGroupで任意のゴルーチンによる処理の完了を待ち受けるための仕組みを提供
	// -> Addメソッドでゴルーチンの数を指定し、Doneメソッドが同数実行されるまで、Waitメソッドでまつ。

	/** crypto/* **/
	s256 := sha256.New()
	io.WriteString(s256, "ABCDE")
	fmt.Printf("%x\n", s256.Sum(nil))
}

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Created time.Time `json:"createdAt"`
}
