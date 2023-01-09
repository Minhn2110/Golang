package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type PostMessage struct {
	RoomId      string `json:"roomId"`
	Channel     string `json:"channel"`
	Text        string `json:"text"`
	Alias       string `json:"alias"`
	Emoji       string `json:"emoji"`
	Avatar      string `json:"avatar"`
	Attachments string `json:"attachments"`
}

func main() {
	client := resty.New()
	payload := PostMessage{
		RoomId: "w8Fb2hTxT2YfFuWPK",
		Text:   "Test noti!",
	}
	resp, err := client.R().
		EnableTrace().
		SetHeaders(map[string]string{
			"X-Auth-Token": "raJcsc2SGTk2Xnanbgp-jrUIZRRGlI13kLH1PyRlWgw",
			"X-User-Id":    "a4nWGM3qT3chmd9Et",
		}).
		SetBody(payload).
		Post("https://chat.smartosc.com/api/v1/chat.postMessage")

	// Explore request info

	fmt.Println("Request Info:")
	fmt.Println("  Header     :", resp.Header())

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())

	fmt.Println("  Body       :\n", resp)
	fmt.Println()
}
