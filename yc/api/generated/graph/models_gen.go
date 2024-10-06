// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

type Link struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Address string `json:"address"`
	User    *User  `json:"user"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Mutation struct {
}

type NewLink struct {
	Title   string `json:"title"`
	Address string `json:"address"`
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Query struct {
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

// `Subscription` is where all the subscriptions your clients can
// request. You can use Schema Directives like normal to restrict
// access.
type Subscription struct {
}

type Time struct {
	UnixTime  int    `json:"unixTime"`
	TimeStamp string `json:"timeStamp"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
