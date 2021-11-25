package main

import (
	"bufio"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const htpasswdFilePath = "/var/www/vichan/vichan-users"
const defaultPassword = "burbfren"

var admins map[string]bool

func main() {
	admins = make(map[string]bool)
	admins["joshb"] = true
	admins["burdmin"] = true
	admins["katst"] = true
	admins["thomass"] = true
	admins["shannonj"] = true
	admins["mattyh"] = true

	setupHTTP()
	http.ListenAndServe(":8315", nil)
}

func setupHTTP() {
	http.HandleFunc("/createuser", func(w http.ResponseWriter, r *http.Request) {
		if !admins[getAuthUser()] {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		switch r.Method {
		case "GET":
			fmt.Fprintf(w, `
			<h1>Create new user with default password</h1>
			<form action="/createuser" method="post">
				<fieldset>
					<label>Username</label>
					<input name="user" />
				</fieldset>
				<fieldset>
					<label>Post Name</label>
					<input name="name" />
				</fieldset>
				<input type="submit" />
			</form>
			<a href="/changepassword">Change Password</a>
			`)
		case "POST":
			var formUsername, formName string
			r.ParseForm()

			fmt.Fprintf(w, "<p>")
			if formUsername = r.PostForm.Get("user"); formUsername == "" {
				fmt.Fprintf(w, "No Username")
			} else if formName = r.PostForm.Get("name"); formName == "" {
				fmt.Fprintf(w, "No post name")
			} else if err := writeUser(formUsername, formName, defaultPassword, true); err != nil {
				fmt.Fprintf(w, "Failed to save new user: %s", err.Error())
			} else {
				fmt.Fprintf(w, "Successfully created new user")
			}
			fmt.Fprintf(w, "</p>")

			fmt.Fprintf(w, `<a href="/createuser">Back</a>`)
		}
	})

	http.HandleFunc("/changepassword", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, `
			<h1>Change my password</h1>
			<form action="/changepassword" method="post">
				<fieldset>
					<label>New Password</label>
					<input type="password" name="new-password" />
				</fieldset>
				<fieldset>
					<label>Verify New Password</label>
					<input type="password" name="new-password-verify" />
				</fieldset>
				<input type="submit" />
			</form>
			<a href="/createuser">Create User (admins only)</a>
			`)
		case "POST":
			var passNew string
			r.ParseForm()

			fmt.Fprintf(w, "<p>")
			if passNew = r.PostForm.Get("new-password"); passNew == "" {
				fmt.Fprintf(w, "No password entered")
			} else if r.PostForm.Get("new-password-verify") != passNew {
				fmt.Fprintf(w, "Passwords don't match")
			} else if err := writeUser(getAuthUser(), "", passNew, false); err != nil {
				fmt.Fprintf(w, "Failed to save user: %s", err.Error())
			} else {
				fmt.Fprintf(w, "Successfully updated password")
			}
			fmt.Fprintf(w, "</p>")

			fmt.Fprintf(w, `<a href="/changepassword">Back</a>`)
		}
	})

}

func writeUser(username, name, passClear string, isCreate bool) error {
	updated := false
	file, err := os.OpenFile(htpasswdFilePath, os.O_CREATE|os.O_RDWR, 0664)
	if err != nil {
		return err
	}
	defer file.Close()

	var newFile strings.Builder

	s := bufio.NewScanner(file)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if strings.HasPrefix(l, username+":") {
			if isCreate {
				return errors.New("User already exists")
			}
			passwd, err := exec.Command("openssl", "passwd", "-6", passClear).Output()
			if err != nil {
				return err
			}
			newFile.WriteString(username)
			newFile.WriteRune(':')
			newFile.WriteString(strings.TrimSpace(string(passwd)))
			newFile.WriteRune('\n')
			updated = true
		} else {
			newFile.WriteString(l)
		}
	}

	if isCreate {
		passwd, err := exec.Command("openssl", "passwd", "-6", passClear).Output()
		if err != nil {
			return err
		}
		newFile.WriteString(username)
		newFile.WriteRune(':')
		newFile.WriteString(strings.TrimSpace(string(passwd)))
		newFile.WriteString("\n# ")
		newFile.WriteString(name)
		newFile.WriteString("\n\n")
		updated = true
	}

	if _, err = file.Seek(0, os.SEEK_SET); err != nil {
		return err
	}

	if err = file.Truncate(0); err != nil {
		return err
	}

	if _, err = file.WriteString(newFile.String()); err != nil {
		return err
	}

	if !updated {
		return errors.New("User does not exist")
	}

	return nil
}

func getAuthUser() string {
	return "joshb"
}
