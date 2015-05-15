package session

// Session abstract session interface
type Session interface {
	Getter
	Setter
	Deleter
	Scavenger
}

// Getter get data from session
type Getter interface {
	Get(token string, key string) interface{}
}

// Setter save data in session
type Setter interface {
	Set(token string, key string, val interface{}) error
}

// Deleter delete data from session
type Deleter interface {
	Delete(token string, key string)
}

// Scavenger expires the token
type Scavenger interface {
	Expire(token string)
}
