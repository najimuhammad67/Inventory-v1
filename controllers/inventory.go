package controllers

import (
	"html/template"
	"inventory-app/middleware"
	"inventory-app/models"
	"inventory-app/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AdminDashboard(w http.ResponseWriter, r *http.Request) {
	session, _ := middleware.Store.Get(r, "session")
	if session.Values["authenticated"] != true || session.Values["role"] != "admin" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	items := models.GetAllItems()
	tmpl := template.Must(template.ParseFiles("views/admin.html"))
	tmpl.Execute(w, map[string]interface{}{
		"Items": items,
	})
}

func InventoryPage(w http.ResponseWriter, r *http.Request) {
	session, _ := middleware.Store.Get(r, "session")
	if session.Values["authenticated"] != true {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	items := models.GetAllItems()
	tmpl := template.Must(template.ParseFiles("views/inventory.html"))
	tmpl.Execute(w, map[string]interface{}{
		"Items": items,
		"Role":  session.Values["role"],
	})
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	session, _ := middleware.Store.Get(r, "session")
	if session.Values["authenticated"] != true || session.Values["role"] != "admin" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	r.ParseForm()
	name := r.Form.Get("name")
	category := r.Form.Get("category")
	quantity, _ := strconv.Atoi(r.Form.Get("quantity"))

	models.InsertItem(name, category, quantity)
	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	session, _ := middleware.Store.Get(r, "session")
	if session.Values["authenticated"] != true || session.Values["role"] != "admin" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	name := r.Form.Get("name")
	category := r.Form.Get("category")
	quantity, _ := strconv.Atoi(r.Form.Get("quantity"))

	models.UpdateItem(id, name, category, quantity)
	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	session, _ := middleware.Store.Get(r, "session")
	if session.Values["authenticated"] != true || session.Values["role"] != "admin" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	models.DeleteItem(id)
	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}

func ExportPDF(w http.ResponseWriter, r *http.Request) {
	session, _ := middleware.Store.Get(r, "session")
	if session.Values["authenticated"] != true {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	items := models.GetAllItems() // ✅ ambil semua data
	utils.ExportPDF(w, items)     // ✅ kirim ke fungsi export
}
