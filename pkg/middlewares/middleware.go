package middlewares

type Middlewares struct {
	Post *PostMiddleware
}

func NewMiddlewares(post *PostMiddleware) *Middlewares {
	return &Middlewares{Post: post}
}
