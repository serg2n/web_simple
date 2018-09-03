package http

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type RegexpResolver struct {
	handlers map[string]http.HandlerFunc
	cache    map[string]*regexp.Regexp
}

func NewPathResolver() *RegexpResolver {
	return &RegexpResolver{
		handlers: make(map[string]http.HandlerFunc),
		cache:    make(map[string]*regexp.Regexp),
	}
}

func (r *RegexpResolver) Add(regex string, handler http.HandlerFunc) {
	r.handlers[regex] = handler
	cache, _ := regexp.Compile(regex)
	r.cache[regex] = cache
}

func (r *RegexpResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path
	for pattern, handlerFunc := range r.handlers {
		if r.cache[pattern].MatchString(check) == true {
			handlerFunc(res, req)
			return
		}
	}

	http.NotFound(res, req)
}

func GetIdFromRequest(req *http.Request, level int) (int32, error) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	id := ""

	if len(parts) > level {
		id = parts[level]
	}

	log.Printf("Extracted id from url: %s, %s", path, id)

	res, err := strconv.ParseInt(id, 10, 32)
	return int32(res), err
}

func InternalServerErrorResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error. Please check logs"))
}

func BadRequestResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bad Request. Please check logs"))
}
