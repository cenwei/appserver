package sessionStore

// SessionStore abstract session store interface
type SessionStore interface {
	Getter
	Setter
	Deleter
	Scavenger
}

// Getter get data from session store
type Getter interface {
	Get(token string, key string) interface{}
}

// Setter save data in session store
type Setter interface {
	Set(token string, key string, val interface{}) error
}

// Deleter delete data from session store
type Deleter interface {
	Delete(token string, key string)
}

// Scavenger expires the token
type Scavenger interface {
	Expire(token string)
}
