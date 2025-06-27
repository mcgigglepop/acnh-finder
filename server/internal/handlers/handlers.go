package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/mcgigglepop/acnh-finder/server/internal/config"
	"github.com/mcgigglepop/acnh-finder/server/internal/forms"
	"github.com/mcgigglepop/acnh-finder/server/internal/models"
	"github.com/mcgigglepop/acnh-finder/server/internal/render"
)

// repository used by the handlers
var Repo *Repository

// the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// ////////////////////////////////////////////////////////////
// ////////////////////////////////////////////////////////////
// /////////////////// GET REQUESTS ///////////////////////////
// ////////////////////////////////////////////////////////////
// ////////////////////////////////////////////////////////////

func (m *Repository) GetIndex(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "index.page.tmpl", &models.TemplateData{})
}

func (m *Repository) LoginGet(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "login.page.tmpl", &models.TemplateData{})
}

func (m *Repository) RegisterGet(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "register.page.tmpl", &models.TemplateData{})
}

func (m *Repository) EmailVerificationGet(w http.ResponseWriter, r *http.Request) {
	email := m.App.Session.GetString(r.Context(), "user_email")

	if email == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	render.Template(w, r, "email-verification.page.tmpl", &models.TemplateData{})
}

func (m *Repository) DashboardGet(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "dashboard.page.tmpl", &models.TemplateData{})
}

