package gcors

type Option interface {
	apply(*options)
}

type options struct {
	origin  string
	methods []string
	headers []string
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithOrigin(origin string) Option {
	return optionFunc(func(opts *options) {
		opts.origin = origin
	})
}

func WithMethods(methods ...string) Option {
	return optionFunc(func(opts *options) {
		opts.methods = methods
	})
}

func WithHeaders(headers ...string) Option {
	return optionFunc(func(opts *options) {
		opts.headers = headers
	})
}
