package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ResolveChromeUserDataDir(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("~/（HOME ディレクトリ）を取得できない: %w", err)
		}
		trimmed = filepath.Join(homeDir, ".cache", "chrome-user-data")
	}

	if err := os.MkdirAll(trimmed, 0o755); err != nil {
		return "", fmt.Errorf("chrome ユーザーデータディレクトリを作成できない: %w", err)
	}
	return trimmed, nil
}
