package router

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/kaung-minkhant/go-restaurent/utils"
)

type cookieData struct {
	RefToken string `json:"refresh_token"`
	Role     string `json:"role"`
}

const cookieKey = "xx-cookie"

func setDataCookie(w http.ResponseWriter, data *cookieData) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		SomethingWentWrongResponse(w)
		return
	}
	formatedData := strings.ReplaceAll(string(jsonData), "\"", "'")

	signature, err := utils.SignData([]byte(formatedData), nil)
	if err != nil {
		SomethingWentWrongResponse(w)
		return
	}
	cookieDataWithSignature := string(formatedData) + "::" + signature
	http.SetCookie(w, &http.Cookie{
		Name:     cookieKey,
		Value:    cookieDataWithSignature,
		HttpOnly: true,
		Path:     "/",
	})
}
func getDataFromCookie(r *http.Request) (*cookieData, error) {
	dataFromCookie, err := r.Cookie(cookieKey)
	if err != nil {
		return nil, utils.ReturnAccessDenied()
	}
	cookieChunks := strings.Split(dataFromCookie.Value, "::")
	if len(cookieChunks) != 2 {
		return nil, utils.ReturnAccessDenied()
	}
	err = utils.VerifySign([]byte(cookieChunks[0]), cookieChunks[1], nil)
	if err != nil {
		return nil, utils.ReturnAccessDenied()
	}
	decodedData := strings.ReplaceAll(cookieChunks[0], "'", "\"")
	data := &cookieData{}
	if err := json.Unmarshal([]byte(decodedData), data); err != nil {
		return nil, utils.ReturnAccessDenied()
	}

	return data, nil
}

func getRefreshTokenFromCookie(r *http.Request) (string, error) {
	data, err := getDataFromCookie(r)
	if err != nil {
		return "", err
	}
	return data.RefToken, nil
}