func (m *Repository) CollectionsGet(w http.ResponseWriter, r *http.Request) {
	// Get the user_id (Cognito sub) from session
	userSub := m.App.Session.GetString(r.Context(), "user_id")
	if userSub == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Fetch collections for the user from DynamoDB
	collections, err := m.App.Dynamo.Collections.GetCollectionsByUser(r.Context(), userSub)
	if err != nil {
		m.App.ErrorLog.Println("Error fetching collections:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Render the collections template
	render.Template(w, r, "collections.page.tmpl", &models.TemplateData{
		Data: map[string]interface{}{
			"Collections": collections,
		},
	})
}

//////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////
///////////////////// POST REQUESTS //////////////////////////
//////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////

func (m *Repository) CreateCollectionPost(w http.ResponseWriter, r *http.Request) {
	// Renew the session token (useful for rotation and session security)
	if err := m.App.Session.RenewToken(r.Context()); err != nil {
		m.App.ErrorLog.Println("Session token renewal failed:", err)
	}

	// Parse the incoming form data
	err := r.ParseForm()
	if err != nil {
		m.App.ErrorLog.Println("ParseForm error:", err)
	}

	// Wrap form data for validation
	form := forms.New(r.PostForm)
	form.Required("title")

	// If form validation fails, re-render the page
	if !form.Valid() {
		m.App.ErrorLog.Println("form validation failed")
		render.Template(w, r, "create-collection.page.tmpl", &models.TemplateData{
			Form: form,
		})
		return
	}

	// Clean input values
	collectionTitle := r.Form.Get("title")
	collectionDescription := r.Form.Get("description")

	// Retrieve user ID from session
	userID := m.App.Session.GetString(r.Context(), "user_id")
	if userID == "" {
		m.App.ErrorLog.Println("invalid user session")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// create collection
	err = m.App.Dynamo.Collections.CreateCollection(r.Context(), userID, collectionTitle, collectionDescription)
	if err != nil {
		m.App.ErrorLog.Println("error creating collection", err)
		m.App.Session.Put(r.Context(), "error", "can't create collection")
		http.Redirect(w, r, "/create-collection", http.StatusSeeOther)
		return
	}

	// Flash and redirect
	m.App.Session.Put(r.Context(), "flash", "Collection created successfully.")
	http.Redirect(w, r, "/collections", http.StatusSeeOther)
}

func (m *Repository) RegisterPost(w http.ResponseWriter, r *http.Request) {
	if err := m.App.Session.RenewToken(r.Context()); err != nil {
		m.App.ErrorLog.Println("Session token renewal failed:", err)
	}

	err := r.ParseForm()
	if err != nil {
		m.App.ErrorLog.Println("ParseForm error:", err)
	}

	form := forms.New(r.PostForm)

	form.Required("email", "password")

	form.IsEmail("email")

	if !form.Valid() {
		render.Template(w, r, "register.page.tmpl", &models.TemplateData{
			Form: form,
		})
		return
	}

	email := strings.TrimSpace(r.Form.Get("email"))
	password := r.Form.Get("password")

	userErr := m.App.CognitoClient.RegisterUser(r.Context(), email, password)
	if userErr != nil {
		m.App.ErrorLog.Println("Cognito RegisterUser failed:", userErr)
		m.App.Session.Put(r.Context(), "error", "Registration failed. Please try again.")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	m.App.Session.Put(r.Context(), "user_email", email)

	m.App.Session.Put(r.Context(), "flash", "Registered successfully.")
	http.Redirect(w, r, "/email-verification", http.StatusSeeOther)
}

func (m *Repository) EmailVerificationPost(w http.ResponseWriter, r *http.Request) {
	if err := m.App.Session.RenewToken(r.Context()); err != nil {
		m.App.ErrorLog.Println("Session token renewal failed:", err)
	}

	email := m.App.Session.GetString(r.Context(), "user_email")

	if email == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		m.App.ErrorLog.Println("ParseForm error:", err)
	}

	form := forms.New(r.PostForm)
	form.Required("otpFirst", "otpSecond", "otpThird", "otpFourth", "otpFifth", "otpSixth")

	if !form.Valid() {
		log.Printf("[DEBUG] Form validation failed: %+v", form.Errors)
		render.Template(w, r, "email-verification.page.tmpl", &models.TemplateData{
			Form: form,
		})
		return
	}

	otpCode := strings.TrimSpace(
		r.Form.Get("otpFirst") +
			r.Form.Get("otpSecond") +
			r.Form.Get("otpThird") +
			r.Form.Get("otpFourth") +
			r.Form.Get("otpFifth") +
			r.Form.Get("otpSixth"),
	)

	_, err := m.App.CognitoClient.ConfirmUser(r.Context(), email, otpCode)
	if err != nil {
		m.App.ErrorLog.Printf("Cognito ConfirmUser failed: %v", err)
		m.App.Session.Put(r.Context(), "error", "Email verification failed. Please try again.")
		http.Redirect(w, r, "/email-verification", http.StatusSeeOther)
		return
	}

	m.App.Session.Remove(r.Context(), "user_email")
	m.App.Session.Put(r.Context(), "flash", "Email Verified.")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (m *Repository) LoginPost(w http.ResponseWriter, r *http.Request) {
	if err := m.App.Session.RenewToken(r.Context()); err != nil {
		m.App.ErrorLog.Println("Session token renewal failed:", err)
	}

	err := r.ParseForm()
	if err != nil {
		m.App.ErrorLog.Println("ParseForm error:", err)
	}

	form := forms.New(r.PostForm)

	form.Required("email", "password")

	form.IsEmail("email")

	if !form.Valid() {
		render.Template(w, r, "login.page.tmpl", &models.TemplateData{
			Form: form,
		})
		return
	}

	email := strings.TrimSpace(r.Form.Get("email"))
	password := r.Form.Get("password")

	auth_response, userErr := m.App.CognitoClient.Login(r.Context(), email, password)
	if userErr != nil {
		m.App.ErrorLog.Println("Cognito Login failed:", userErr)
		m.App.Session.Put(r.Context(), "error", "Login failed. Please try again.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	sub, err := m.App.CognitoClient.ExtractSubFromToken(r.Context(), auth_response.IdToken)

	if err != nil {
		// handle error
	}

	m.App.Session.Put(r.Context(), "user_id", sub)
	m.App.Session.Put(r.Context(), "id_token", auth_response.IdToken)
	m.App.Session.Put(r.Context(), "access_token", auth_response.AccessToken)
	m.App.Session.Put(r.Context(), "refresh_token", auth_response.RefreshToken)

	m.App.Session.Put(r.Context(), "flash", "login successfully.")
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}