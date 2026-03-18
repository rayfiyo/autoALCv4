package usecase

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/rayfiyo/autoALCv4/internal/browser"
	"github.com/rayfiyo/autoALCv4/internal/domain"
	"github.com/rayfiyo/autoALCv4/utils"
)

type RunInput struct {
	DotEnvPath    string
	Headless      bool
	WaitingTimeMS int
	UserDataDir   string
}

func Run(input RunInput) error {
	creds, err := loadCredentials(input.DotEnvPath)
	if err != nil {
		return err
	}

	chromeDir, err := utils.ResolveChromeUserDataDir(input.UserDataDir)
	if err != nil {
		return fmt.Errorf("chrome ユーザーデータディレクトリ解決に失敗した: %w", err)
	}

	loginCfg := browser.LoginConfig{
		Headless:      input.Headless,
		WaitingTimeMS: input.WaitingTimeMS,
		UserDataDir:   chromeDir,
	}
	if err := browser.Login(context.Background(), creds, loginCfg); err != nil {
		return err
	}

	return nil
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
