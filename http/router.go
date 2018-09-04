package http

type Router struct {
	cc *ContactController
}

func (r *Router) configureRouting(resolver *RegexpResolver) {
	resolver.Add("GET /contact/?$", r.cc.Contacts)
	//resolver.Add("POST /user/?$", CreateContact)
	resolver.Add("GET /contact/([0-9]+)/?$", r.cc.Contact)
	//resolver.Add("PUT /user/([0-9]+)/?$", UpdateContact)
	//resolver.Add("DELETE /user/([0-9]+)/?$", DeleteContact)
}

func NewRouter(contactController *ContactController) *Router {
	return &Router{
		cc: contactController,
	}
}
