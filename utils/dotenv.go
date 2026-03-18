package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadDotEnv(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("env ファイルを開けない: %w", err)
	}
	defer func() {
		_ = file.Close()
	}()

	values := make(map[string]string)
	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, ok := strings.Cut(line, "=")
		if !ok {
			return nil, fmt.Errorf("env ファイルの形式が不正である: %s:%d", path, lineNum)
		}

		key = strings.TrimSpace(key)
		if key == "" {
			return nil, fmt.Errorf("env ファイルのキーが空である: %s:%d", path, lineNum)
		}
		values[key] = strings.TrimSpace(value)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("env ファイル読み取りに失敗した: %w", err)
	}

	return values, nil
}
