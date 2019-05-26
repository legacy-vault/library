package httphelper

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/legacy-vault/library/go/header"
	"github.com/legacy-vault/library/go/mime"
)

// Functions which help in replying to HTTP Requests.

// ReplyTextWithCode Function replies to the HTTP Request with the specified
// Text and HTTP Status Code.
func ReplyTextWithCode(
	w http.ResponseWriter,
	httpStatusCode int,
	replyText string,
) {
	var xerr error

	w.WriteHeader(httpStatusCode)
	_, xerr = w.Write([]byte(replyText))
	if xerr != nil {
		log.Println(replyText)
		log.Println(xerr)
	}
}

// ReplyErrorWithCode Function replies to the HTTP Request with an Error and
// the specified HTTP Status Code.
func ReplyErrorWithCode(
	w http.ResponseWriter,
	httpStatusCode int,
	err error,
) {

	ReplyTextWithCode(w, httpStatusCode, err.Error())
}

// ReplyErrorInternal Function replies to the HTTP Request with an Error and
// 'Internal Server Error' HTTP Status Code.
func ReplyErrorInternal(
	w http.ResponseWriter,
	err error,
) {

	ReplyErrorWithCode(w, http.StatusInternalServerError, err)
}

// ReplyJSON Function sends an Object in JSON Format to the HTTP Output Stream.
func ReplyJSON(
	w http.ResponseWriter,
	replyObject interface{},
) {

	var err error
	var response []byte

	// Encode an Object with JSON Format.
	response, err = json.Marshal(replyObject)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

	// Send the Reply.
	w.Header().Set(header.HttpHeaderContentType, mime.TypeApplicationJson)
	_, err = w.Write(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}
