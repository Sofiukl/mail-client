package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"

	"github.com/spf13/cobra"
)

const (
	// URL - notification server
	URL         = "http://localhost:8000/notification/N2"
	contentType = "application/json"
	httpMethod  = "POST"
)

// HelloCommand - hello command
var HelloCommand = &cobra.Command{
	Use: "hello",
	Run: func(cobraCmd *cobra.Command, args []string) {
		sub, err := cobraCmd.Flags().GetString("subject")
		body, _ := cobraCmd.Flags().GetString("body")
		notificationTemplate, _ := cobraCmd.Flags().GetString("templatefile")
		fmt.Println("Sending notification...", sub, body, notificationTemplate)
		bodyHTML, _ := ParseTemplate(notificationTemplate, DataSource{"sample name"})
		requestData := Notification{[]string{"sofikul.mallick786@gmail.com"},
			sub,
			bodyHTML}
		requestJSON, _ := json.Marshal(requestData)
		res, err := http.Post(URL, contentType, bytes.NewBuffer(requestJSON))
		if err != nil {
			fmt.Println("Error occurred.. ", err)
		} else {
			data, _ := ioutil.ReadAll(res.Body)
			fmt.Println("Sent mail.. ", string(data))
		}
	},
}
