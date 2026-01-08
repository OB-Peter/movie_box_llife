package main

import (
	"errors"
	"net/http"
	"time"

	"obpeterapp.com/internal/data"
	"obpeterapp.com/internal/validator"
)

func (app *application) CreatedAuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := app.readJSON(w, r, &input); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	data.ValidateEmail(v, input.Email)
	data.ValidatePasswordPlaintext(v, input.Password)

	if !v.Valid() {
		app.failedvalidationResponse(w, r, v.Errors) // consider renaming to failedValidationResponse
		return
	}

	// NOTE: This handler's body behaves like LOGIN, not REGISTER.
	// It checks existing user and password, then issues a token.
	user, err := app.models.Users.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.invalidCredentialsResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	match, err := user.Password.Matches(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if !match {
		app.invalidCredentialsResponse(w, r)
		return
	}

	token, err := app.models.Tokens.New(user.ID, 24*time.Hour, data.ScopeAuthentication)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err := app.writeJSON(w, http.StatusCreated, envelope{"authentication_token": token}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

const (
	ScopeActivation     = "activation"
	ScopeAuthentication = "authentication" // Include a new authentication scope.
)

// Add struct tags to control how the struct appears when encoded to JSON.
type Token struct {
	Plaintext string    `json:"token"`
	Hash      []byte    `json:"-"`
	UserID    int64     `json:"-"`
	Expiry    time.Time `json:"expiry"`
	Scope     string    `json:"-"`
}
