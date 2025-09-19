package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/oklog/ulid/v2"
)

func main() {
	// コマンドライン引数が渡されているかチェック
	if len(os.Args) < 2 {
		// エラーメッセージを標準エラー出力に表示
		fmt.Fprintln(os.Stderr, "エラー: 変換対象の16進エスケープ文字列を引数に指定してください。")
		fmt.Fprintln(os.Stderr, `使用法例: BytesToStrings '"\x01\x84\x08\x71\x18\x76\x08\x99\x84\x08\x25\x13\x12\x42\x31\x01"'`)
		os.Exit(1) // エラーで終了
	}

	// 引数を取得し、エスケープシーケンスを解釈
	// コマンドラインから渡される文字列は、Goの文字列リテラルとして解釈するために
	// ダブルクォートで囲まれている必要があります。
	// strconv.Unquoteは、そのような引用符付きの文字列を解釈して、
	// 実際のバイトシーケンスに変換します。
	inputStr := os.Args[1]
	binaryStr, err := strconv.Unquote(inputStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "エラー: 引数の形式が正しくありません。文字列をダブルクォートで囲み、有効なエスケープシーケンスを使用してください。")
		fmt.Fprintf(os.Stderr, "詳細: %v\n", err)
		os.Exit(1)
	}

	// 解釈後のデータが16バイトであることを確認 (ULIDは16バイト)
	byteData := []byte(binaryStr)
	if len(byteData) != 16 {
		fmt.Fprintf(os.Stderr, "エラー: 入力データは16バイトである必要がありますが、現在 %d バイトです。\n", len(byteData))
		os.Exit(1)
	}

	// バイトデータをULID型に変換
	var ulidValue ulid.ULID
	copy(ulidValue[:], byteData) // バイトスライスからULID型の配列に中身をコピー

	// ULIDを文字列としてコンソールに表示
	fmt.Println(strings.ToLower(ulidValue.String()))
}
