package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/insanXYZ/golang-mux/model"
	"net/http"
	"strings"
)

type ContactsController struct {
	DB *sql.DB
}

func NewContactsController(DB *sql.DB) *ContactsController {
	return &ContactsController{DB: DB}
}

func (controller *ContactsController) GetAllContacts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		rows, err := controller.DB.QueryContext(r.Context(), "select * from contacts")
		if err != nil {
			fmt.Fprint(w, err.Error())
		}

		defer rows.Close()

		var contacts []model.Contacts

		for rows.Next() {
			var contact model.Contacts
			err := rows.Scan(&contact.ID, &contact.Name, &contact.Email, &contact.CreatedAt, &contact.UpdatedAt)
			if err != nil {
				return
			}
			contacts = append(contacts, contact)
		}

		encoder := json.NewEncoder(w)
		err = encoder.Encode(contacts)
		if err != nil {
			fmt.Fprint(w, err.Error())
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "bad request")
	}
}

func (controller *ContactsController) Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		var req model.InsertContacts
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			fmt.Fprint(w, err.Error())
		}

		_, err = controller.DB.ExecContext(r.Context(), "insert into contacts(name,email) values(?,?)", req.Name, req.Email)
		if err != nil {
			return
		}

		fmt.Fprintf(w, "success insert to database")

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "bad request")
	}
}

func (controller *ContactsController) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		id := strings.Split(r.URL.Path, "/")[2]
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "parameter id required")
		}

		_, err := controller.DB.ExecContext(r.Context(), "delete from contacts where id = ?", id)
		if err != nil {
			fmt.Fprint(w, err.Error())
		}

		fmt.Fprint(w, "success delete contact")

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "bad request")
	}
}

func (controller *ContactsController) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {

		id := strings.Split(r.URL.Path, "/")[2]
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "parameter id required")
		}

		var req model.UpdateContacts
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err.Error())
		}

		_, err = controller.DB.ExecContext(r.Context(), "update contacts set name = coalesce(?, name), email = coalesce(?, email) where id = ?", req.Name, req.Email, id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err.Error())
		}
		fmt.Fprint(w, "success update contact")

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "bad request")
	}
}
