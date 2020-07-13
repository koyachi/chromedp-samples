package main

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"

	"github.com/koyachi/chromedp-samples/samples"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()

	if err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate("https://github.com/koyachi"),
		samples.CaptureAction,
	}); err != nil {
		log.Fatal(err)
	}
}
