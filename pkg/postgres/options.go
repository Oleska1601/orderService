package postgres

import "time"

type Option func(*Postgres)

func SetMaxPoolSize(newMaxPoolSize int) Option {
	return func(p *Postgres) {
		p.maxPoolSize = newMaxPoolSize
	}
}

func SetMaxConnAttempts(newMaxConnAttempts int) Option {
	return func(p *Postgres) {
		p.maxConnAttempts = newMaxConnAttempts
	}
}

func SetMaxConnTimeout(newMaxConnTimeout time.Duration) Option {
	return func(p *Postgres) {
		p.maxConnTimeout = newMaxConnTimeout
	}
}
