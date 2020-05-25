package user

type User struct {
	FirstName   string
	LastName    string
}
func New(firstName string, lastName string, totalLeave int, leavesTaken int) User {
	user := User {firstName, lastName}
	return user
}