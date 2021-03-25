package opt

import "time"

type options struct {
	timeout time.Duration
	caching bool
}

type Option interface {
	apply(*options)
}
type optionFunc func(*options)

func (f optionFunc)apply(o *options) {
	f(o)
}

func WithTimeout(t time.Duration) Option {
	return optionFunc(func(o *options){
		o.timeout = t
	})
}
func WithCaching(cache bool) Option {
	return optionFunc(func(o *options) {
		o.caching = cache
	})
}
var (
	defaultTimeout = time.Duration(10)*time.Second
	defaultCaching = true
)
func Connect(addr string, opts ...Option){
	optDefault := options{
		timeout:defaultTimeout,
		caching:defaultCaching,
	}
	for _,o:=range opts {
		o.apply(&optDefault)
	}
}