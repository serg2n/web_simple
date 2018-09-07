package http

import (
	"net/http"
	"simple-web-app/constants"
)

type Router struct {
	cc *ContactController
}

func (r *Router) configureRouting(resolver *RegexpResolver) {

	resolver.AddFileHandler("GET /$", http.Dir(constants.ASSETS_PATH))
	resolver.AddFileHandler("GET /js", http.Dir(constants.ASSETS_PATH))

	//resolver.Add("GET /$", IndexHandler)

	resolver.Add("GET /contact/?$", r.cc.Contacts)
	resolver.Add("POST /contact/?$", r.cc.CreateContact)
	resolver.Add("GET /contact/([0-9]+)/?$", r.cc.Contact)
	resolver.Add("PUT /contact/([0-9]+)/?$", r.cc.UpdateContact)
	resolver.Add("DELETE /contact/([0-9]+)/?$", r.cc.DeleteContact)
}

func NewRouter(contactController *ContactController) *Router {
	return &Router{
		cc: contactController,
	}
}
