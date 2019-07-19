package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"time"
)

var (
	baseURL string

	campaign          string
	content           string
	medium            string
	source            string
	funnel_experiment string
	funnel_variation  string

	lang    string
	os      string
	product string

	hmacKey string

	numUrls int
)

const randCharset = "abcdefghijklmnopqrstuvwxyz"

func randomString(length int) string {
	res := make([]byte, length)

	for i := range res {
		res[i] = randCharset[rand.Intn(len(randCharset))]
	}

	return string(res)
}

func init() {
	rand.Seed(time.Now().UnixNano())

	flag.StringVar(&baseURL, "baseurl", "https://stubattribution-default.stage.mozaws.net", "base stub attribution service url")

	flag.StringVar(&campaign, "campaign", "testcampaign", "campaign")
	flag.StringVar(&content, "content", "testcontent", "content")
	flag.StringVar(&medium, "medium", "testmedium", "medium")
	flag.StringVar(&source, "source", "mozilla.com", "source")
	flag.StringVar(&funnel_experiment, "funnel_experiment", "exp1", "funnel_experiment")
	flag.StringVar(&funnel_variation, "funnel_variation", "var1", "funnel_variation")

	flag.StringVar(&lang, "lang", "en-US", "")
	flag.StringVar(&os, "os", "win", "")
	flag.StringVar(&product, "product", "test-stub", "")

	flag.StringVar(&hmacKey, "hmackey", "testkey", "test hmac key")

	flag.IntVar(&numUrls, "numurls", 1, "adds a random string to campaign, generates number of urls specified")
}

func genCode() string {
	query := url.Values{}
	query.Set("campaign", campaign)
	query.Set("content", content)
	query.Set("medium", medium)
	query.Set("source", source)
	query.Set("funnel_experiment", funnel_experiment)
	query.Set("funnel_variation", funnel_variation)
	query.Set("timestamp", fmt.Sprintf("%d", time.Now().UTC().Unix()))

	b64Query := base64.URLEncoding.WithPadding('.').EncodeToString([]byte(query.Encode()))
	return b64Query
}

func hmacSig(code string) string {
	mac := hmac.New(sha256.New, []byte(hmacKey))
	mac.Write([]byte(code))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

func genURL(code, sig string) string {
	query := url.Values{}
	query.Set("attribution_code", code)
	query.Set("attribution_sig", sig)

	query.Set("lang", lang)
	query.Set("os", os)
	query.Set("product", product)

	u, err := url.Parse(baseURL)
	if err != nil {
		log.Fatal("Could not parse url:", err)
	}
	u.RawQuery = query.Encode()
	return u.String()
}

func main() {
	flag.Parse()

	originalCampaign := campaign
	for i := 0; i < numUrls; i++ {
		if i > 0 {
			campaign = originalCampaign + randomString(12)
		}
		code := genCode()
		sig := hmacSig(url.QueryEscape(code))
		fmt.Println(genURL(code, sig))
	}
}
