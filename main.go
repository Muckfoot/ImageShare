package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/everdev/mack"
	"github.com/therecipe/qt/widgets"
)

var (
	path string
	root string = "/Users/rekfuki/Documents/Coding/ImageShare/"
)

func main() {
	logf, err := os.OpenFile(root+"bot-erros.log",
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	defer logf.Close()
	log.SetOutput(logf)

	widgets.NewQApplication(len(os.Args), os.Args)

	NewImageShareForm().Show()

	var config Configuration
	loadJSON(root+"config.json", &config)

	var auth Authentication
	loadJSON(root+"auth.json", &auth)

	auth.Access_Token, auth.Refresh_Token, auth.Expires_In = getAccessToken(auth)

	path = config.Path
	previousSS := config.PreviousSS
	log.Print(previousSS)
	log.Print(path)

	hasUpdated := false

	go func() {
		for {
			hasUpdated, previousSS = update(path, previousSS)
			if hasUpdated {
				upload(config.Path, previousSS, auth.Access_Token)
			}
			time.Sleep(time.Second * 1)
		}
	}()

	widgets.QApplication_Exec()

	// sc := make(chan os.Signal, 1)
	// signal.Notify(
	// 	sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	// <-sc

	config.PreviousSS = previousSS

	file, err := json.Marshal(config)
	checkErr(err)

	err = ioutil.WriteFile(root+"config.json", file, 0644)
	checkErr(err)
}

func update(path string, previousSS string) (bool, string) {
	files, err := ioutil.ReadDir(path)
	checkErr(err)

	latestSS := files[len(files)-1].Name()

	if latestSS != previousSS {
		previousSS = latestSS

		return true, previousSS
	}
	return false, previousSS

}

func getAccessToken(auth Authentication) (string, string, int64) {
	authUrl := "https://api.imgur.com/oauth2/token"
	resp, err := http.PostForm(authUrl,
		url.Values{
			"refresh_token": {auth.Refresh_Token}, "client_id": {auth.Client_Id},
			"client_secret": {auth.Client_Secret}, "grant_type": {"refresh_token"}})
	checkErr(err)

	var authResp AuthResponse
	err = json.NewDecoder(resp.Body).Decode(&authResp)
	checkErr(err)

	return authResp.Access_Token, authResp.Refresh_Token, authResp.Expires_In
}

func upload(path string, ssPath string, accessToken string) {
	file, err := ioutil.ReadFile(path + ssPath)
	checkErr(err)

	imgString := base64.StdEncoding.EncodeToString(file)

	authUrl := "https://api.imgur.com/3/image"
	req, err := http.NewRequest("POST", authUrl, strings.NewReader(imgString))
	checkErr(err)

	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	checkErr(err)

	var uploadResponse UploadResponse
	err = json.NewDecoder(resp.Body).Decode(&uploadResponse)
	if err != nil {
		switch err.Error() {
		case "EOF":
			upload(path, ssPath, accessToken)
			return
		default:
			checkErr(err)
			return
		}
	}
	// checkErr(err)

	if uploadResponse.Success {
		link := uploadResponse.Data.Link
		mack.SetClipboard(link)
		mack.Beep(1)
		err = mack.Notify(link, "Image has been uploaded to imgur")
		checkErr(err)
		updateHistoryList(link, ssPath, 1)
	} else {
		mack.Beep(1)
		err = mack.Notify("Image failed to upload")
		checkErr(err)
	}

}
