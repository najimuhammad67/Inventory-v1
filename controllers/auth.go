package controllers

import (
	"html/template"
	"inventory-app/middleware"
	"inventory-app/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// Tampilkan halaman login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("views/login.html"))
		tmpl.Execute(w, nil)
		return
	}

	// Proses login
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	user, err := models.GetUserByUsername(username)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	session, _ := middleware.Store.Get(r, "session")
	session.Values["authenticated"] = true
	session.Values["username"] = user.Username
	session.Values["role"] = user.Role
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Tampilkan halaman register atau proses registrasi
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("views/register.html"))
		tmpl.Execute(w, nil)
		return
	}

	// Proses register
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// Buat akun baru dengan role default: user
	models.CreateUser(username, string(hashed), "user")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Logout dan hapus session
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := middleware.Store.Get(r, "session")
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
