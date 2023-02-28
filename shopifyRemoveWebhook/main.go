package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	goshopify "github.com/bold-commerce/go-shopify"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type License struct {
	Code            string
	TenantID        string
	TierID          string
	ProductID       string
	Status          string
	IsTrial         int
	StartDate       string
	EndDate         string
	LastTierID      string
	LastTierEndDate string
	NextTier        string
}

type TenantPassport struct {
	Tenant            string
	TenantID          string
	CallbackURL       string
	IdentityURL       string
	SecretKey         string
	AccessKey         string
	ConsumerKey       string
	ConsumerSecretKey string
	Platform          string
	ProductID         string
	Status            string
	Metadata_Json     string
	Store_Hash        string
}

func main() {
	// metadata_json->>\"$.store_url\"

	// for index, item := range tentantPassport {
	// 	// fmt.Println("index", index)
	// 	// fmt.Println("item", item)
	// 	if index < 5 {
	// 		DeregisterWebHook(item.Store_Hash, item.AccessKey)
	// 	}
	// }
	test, _ := VerifyAuthorizationURL()
	fmt.Println("test", test)
}

// Verifying URL callback parameters.
func VerifyAuthorizationURL() (bool, error) {
	u := &url.URL{RawQuery: "code=513cf480dd258fec6d4a87b5ddf18e17\u0026embedded=1\u0026hmac=a608efa55e9a103ffe674424ba2fd82ed005255e6d6c1cda53e77824cc7edb9c\u0026host=bWluaHRlc3QyMTEwMS5teXNob3BpZnkuY29tL2FkbWlu\u0026locale=en\u0026session=20b31ffb6c6aef60b4bf58d7ee3a0f7f3ecc530d4a27ec5c5844f3579ab16588\u0026shop=minhtest21101.myshopify.com\u0026state=bWluaHRlc3QyMTEwMS5teXNob3BpZnkuY29tfGF1dC1zaG9waWZ5fDIwMjMtMDItMjEgMDQ6NDQ6MjguNzg3MDkyMDc3ICswMDAwIHV0YyBtPSsxMjg3LjgzNjQ5NTcxMg%3D%3D\u0026timestamp=1676963685"}
	q := u.Query()
	messageMAC := q.Get("hmac")

	// Remove hmac and signature and leave the rest of the parameters alone.
	q.Del("hmac")
	q.Del("signature")

	message, err := url.QueryUnescape(q.Encode())
	fmt.Println("message", message)
	fmt.Println("messageMAC", messageMAC)

	return VerifyMessage(message, messageMAC), err
}

func VerifyMessage(message, messageMAC string) bool {
	mac := hmac.New(sha256.New, []byte("ca313e0114a82b5046ac70215104e297"))
	mac.Write([]byte(message))
	expectedMAC := mac.Sum(nil)
	fmt.Println("message", messageMAC)
	// shopify HMAC is in hex so it needs to be decoded
	actualMac, _ := hex.DecodeString(messageMAC)

	fmt.Println("expectedMAC", expectedMAC)
	fmt.Println("actualMac", actualMac)

	return hmac.Equal(actualMac, expectedMAC)
}

func GetShopifyAccessKeyAndStoreHash() ([]TenantPassport, error) {
	dsn := "root:Anhminh96@@tcp(127.0.0.1:3306)/warthog?charset=utf8mb4&parseTime=True&loc=Local"
	connectFailed := "Could not connect the data base"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})

	if err != nil {
		fmt.Println("error", err)
		panic(connectFailed)
	}

	var tentantPassport []TenantPassport
	db = db.Table("licenses")
	db = db.Select("access_key", "JSON_UNQUOTE(JSON_EXTRACT(metadata_json, '$.store_hash')) as store_hash")
	db = db.Joins("INNER JOIN tenant_passports ON tenant_passports.tenant_id = licenses.tenant_id")
	db = db.Joins("INNER JOIN tenants ON tenants.id = tenant_passports.tenant_id")
	db = db.Where("code LIKE ? AND licenses.deleted_at = '' AND tenant_passports.deleted_at = '' ", "AUT-Shopify%")
	db.Find(&tentantPassport)

	if len(tentantPassport) > 1 {
		return tentantPassport, nil
	}
	return nil, err
}

func DeregisterWebHook(store_hash string, access_key string) {
	var params interface{}
	atom8ShopifyApp := goshopify.App{
		ApiKey:    "abcd",
		ApiSecret: "a",
	}
	client := goshopify.NewClient(atom8ShopifyApp, store_hash, access_key)
	webhooks, err := client.Webhook.List(params)
	if len(webhooks) > 0 {
		for _, webhook := range webhooks {
			if strings.Contains(webhook.Address, "grit.software") {
				fmt.Println("webhook", webhook)
				// client.Webhook.Delete(webhook.ID)
			}
		}
	}

	if err != nil {
		fmt.Println(err)
	}

}
