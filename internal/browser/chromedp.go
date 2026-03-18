package browser

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/rayfiyo/autoALCv4/internal/domain"
)

type LoginConfig struct {
	Headless      bool
	WaitingTimeMS int
	UserDataDir   string
}

func Login(ctx context.Context, creds domain.Credentials, cfg LoginConfig) error {
	if cfg.WaitingTimeMS <= 0 {
		cfg.WaitingTimeMS = 1200
	}
	waitDur := time.Duration(cfg.WaitingTimeMS) * time.Millisecond

	allocatorOptions := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", cfg.Headless),
		chromedp.UserDataDir(cfg.UserDataDir),
	)

	allocCtx, cancelAlloc := chromedp.NewExecAllocator(ctx, allocatorOptions...)
	defer cancelAlloc()

	taskCtx, cancelTask := chromedp.NewContext(allocCtx)
	defer cancelTask()

	var currentURL string
	err := chromedp.Run(taskCtx,
		chromedp.Navigate(domain.LoginPageURL),
		chromedp.Sleep(waitDur),
		chromedp.WaitVisible(`#AccountId`, chromedp.ByQuery),
		chromedp.SendKeys(`#AccountId`, creds.ID, chromedp.ByQuery),
		chromedp.Sleep(waitDur),
		chromedp.SendKeys(`#Password`, creds.Password, chromedp.ByQuery),
		chromedp.Sleep(waitDur),
		chromedp.Click(`#BtnLogin`, chromedp.ByQuery),
		chromedp.Sleep(waitDur),
		chromedp.Evaluate(`window.location.href`, &currentURL),
	)
	if err != nil {
		return fmt.Errorf("ログイン操作に失敗した: %w", err)
	}

	if strings.TrimSpace(currentURL) != domain.TopPageURL {
		return fmt.Errorf("ログイン後の遷移先 URL が不正である: got=%s want=%s",
			currentURL, domain.TopPageURL)
	}

	return nil
}
