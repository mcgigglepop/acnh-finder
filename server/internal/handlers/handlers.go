package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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
	// Get the user_id (Cognito sub) from session
	userSub := m.App.Session.GetString(r.Context(), "user_id")
	if userSub == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	user, err := m.App.Dynamo.UserProfile.GetUserProfile(r.Context(), userSub)
	if err != nil {
		log.Printf("Couldn't fetch user: %v", err)
		// redirect to error page or some default page
	} else {
		if user.Hemisphere == "unset" {
			http.Redirect(w, r, "/choose-hemisphere", http.StatusSeeOther)
		}
	}

	m.App.Session.Put(r.Context(), "user_hemisphere", user.Hemisphere)

	render.Template(w, r, "dashboard.page.tmpl", &models.TemplateData{})
}

func (m *Repository) ChooseHemisphereGet(w http.ResponseWriter, r *http.Request) {
	// Get the user_id (Cognito sub) from session
	userSub := m.App.Session.GetString(r.Context(), "user_id")
	if userSub == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	user, err := m.App.Dynamo.UserProfile.GetUserProfile(r.Context(), userSub)
	if err != nil {
		log.Printf("Couldn't fetch user: %v", err)
		m.App.Session.Put(r.Context(), "flash", "something went wrong")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	render.Template(w, r, "choose-hemisphere.page.tmpl", &models.TemplateData{
		Data: map[string]interface{}{
			"Hemisphere": user.Hemisphere,
		},
	})
}

func (m *Repository) FishDashboardGet(w http.ResponseWriter, r *http.Request) {
	// Get the user_id (Cognito sub) from session
	userHemisphere := m.App.Session.GetString(r.Context(), "user_hemisphere")
	if userHemisphere == "" {
		http.Redirect(w, r, "/choose-hemisphere", http.StatusSeeOther)
		return
	}

	render.Template(w, r, "fish-dashboard.page.tmpl", &models.TemplateData{
		Data: map[string]interface{}{
			"Hemisphere": userHemisphere,
		},
	})
}

func (m *Repository) GetAvailableFish(w http.ResponseWriter, r *http.Request) {
	// Get the user_id (Cognito sub) from session
	userHemisphere := m.App.Session.GetString(r.Context(), "user_hemisphere")
	if userHemisphere == "" {
		http.Redirect(w, r, "/choose-hemisphere", http.StatusSeeOther)
		return
	}

	userID := m.App.Session.GetString(r.Context(), "user_id") // should be set during auth
	if userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	monthStr := r.URL.Query().Get("month")
	timeStr := r.URL.Query().Get("time")

	month, err := strconv.Atoi(monthStr)
	if err != nil || month < 1 || month > 12 {
		http.Error(w, "invalid month", http.StatusBadRequest)
		return
	}

	fish, err := m.App.Dynamo.Fish.ListAvailableFish(r.Context(), userID, month, timeStr, userHemisphere)
	if err != nil {
		log.Printf("failed to list available fish: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(fish)
}

//////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////
///////////////////// POST REQUESTS //////////////////////////
//////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////

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

func (m *Repository) ChooseHemispherePost(w http.ResponseWriter, r *http.Request) {
	if err := m.App.Session.RenewToken(r.Context()); err != nil {
		m.App.ErrorLog.Println("Session token renewal failed:", err)
	}

	err := r.ParseForm()
	if err != nil {
		m.App.ErrorLog.Println("ParseForm error:", err)
	}

	form := forms.New(r.PostForm)

	if !form.Valid() {
		render.Template(w, r, "choose-hemisphere.page.tmpl", &models.TemplateData{
			Form: form,
		})
		return
	}

	hemisphere := r.Form.Get("hemisphere")
	if hemisphere != "north" && hemisphere != "south" {
		http.Error(w, "Invalid hemisphere selected", http.StatusBadRequest)
		return
	}

	// Get the user_id (Cognito sub) from session
	userSub := m.App.Session.GetString(r.Context(), "user_id")
	if userSub == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	err = m.App.Dynamo.UserProfile.UpdateUserHemisphere(r.Context(), userSub, hemisphere)

	m.App.Session.Put(r.Context(), "flash", "hemisphere confirmed")
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func (m *Repository) UpdateUserFish(w http.ResponseWriter, r *http.Request) {
	userID := m.App.Session.GetString(r.Context(), "user_id")
	if userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var payload struct {
		FishID string `json:"fish_id"`
		Caught bool   `json:"caught"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	if payload.FishID == "" {
		http.Error(w, "missing fish_id", http.StatusBadRequest)
		return
	}

	var err error
	if payload.Caught {
		err = m.App.Dynamo.UserFish.PutCaughtFish(r.Context(), userID, payload.FishID)
	} else {
		err = m.App.Dynamo.UserFish.DeleteCaughtFish(r.Context(), userID, payload.FishID)
	}

	if err != nil {
		log.Printf("Failed to update userfish: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
