package models

type User struct {
	ID       int64
	Login    string
	Password string
}

type Card struct {
	ID         int64
	Title      string
	Bank       string
	Owner      string
	Number     string
	DataEnd    string
	SecretCode string
}

type Password struct {
	ID       int64
	Title    string
	Owner    string
	Login    string
	Password string
}

type BinaryData struct {
	ID    int64
	Title string
	Owner int64
	Data  []byte
}
