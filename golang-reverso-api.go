package reverso

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type translation_api_return struct {
	ContextResults *struct {
		Results []struct {
			Translation     string `json:"translation"`
			Transliteration string `json:"transliteration"`
		} `json:"results"`
	} `json:"contextResults,omitempty"`
	Extra_translation *[]string `json:"translation,omitempty"`
}

type translation_api_request struct {
	Format  string `json:"format"`
	From    string `json:"from"`
	Input   string `json:"input"`
	Options struct {
		SentenceSplitter  bool   `json:"sentenceSplitter"`
		Origin            string `json:"origin"`
		LanguageDetection bool   `json:"languageDetection"`
		ContextResults    bool   `json:"contextResults"`
	} `json:"options"`
	To string `json:"to"`
}

func Get_Translation(input_lang string, output_lang string, query string) translation_api_return {

	url := "https://api.reverso.net/translate/v1/translation"
	request_obj := &translation_api_request{
		Format: "text",
		From:   get_abbr(input_lang),
		Input:  query,
		Options: struct {
			SentenceSplitter  bool   `json:"sentenceSplitter"`
			Origin            string `json:"origin"`
			LanguageDetection bool   `json:"languageDetection"`
			ContextResults    bool   `json:"contextResults"`
		}{
			SentenceSplitter:  true,
			Origin:            "translation.web",
			ContextResults:    true,
			LanguageDetection: true,
		},
		To: get_abbr(output_lang),
	}
	request_obj_json, _ := json.Marshal(request_obj)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(request_obj_json))

	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "en-US,en;q=0.9")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("origin", "https://www.reverso.net")
	req.Header.Add("priority", "u=1, i")
	req.Header.Add("referer", "https://www.reverso.net/")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"134\", \"Not:A-Brand\";v=\"24\", \"Google Chrome\";v=\"134\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36")
	req.Header.Add("x-reverso-origin", "translation.web")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	fmt.Println(res.StatusCode)
	tr := translation_api_return{}
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
	json.Unmarshal(body, &tr)
	fmt.Println(tr)

	return tr

}

// converting languages to the abbreviation used in url
func get_abbr(prompt string) string {
	abbr := ""
	prompt = strings.ToLower(prompt)
	switch prompt {
	case "arabic":
		abbr = "ara"
	case "chinese":
		abbr = "chi"
	case "czech":
		abbr = "cze"
	case "danish":
		abbr = "dan"
	case "dutch":
		abbr = "dut"
	case "english":
		abbr = "eng"
	case "french":
		abbr = "fra"
	case "german":
		abbr = "ger"
	case "greek":
		abbr = "gre"
	case "hebrew":
		abbr = "heb"
	case "hindi":
		abbr = "hin"
	case "hungarian":
		abbr = "hun"
	case "italian":
		abbr = "ita"
	case "japanese":
		abbr = "jpn"
	case "korean":
		abbr = "kor"
	case "persian":
		abbr = "per"
	case "polish":
		abbr = "pol"
	case "portuguese":
		abbr = "por"
	case "romanian":
		abbr = "rom"
	case "russian":
		abbr = "rus"
	case "slovak":
		abbr = "slo"
	case "spanish":
		abbr = "spa"
	case "swedish":
		abbr = "swe"
	case "thai":
		abbr = "tha"
	case "turksih":
		abbr = "tur"
	case "ukranian":
		abbr = "ukr"

	}
	return abbr
}
