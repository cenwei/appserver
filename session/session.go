package session

// Session abstract session interface
type Session interface {
	Getter
	Setter
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

// Scavenger delete data from session
type Scavenger interface {
	Delete(token string, key string)
}
