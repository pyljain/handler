package openai

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"time"
)

const (
	openAIURL = "https://api.openai.com"
)

type OpenAIProxy struct {
}

func (p *OpenAIProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	userRequestStart := time.Now()

	// ***TODO: Please make sure to check the request path to only support the ones allowlisted for the company.
	// Exclude DALL-E and Assistants API, fine tuning etc.****

	url, err := url.Parse(openAIURL)
	if err != nil {
		fmt.Printf("Unable to parse the target URL: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	r.Host = "api.openai.com"
	// rec := httptest.NewRecorder()
	clientReq, _ := httputil.DumpRequest(r, true)
	// Request can be inspected and handled for model, prompt checking etc. with the current Logger / other middleware we have.
	log.Printf("Client request is %s", clientReq)

	rec := httptest.NewRecorder()
	if r.Method == http.MethodConnect {
		w.WriteHeader(http.StatusOK)
		return
	}

	llmCallStartTime := time.Now()
	proxy.ServeHTTP(rec, r)
	llmProcessingTimeTaken := time.Since(llmCallStartTime)

	// Copying the response buffer for inspection and handling as needed
	var newBuffer bytes.Buffer
	newBuffer.Write(rec.Body.Bytes())

	var reader io.Reader
	switch rec.Header().Get("Content-Encoding") {
	// Handling gzip and other formats to accommoate streaming and non-streaming responses
	case "gzip":
		reader, err = gzip.NewReader(&newBuffer)
		if err != nil {
			log.Printf("Error in zip %s", err)
		}
	default:
		reader = &newBuffer
	}

	respData, _ := io.ReadAll(reader)

	log.Printf("Response from OpenAI is %s", respData)
	for k, v := range rec.Header() {
		w.Header().Add(k, v[0])
	}
	userRequestTotalTime := time.Since(userRequestStart)
	w.Header().Add("X-R2D2-TOTAL-TIME-TAKEN", fmt.Sprintf("%f", userRequestTotalTime.Seconds()))
	w.Header().Add("X-R2D2-LLM-PROCESSING-TIME", fmt.Sprintf("%f", llmProcessingTimeTaken.Seconds()))
	w.Header().Add("X-R2D2-PROXY-PROCESSING-TIME", fmt.Sprintf("%f", userRequestTotalTime.Seconds()-llmProcessingTimeTaken.Seconds()))

	w.WriteHeader(rec.Code)

	rec.Body.WriteTo(w)

	// log.Printf("Body is %s", rec.Body)
}

/*
1. Log request, model being used - supported
2. Log response - supported
3. Try multiple APIs such as completion and embeddings - supported
4. Try streaming - supported
5. Try gRPC with a another reverse proxy
*/
