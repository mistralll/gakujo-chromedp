package scrape

import (
	"context"

	"github.com/chromedp/chromedp"
)

func GetHomeHTML(id string, pass string) (string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var ans string

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://gakujo.shizuoka.ac.jp/portal/"),
		chromedp.Click("#left_container > div.left-module-top.bg_color > div > div > a", chromedp.NodeVisible),
		chromedp.ScrollIntoView(`.wrapper > footer:nth-child(2)`),

		chromedp.SendKeys(`#username`, id),
		chromedp.SendKeys(`#password`, pass),
		chromedp.Click(`body > div > div > div > div > form > div:nth-child(3) > button`),

		chromedp.ScrollIntoView(`#footer-lower`),

		chromedp.OuterHTML("html", &ans, chromedp.ByQuery),
	)
	if err != nil {
		return ans, err
	}

	return ans, nil
}
