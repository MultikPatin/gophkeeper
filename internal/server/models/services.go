package models

type User struct {
	ID       int64
	Login    string
	Password string
}
type Password struct {
	ID       int64
	Title    string
	UserID   int64
	Login    []byte
	Password []byte
}

type Card struct {
	ID         int64
	Title      string
	UserID     int64
	Bank       []byte
	Number     []byte
	DataEnd    []byte
	SecretCode []byte
}

type BinaryData struct {
	ID     int64
	Title  string
	UserID int64
	Data   []byte
}
