package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	defaultPort := os.Getenv("PORT")
	if defaultPort == "" {
		defaultPort = "8282"
	}

	port := flag.String("p", defaultPort, "The port to listen on")
	flag.Parse()

	addr := "0.0.0.0:" + *port
	server := http.Server{
		Addr:    addr,
		Handler: &echoServer{},
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

type echoServer struct{}

func (*echoServer) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	responseDoc := make(map[string]interface{})

	responseDoc["method"] = request.Method
	responseDoc["URL"] = extractUrlInfo(request.URL)
	responseDoc["headers"] = request.Header

	bodyAsBytes, err := ioutil.ReadAll(request.Body)
	if maybeReplyError(err, responseWriter) {
		return
	}
	responseDoc["body"] = string(bodyAsBytes)

	responseAsBytes, err := json.Marshal(responseDoc)
	if maybeReplyError(err, responseWriter) {
		return
	}

	responseWriter.Write(responseAsBytes)
}

func extractUrlInfo(url *url.URL) map[string]interface{} {
	urlDoc := make(map[string]interface{})

	urlDoc["scheme"] = url.Scheme
	urlDoc["host"] = url.Host
	urlDoc["path"] = url.Path
	urlDoc["rawPath"] = url.RawPath
	urlDoc["query"] = url.RawQuery
	urlDoc["fragment"] = url.Fragment

	if user := url.User; user == nil {
		urlDoc["userInfo"] = nil
	} else {
		userInfo := make(map[string]interface{})
		userInfo["user"] = user.Username()
		passwordValue, passwordPresent := user.Password()
		userInfo["passwordPresent"] = passwordPresent
		userInfo["passwordValue"] = passwordValue

		urlDoc["userInfo"] = userInfo
	}

	return urlDoc
}

func maybeReplyError(err error, responseWriter http.ResponseWriter) bool {
	if err == nil {
		return false
	} else {
		http.Error(responseWriter, err.Error(), 500)
		return true
	}
}
