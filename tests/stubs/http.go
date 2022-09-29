package stubs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/dougmendes/gondalf/pkg/infra/controller"
	"golang.org/x/exp/slices"
)

type RouterStub struct {
	controller.Router
	routes []Route
}

type Route struct {
	pattern string
	handler func(w http.ResponseWriter, r *http.Request) error
}

func (r RouterStub) Exec(rawURL string, w http.ResponseWriter, req *http.Request) {
	route := r.getRouterFromPattern(rawURL)
	params := route.getURIParams(rawURL)
	newReq := req.WithContext(WithParams(req.Context(), params))
	route.handler(w, newReq)
}

func (r RouterStub) getRouterFromPattern(url string) Route {
	routeIndex := slices.IndexFunc(r.routes, func(route Route) bool {
		re := regexp.MustCompile("{(.?+)}")
		newPattern := re.ReplaceAllString(route.pattern, ".*")
		re = regexp.MustCompile(newPattern)
		return re.MatchString(url)
	})
	if routeIndex < 0 {
		log.Fatal(context.Background(), fmt.Sprintf("route GET %s not found", url))
	}
	route := r.routes[routeIndex]

	return route
}
func (r Route) getURIParams(rawURL string) URIParams {
	result := make(URIParams)

	re := regexp.MustCompile("{(.+?)}")
	for _, val := range re.FindAllString(r.pattern, -1) {
		name := re.ReplaceAllString(val, "$1")
		value := r.matchPatternSectionWithURLValue(val, rawURL)
		result[name] = value
	}
	return result
}

func (r Route) matchPatternSectionWithURLValue(expectedSection, url string) string {
	for i, section := range strings.Split(r.pattern, "/") {
		if section == expectedSection {
			return strings.Split(url, "/")[i]
		}
	}
	return ""
}

type ResponseWriterStub struct {
	StatusCode int
	Data       []byte
}

func (w ResponseWriterStub) Header() http.Header {
	return make(http.Header)
}

func (w *ResponseWriterStub) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
}

func (w *ResponseWriterStub) Write(data []byte) (int, error) {
	w.Data = data
	return len(data), nil
}

// web
func WithParams(ctx context.Context, params URIParams) context.Context {
	return context.WithValue(ctx, uriParamsContextKey{}, params)
}

type uriParamsContextKey struct{}
type URIParams map[string]string

func Params(r *http.Request) URIParams {
	if v, ok := r.Context().Value(uriParamsContextKey{}).(URIParams); ok {
		return v
	}

	return nil
}

func (p URIParams) String(param string) (string, error) {
	v, ok := p[param]
	if !ok {
		return "", errors.New(fmt.Sprintf("%s uri param is not found: %s", http.StatusInternalServerError, param))
	}
	return v, nil
}
func EncodedJSON(w http.ResponseWriter, v interface{}, code int) error {
	if headerer, ok := v.(http.Header); ok {
		for k, values := range headerer {
			for _, v := range values {
				w.Header().Add(k, v)
			}
		}
	}

	var jsonData []byte
	var err error
	switch v := v.(type) {
	case []byte:
		jsonData = v
	case io.Reader:
		jsonData, err = ioutil.ReadAll(v)
	default:
		jsonData, err = json.Marshal(v)
	}

	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(code)
	if _, err := w.Write(jsonData); err != nil {
		return err
	}
	return nil
}
