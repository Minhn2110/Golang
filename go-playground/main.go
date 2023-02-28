package main

import (
	"fmt"
	"strings"
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
	// client := resty.New()
	// payload := PostMessage{
	// 	RoomId: "w8Fb2hTxT2YfFuWPK",
	// 	Text:   "Test noti!",
	// }
	// resp, err := client.R().
	// 	EnableTrace().
	// 	SetHeaders(map[string]string{
	// 		"X-Auth-Token": "raJcsc2SGTk2Xnanbgp-jrUIZRRGlI13kLH1PyRlWgw",
	// 		"X-User-Id":    "a4nWGM3qT3chmd9Et",
	// 	}).
	// 	SetBody(payload).
	// 	Post("https://chat.smartosc.com/api/v1/chat.postMessage")

	// // Explore request info

	// fmt.Println("Request Info:")
	// fmt.Println("  Header     :", resp.Header())

	// // Explore response object
	// fmt.Println("Response Info:")
	// fmt.Println("  Error      :", err)
	// fmt.Println("  Status Code:", resp.StatusCode())
	// fmt.Println("  Status     :", resp.Status())

	// fmt.Println("  Body       :\n", resp)
	// fmt.Println()

	authHeader := "Bearer eyJraWQiOiJZU2VKeUhzV2hmS08rcU00TVMzalRYR2Z6WXB2dzBmdnZreTg0NXk2REZzPSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiJiZTkzODdkMy01YzU2LTQ2YTUtOGU0ZC01YWY2MTViOGQ2OGQiLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb21cL3VzLWVhc3QtMV8yeGV0SHZCalAiLCJjb2duaXRvOnVzZXJuYW1lIjoiYmlnY29tbWVyY2V8Ym98ZXNlZWc5NG05NXwyMDUyMzY3IiwiY3VzdG9tOnRlbmFudF9pZCI6ImVhMzhmYWQ1LTRmM2EtNDI4Ni1hNmI3LTg2MzAyM2YwNmUyOSIsImN1c3RvbTpwcm9kdWN0X2NvZGUiOiJCTyIsImF1ZCI6IjYzbjVnZ3AwaGxla3RhNjlyaWwwaHAxZmkiLCJldmVudF9pZCI6IjYyNGNkM2Y0LThkOTgtNDIwOC1hNjY3LWM1NmM5YmU5NWQzNSIsInRva2VuX3VzZSI6ImlkIiwiYXV0aF90aW1lIjoxNjc0MTI0NDI5LCJleHAiOjE2NzQxMjgwMjksImlhdCI6MTY3NDEyNDQyOSwiZW1haWwiOiJtaW5obmhhQHNtYXJ0b3NjLmNvbSJ9.BGMoehX2pM7IHqDBJ3IwnpKxLEOKbkDGYK8iKqVRhAEjLFUCj_m6nMA0MlgxybgKmpvIUPaRSTLAARDD8lPen9v8yyodHMBd5tzXWVjK-oKNHpr82lu5JcM9kJcYbi_f7qNYcESqQGKkZEoRRng2Ubbe6SXkbk26lX_GC3ak73Kz69tk9wr5O_OHNY9oJ97S5tSshGtiYy06jfkXsX0vVRjFe4-6OlCzCwkq0qkIQqDW_RO5ei-KvksEdMHM4_hBXq51KyiGXKBMSAIo5reG3KlyYVyV55jz4UQxJx7Kn0YTGoVmARqcp0PmYsg6QOD0FXdusUntusUHgd-6s2p4Hg"
	splitToken := strings.Split(authHeader, "Bearer")
	// token := strings.TrimSpace(splitToken[1])

	fmt.Println(splitToken[1])
}
