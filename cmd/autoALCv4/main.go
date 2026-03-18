package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rayfiyo/autoALCv4/internal/usecase"
)

func main() {
	input, err := parseRunInput(os.Args[1:])
	if err != nil {
		exitWithError(err)
	}

	if err := usecase.Run(input); err != nil {
		exitWithError(err)
	}
}

func parseRunInput(args []string) (usecase.RunInput, error) {
	input := usecase.RunInput{
		DotEnvPath:    ".env",
		Headless:      true,
		WaitingTimeMS: 1200,
	}

	fs := flag.NewFlagSet("autoALCv4", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	fs.StringVar(&input.DotEnvPath, "dot-env", input.DotEnvPath, "env ファイルのパス")
	fs.BoolVar(&input.Headless, "headless", input.Headless, "ヘッドレスモードの有効/無効")
	fs.IntVar(&input.WaitingTimeMS,
		"waiting-time", input.WaitingTimeMS, "各処理の待ち時間 [ミリ秒]")
	fs.IntVar(&input.WaitingTimeMS,
		"wt", input.WaitingTimeMS, "各処理の待ち時間 [ミリ秒]")
	fs.StringVar(&input.UserDataDir,
		"user-data-dir", "", "Chrome ユーザーデータディレクトリ")

	if err := fs.Parse(args); err != nil {
		return usecase.RunInput{}, fmt.Errorf("フラグ解析に失敗した: %w", err)
	}
	return input, nil
}

func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
	os.Exit(1)
}
