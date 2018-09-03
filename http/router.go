package http

import "simple-web-app/http/controller"

func configureRouting(resolver *RegexpResolver) {
	resolver.Add("GET /contact/?$", controller.Contact)
}
