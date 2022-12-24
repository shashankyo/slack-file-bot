package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4527391525223-4565514243988-jp9i0rLYWv9gFQz2EuGPccSz")
	os.Setenv("CHANNEL_ID", "C04FXQ3SH35")
	api := slack.New(os.Getenv("SLACK_BOT_token"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"20221215145037english.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name:%s. URL:%s\n", file.Name, file.URL)
	}
}
