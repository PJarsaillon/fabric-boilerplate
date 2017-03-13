package entities

type ECertResponse struct {
	OK string `json:"OK"`
}

type TestData struct {
	Users  []User  `json:"users"`
	Things []Thing `json:"things"`
}

type TestDataElement interface {
	ID() string
}

type User struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Hash     string `json:"hash"`
}

type Thing struct {
	ThingID      string `json:"thingID"`
	SomeProperty string `json:"someProperty"`
	UserID       string `json:"userID"`
	Toto         string `json:"toto"`
	Tata         string `json:"tata"`
}

type UserAuthenticationResult struct {
	User          User
	Authenticated bool
}

type Users struct {
	Users []User `json:"users"`
}

func (t *User) ID() string {
	return t.Username
}

func (t *Thing) ID() string {
	return t.ThingID
}
