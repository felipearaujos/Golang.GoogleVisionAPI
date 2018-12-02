package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/vision/v1"

	Models "./Models"
)

const developerKey = "not today!"

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}

func ProcessImage(w http.ResponseWriter, r *http.Request) {
	var image Models.ProcessImageRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&image)
	if err != nil {
		http.Error(w, "Parsing error", http.StatusBadRequest)
		return
	}

	response, err := exec(image.Image)
	if err != nil {
		http.Error(w, "Execution error", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func exec(file string) (string, error) {

	enc := file
	img := &vision.Image{Content: enc}

	feature := &vision.Feature{
		Type: "SAFE_SEARCH_DETECTION",
	}

	req := &vision.AnnotateImageRequest{
		Image:    img,
		Features: []*vision.Feature{feature},
	}

	batch := &vision.BatchAnnotateImagesRequest{
		Requests: []*vision.AnnotateImageRequest{req},
	}

	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}
	svc, err := vision.New(client)
	if err != nil {
		log.Fatal(err)
	}
	res, err := svc.Images.Annotate(batch).Do()
	if err != nil {
		log.Fatal(err)
	}

	body, err := json.Marshal(res.Responses[0].SafeSearchAnnotation)

	response := string(body)

	return response, err
}

func handleError(err error, w http.ResponseWriter) bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return true
	}

	w.WriteHeader(http.StatusOK)
	return false
}

func main() {
	handlerRequest()
}
