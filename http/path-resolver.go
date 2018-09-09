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
	fileHandler map[string]http.Handler

	cache    map[string]*regexp.Regexp
	fileHandlerCache map[string]*regexp.Regexp
}

func NewPathResolver() *RegexpResolver {
	return &RegexpResolver{
		handlers: make(map[string]http.HandlerFunc),
		fileHandler: make(map[string]http.Handler),
		cache:    make(map[string]*regexp.Regexp),
		fileHandlerCache: make(map[string]*regexp.Regexp),
	}
}

func (r *RegexpResolver) Add(regex string, handler http.HandlerFunc) {
	r.handlers[regex] = handler
	cache, _ := regexp.Compile(regex)
	r.cache[regex] = cache
}

func (r *RegexpResolver) AddFileHandler(regex string, root http.FileSystem) {
	r.fileHandler[regex] = http.FileServer(root)
	cache, _ := regexp.Compile(regex)
	r.fileHandlerCache[regex] = cache

}

func (r *RegexpResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path

	for pattern, handler := range r.fileHandler {
		if r.fileHandlerCache[pattern].MatchString(check) == true {
			handler.ServeHTTP(res, req)
			return
		}
	}

	for pattern, handlerFunc := range r.handlers {
		if r.cache[pattern].MatchString(check) == true {
			handlerFunc(res, req)
			return
		}
	}

	log.Printf("%s handler not found", check)
	http.NotFound(res, req)
}

func IdFromRequest(req *http.Request, level int) (int, error) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	id := ""

	if len(parts) > level {
		id = parts[level]
	}

	log.Printf("Extracted id from url: %s, %s", path, id)

	res, err := strconv.ParseInt(id, 10, 32)
	return int(res), err
}

func InternalServerErrorResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error. Please check logs"))
}

func BadRequestResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bad Request. Please check logs"))
}
