package http

type Router struct {
	cc *ContactController
	ic *IndexController
}

func (r *Router) configureRouting(resolver *RegexpResolver) {

	resolver.Add("GET /", r.ic.indexHandler)

	resolver.Add("GET /contact/?$", r.cc.Contacts)
	resolver.Add("POST /contact/?$", r.cc.CreateContact)
	resolver.Add("GET /contact/([0-9]+)/?$", r.cc.Contact)
	resolver.Add("PUT /contact/([0-9]+)/?$", r.cc.UpdateContact)
	resolver.Add("DELETE /contact/([0-9]+)/?$", r.cc.DeleteContact)
}

func NewRouter(contactController *ContactController, indexController *IndexController) *Router {
	return &Router{
		cc: contactController,
	}
}
