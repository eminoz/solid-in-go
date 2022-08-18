package main

import "fmt"

type user interface {
	getUser(u string) string
	getAllUser(str string) [2]string
}

type manager interface {
	user
	deleteUser(s string) string
}

type users struct {
}

func (u users) getUser(usr string) string {

	return usr
}
func (u users) getAllUser(str string) [2]string {
	strings := [...]string{"ali", "emin"}
	return strings
}

type allUsers struct {
}

func (a allUsers) getUser(u string) string {
	return u
}
func (a allUsers) deleteUser(s string) string {
	return "user deleted"
}
func (a allUsers) getAllUser(str string) [2]string {
	strings := [...]string{str, "emin"}
	return strings
}

func main() {
	u := users{}
	fmt.Println(u.getUser("burak"), u.getAllUser("emin"))
	a := allUsers{}
	fmt.Println(a.deleteUser("emin"), a.getUser("ali"), a.getAllUser("emin"))
}
