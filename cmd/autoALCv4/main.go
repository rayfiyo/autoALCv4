package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/rayfiyo/autoALCv4/internal/browser"
	"github.com/rayfiyo/autoALCv4/internal/domain"
	"github.com/rayfiyo/autoALCv4/utils"
)

func main() {
	var dotEnvPath string
	var headless bool
	var waitingTimeMS int
	var userDataDir string

	flag.StringVar(&dotEnvPath, "dot-env", ".env", "env ファイルのパス")
	flag.BoolVar(&headless, "headless", true, "ヘッドレスモードの有効/無効")
	flag.IntVar(&waitingTimeMS, "waiting-time", 1200, "各処理の待ち時間 [ミリ秒]")
	flag.IntVar(&waitingTimeMS, "wt", 1200, "各処理の待ち時間 [ミリ秒]")
	flag.StringVar(&userDataDir, "user-data-dir", "", "Chrome ユーザーデータディレクトリ")
	flag.Parse()

	creds, err := loadCredentials(dotEnvPath)
	if err != nil {
		exitWithError(err)
	}

	chromeDir, err := utils.ResolveChromeUserDataDir(userDataDir)
	if err != nil {
		exitWithError(
			fmt.Errorf("chrome ユーザーデータディレクトリ解決に失敗した: %w", err),
		)
	}

	loginCfg := browser.LoginConfig{
		Headless:      headless,
		WaitingTimeMS: waitingTimeMS,
		UserDataDir:   chromeDir,
	}
	if err := browser.Login(context.Background(), creds, loginCfg); err != nil {
		exitWithError(err)
	}
}

func loadCredentials(dotEnvPath string) (domain.Credentials, error) {
	id := os.Getenv("ID")
	password := os.Getenv("PASSWORD")

	if id != "" && password != "" {
		return domain.NewCredentials(id, password)
	}

	envMap, err := utils.LoadDotEnv(dotEnvPath)
	if err != nil {
		return domain.Credentials{},
			fmt.Errorf("環境変数と env ファイルからクレデンシャルを取得できない: %w", err)
	}

	if id == "" {
		id = envMap["ID"]
	}
	if password == "" {
		password = envMap["PASSWORD"]
	}

	creds, credErr := domain.NewCredentials(id, password)
	if credErr != nil {
		if errors.Is(credErr, domain.ErrInvalidID) ||
			errors.Is(credErr, domain.ErrInvalidPassword) {
			return domain.Credentials{},
				fmt.Errorf("クレデンシャルが不正である: %w", credErr)
		}
		return domain.Credentials{}, credErr
	}

	return creds, nil
}

func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
	os.Exit(1)
}
