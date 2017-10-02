package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

import "golang.org/x/crypto/bcrypt"

import (

    "net/http"
    "html/template"
    "log"

   
)


var db *sql.DB
var err error

var tpl *template.Template


func init(){
        tpl = template.Must(template.ParseGlob("templates/*gohtml"))
}

func main() {
        db, err = sql.Open("mysql", "root:blue42sethike@/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
        http.HandleFunc("/",index)
        http.HandleFunc("/about",about)
        http.HandleFunc("/terms",terms)
        http.HandleFunc("/privacy", privacy)
        http.HandleFunc("/contact",contact)
        http.HandleFunc("/faq",faq)
        http.HandleFunc("/account-archive",accountarchive)
        http.HandleFunc("/account-close",accountclose)
        http.HandleFunc("/account-favorites",accountfavorites)
        http.HandleFunc("/account-home",accounthome)
        http.HandleFunc("/account-posts",accountposts)
        http.HandleFunc("/account-saved-search",accountsavedsearch)
        http.HandleFunc("/account-security",accountsecurity)
        http.HandleFunc("/automobiles",automobiles)
        http.HandleFunc("/community",community)
        http.HandleFunc("/forgot-password",forgotpassword)
        http.HandleFunc("/forsale",forsale)
        http.HandleFunc("/home",home)
        http.HandleFunc("/jobs",jobs)
        http.HandleFunc("/learning",learning)
        http.HandleFunc("/login",login)
        http.HandleFunc("/pets",pets)
        http.HandleFunc("/post-automobiles",postautomobiles)
        http.HandleFunc("/post-community", postcommunity)
        http.HandleFunc("/post-for-sale",postforsale)
        http.HandleFunc("/post-jobs",postjobs)
        http.HandleFunc("/post-learning",postlearning)
        http.HandleFunc("/post-pets",postpets)
        http.HandleFunc("/post-property",postproperty)
        http.HandleFunc("/post-services",postservices)
        http.HandleFunc("/post",post)
        http.HandleFunc("/property",property)
        http.HandleFunc("/services",services)
        http.HandleFunc("/signup",signup)
        http.HandleFunc("/search",search)               
        http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
        http.ListenAndServe(":8080", nil)
}


func index(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func about(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func terms(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "terms.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func privacy(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "privacy.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func contact(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func faq(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "faq.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func accountarchive(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "account-archive.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func accountclose(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "account-close.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func accountfavorites(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "account-favorites.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func accounthome(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "account-home.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func accountposts(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "account-posts.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func accountsavedsearch(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "account-saved-search.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func accountsecurity(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "account-security.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func automobiles(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "automobiles.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func community(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "community.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func forgotpassword(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "forgot-password.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func forsale(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "forsale.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func home(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "home.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func jobs(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "jobs.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func learning(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "learning.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}

func pets(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "pets.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func postautomobiles(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "post-automobiles.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func postcommunity(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "post-community.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func postforsale(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "post-for-sale.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func postjobs(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "post-jobs.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func postlearning(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "post-learning.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func postpets(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "post-pets.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func postproperty(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "post-property.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func postservices(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "post-services.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func post(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "post.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func property(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "property.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func services(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "services.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
func login(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "login.gohtml")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	var databaseUsername string
	var databasePassword string

	err := db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&databaseUsername, &databasePassword)

	if err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil {
		http.Redirect(res, req, "/post", 301)
		return
	}

	http.Redirect(res, req, "/post", 301)
      
}
func signup(res http.ResponseWriter, req *http.Request) {
    if req.Method != "POST" {
        http.ServeFile(res, req, "signup.gohtml")
        return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")
        confirm := req.FormValue("confirm")

	var user string

    err := db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)

    switch {
    case err == sql.ErrNoRows:
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        if err != nil {
            http.Error(res, "Server error, unable to create your account.", 500)    
            return
        
        }  
   if confirm != password {
         http.Error(res, "passwords do not match, unable to create your account.", 500)    
            return
        http.Redirect(res, req, "/signup", 301)
   }
        _, err = db.Exec("INSERT INTO users(username, password) VALUES(?, ?)", username, hashedPassword)
        if err != nil {
            http.Error(res, "Server error, unable to create your account.", 500)    
            return
        }

        res.Write([]byte("User created!"))
        return
    case err != nil: 
        http.Error(res, "Server error, unable to create your account.", 500)    
        return
    default: 
        http.Redirect(res, req, "/post", 301)
    }
}
func search(w http.ResponseWriter, r *http.Request) {
        err := tpl.ExecuteTemplate(w, "search.gohtml", nil)
        if err != nil {
                log.Println(err)
        }
}
