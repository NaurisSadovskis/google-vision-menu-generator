package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
	"golang.org/x/net/context"
)

var ctx context.Context

func isNonDescriptiveWord(w string) bool {
	var isBanned bool
	ignoredWords := []string{"food", "dish", "cuisine", "fried food", "meat", "fruit", "produce"}
	for _, word := range ignoredWords {
		if word == w {
			isBanned = true
		}
	}
	return isBanned
}

func generateLabels(entry string, menu *map[string][]string, client *vision.ImageAnnotatorClient) {

	file, err := os.Open("./menu/" + entry)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	defer file.Close()

	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	labels, err := client.DetectLabels(ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed to detect labels: %v", err)
	}

	var l []string
	for _, label := range labels {
		if isNonDescriptiveWord(label.Description) {
			continue
		}
		l = append(l, label.Description)
	}

	(*menu)[entry] = l
}

func main() {

	ctx = context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	menu := make(map[string][]string)

	files, err := ioutil.ReadDir("./menu")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		menu[f.Name()] = []string{""}
	}

	for entry, _ := range menu {
		generateLabels(entry, &menu, client)
	}

	fmt.Printf("***Google Vision API results***\n%s\n\n", menu)

	fmt.Println("***Today's Menu Consists Of***")
	fmt.Println("Todays menu consists of:")
	for _, v := range menu {
		fmt.Println("*", v[0])
	}
}
