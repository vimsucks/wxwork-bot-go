package wxworkbot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"
)

type uploadedMediaResponse struct {
	wxWorkResponse
	UploadedMedia
}

type UploadedMedia struct {
	Type      string `json:"type"'`
	MediaID   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

func uploadApiUrl(key *string) string {
	return fmt.Sprintf(
		"https://qyapi.weixin.qq.com/cgi-bin/webhook/upload_media?key=%s&type=file",
		*key,
	)
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func (bot *WxWorkBot) UploadMedia(fileName string, fileBytes *[]byte) (*UploadedMedia, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="media"; filename="%s"; filelength=%d`,
			escapeQuotes(fileName), len(*fileBytes)))
	h.Set("Content-Type", "application/octet-stream")

	part, err := writer.CreatePart(h)
	if err != nil {
		return nil, err
	}
	io.Copy(part, bytes.NewReader(*fileBytes))
	writer.Close()

	req, err := http.NewRequest(http.MethodPost, uploadApiUrl(&bot.Key), body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	if err != nil {
		return nil, err
	}

	resp, err := bot.Client.Do(req)
	if err != nil {
		return nil, err
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var wxWorkResp uploadedMediaResponse
	err = json.Unmarshal(respBody, &wxWorkResp)
	if err != nil {
		return nil, err
	}
	if wxWorkResp.ErrorCode != 0 && wxWorkResp.ErrorMessage != "" {
		return nil, errors.New(string(respBody))
	}
	return &wxWorkResp.UploadedMedia, nil
}
