package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/misikdmytro/granny-grams/client"
	"github.com/misikdmytro/granny-grams/config"
	"github.com/misikdmytro/granny-grams/processor"
)

func main() {
	log.Println("Starting Granny Grams")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	log.Println("Loading configuration")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Generating image")
	var oac client.OpenAPIClient
	if cfg.Environment == "development" {
		oac = client.NewFakeOpenAPIClient()
	} else {
		oac = client.NewOpenAPIClient(cfg.OpenAPIBaseURL, cfg.OpenAPIToken)
	}
	resp, err := oac.GenerateImage(ctx, "a holiday cat", 1, "1024x1024")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Downloading image")
	wc := client.NewWebClient()
	err = wc.DownloadImageToFile(ctx, resp.Data[0].URL, "cat.png")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Adding text to image")
	ip := processor.NewImageProcessor()
	err = ip.AddTextToImage(ctx, "cat.png", "cat-with-text.png", "Happy Holidays!")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Done")
}
