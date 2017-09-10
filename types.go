package main

import "github.com/therecipe/qt/core"

type Configuration struct {
	Path       string `json:"path"`
	PreviousSS string `json:"previousSS"`
}
type Authentication struct {
	Client_Id        string `json:"client_id"`
	Client_Secret    string `json:"client_secret"`
	Access_Token     string `json:"access_token"`
	Expires_In       int64  `json:"expires_in"`
	Token_Type       string `json:"token_type"`
	Refresh_Token    string `json:"refresh_token"`
	Account_Username string `json:"account_username"`
	Account_id       int64  `json:"account_id"`
	Scope            string `json:"scope"`
}

type AuthResponse struct {
	Access_Token     string `json:"access_token"`
	Expires_In       int64  `json:"expires_in"`
	Token_Type       string `json:"token_type"`
	Scope            string `json:"scope"`
	Refresh_Token    string `json:"refresh_token"`
	Account_Id       int64  `json:"account_id"`
	Account_Username string `json:"account_username"`
}

type UploadResponse struct {
	Data struct {
		ID          string        `json:"id"`
		Title       interface{}   `json:"title"`
		Description interface{}   `json:"description"`
		Datetime    int           `json:"datetime"`
		Type        string        `json:"type"`
		Animated    bool          `json:"animated"`
		Width       int           `json:"width"`
		Height      int           `json:"height"`
		Size        int           `json:"size"`
		Views       int           `json:"views"`
		Bandwidth   int           `json:"bandwidth"`
		Vote        interface{}   `json:"vote"`
		Favorite    bool          `json:"favorite"`
		Nsfw        interface{}   `json:"nsfw"`
		Section     interface{}   `json:"section"`
		AccountURL  interface{}   `json:"account_url"`
		AccountID   int           `json:"account_id"`
		IsAd        bool          `json:"is_ad"`
		InMostViral bool          `json:"in_most_viral"`
		Tags        []interface{} `json:"tags"`
		AdType      int           `json:"ad_type"`
		AdURL       string        `json:"ad_url"`
		InGallery   bool          `json:"in_gallery"`
		Deletehash  string        `json:"deletehash"`
		Name        string        `json:"name"`
		Link        string        `json:"link"`
	} `json:"data"`
	Success bool `json:"success"`
	Status  int  `json:"status"`
}

type Size struct {
	Width  int
	Height int
}

type AppGeoSettings struct {
	MainWindowGeo struct {
		Pos  []int
		Size []int
	}
	ToolsListView      Size
	SessionHistoryList Size
	ImagePreviewFrame  Size
}

type WSHL struct {
	core.QObject
	_ func(string, string, string) `signal:"writeSessionHistoryList"`
}
