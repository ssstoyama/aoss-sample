package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	search "github.com/opensearch-project/opensearch-go/v2"
	searchapi "github.com/opensearch-project/opensearch-go/v2/opensearchapi"
	requestsigner "github.com/opensearch-project/opensearch-go/v2/signer/awsv2"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context) (string, error) {
	endpoint := os.Getenv("AOSS_ENDPOINT")
	if endpoint == "" {
		return "", errors.New("AOSS_ENDPOINT is empty")
	}

	awsCfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", err
	}

	signer, err := requestsigner.NewSignerWithService(awsCfg, "aoss")
	if err != nil {
		return "", err
	}

	client, err := search.NewClient(search.Config{
		Addresses: []string{endpoint},
		Signer:    signer,
	})
	if err != nil {
		return "", err
	}

	indexReq := searchapi.IndexRequest{
		Index:      "test-index",
		DocumentID: "1",
		Body:       strings.NewReader(`{"title": "タイトル"}`),
	}
	_, err = indexReq.Do(ctx, client)
	if err != nil {
		return "", err
	}

	getReq := searchapi.GetRequest{
		Index:      "test-index",
		DocumentID: "1",
	}
	getRes, err := getReq.Do(ctx, client)
	if err != nil {
		return "", err
	}
	defer getRes.Body.Close()
	if getRes.StatusCode != http.StatusOK {
		return "", fmt.Errorf("StatusCode=%d", getRes.StatusCode)
	}
	body, err := ioutil.ReadAll(getRes.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
