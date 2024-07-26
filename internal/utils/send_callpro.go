package utils

import (
	"fmt"
	"net/http"
	"strings"
)

func SendMessage(phone, message string) error {
	message = strings.ReplaceAll(message, " ", "%20")
	url := fmt.Sprintf("https://api.messagepro.mn/send?to=%v&from=72887388&text=%v", phone, message)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-api-key", "e15b92d6da557174aeb74b29f5243f77")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return nil
	}

	fmt.Println("response Status:", resp)
	return fmt.Errorf("status code: %d", resp.StatusCode)
}
