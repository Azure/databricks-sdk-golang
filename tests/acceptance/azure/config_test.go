// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package azure_test

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
	databricks "github.com/Azure/databricks-sdk-golang"
	dbAzure "github.com/Azure/databricks-sdk-golang/azure"
)

var c *dbAzure.DBClient

const (
	DATABRICKS_HOST_KEY  = "DATABRICKS_HOST"
	DATABRICKS_TOKEN_KEY = "DATABRICKS_TOKEN"
)

func init() {
	_, hostSet := os.LookupEnv(DATABRICKS_HOST_KEY)
	_, tokenSet := os.LookupEnv(DATABRICKS_TOKEN_KEY)

	if !hostSet || !tokenSet {
		e := godotenv.Load()
		if e != nil {
			log.Fatalln("Failed to load environment variables.")
		}
	}

	opt := databricks.NewDBClientOption("", "", os.Getenv(DATABRICKS_HOST_KEY), os.Getenv(DATABRICKS_TOKEN_KEY), nil, false, 0)
	c = dbAzure.NewDBClient(opt)
}

func randSeq(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
