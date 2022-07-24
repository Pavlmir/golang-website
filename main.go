package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type User struct {
	Name string
	Age uint16
	Money int16
	avg_grades, happiness float64
	Hobbies []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User: %s, %d, %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	// http://localhost:8080/
	bob := User{Name: "Bob", Age: 25, Money: -50, avg_grades: 4.2, happiness: 0.5, Hobbies: []string{"Football", "Skate", "Dance"}}
	bob.setNewName("Alex")
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
	fmt.Fprint(w, "go is super easy!" + bob.getAllInfo())
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Contacts page")
}

func handleRequest(){
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
