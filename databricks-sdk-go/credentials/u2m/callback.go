package u2m

import (
	"context"
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed page.tmpl
var pageTmpl string

type oauthResult struct {
	Error            string
	ErrorDescription string
	State            string
	Code             string
	Issuer           string
	Host             string
}

// callbackServer is a server that listens for the redirect from the Databricks
// identity provider. It renders a page.html template that shows the result of
// the authentication attempt.
type callbackServer struct {
	// ctx is the context used when waiting for the redirect from the identity
	// provider. This is needed because the Handler() method from the oauth2
	// library does not accept a context.
	ctx context.Context

	// srv is the server that listens for the redirect from the identity provider.
	srv http.Server

	// browser is a function that opens a browser to the given URL.
	browser func(string) error

	// arg is the OAuth argument used to authenticate.
	arg OAuthArgument

	// renderErrCh is a channel that receives an error if there is an error
	// rendering the page.html template.
	renderErrCh chan error

	// feedbackCh is a channel that receives the result of the authentication
	// attempt.
	feedbackCh chan oauthResult

	// tmpl is the template used to render the response page after the user is
	// redirected back to the callback server.
	tmpl *template.Template
}

// newCallbackServer creates a new callback server that listens for the redirect
// from the Databricks identity provider.
func (a *PersistentAuth) newCallbackServer() (*callbackServer, error) {
	tmpl, err := template.New("page").Funcs(template.FuncMap{
		"title": func(in string) string {
			title := cases.Title(language.English)
			return title.String(strings.ReplaceAll(in, "_", " "))
		},
	}).Parse(pageTmpl)
	if err != nil {
		return nil, err
	}
	cb := &callbackServer{
		feedbackCh:  make(chan oauthResult),
		renderErrCh: make(chan error),
		tmpl:        tmpl,
		ctx:         a.ctx,
		browser:     a.browser,
		arg:         a.oAuthArgument,
	}
	cb.srv.Handler = cb
	go func() {
		_ = cb.srv.Serve(a.ln)
	}()
	return cb, nil
}

// Close closes the callback server.
func (cb *callbackServer) Close() error {
	return cb.srv.Close()
}

// ServeHTTP renders the page.html template.
func (cb *callbackServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res := oauthResult{
		Error:            r.FormValue("error"),
		ErrorDescription: r.FormValue("error_description"),
		Code:             r.FormValue("code"),
		State:            r.FormValue("state"),
		Issuer:           r.FormValue("iss"),
		Host:             cb.getHost(),
	}
	if res.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	err := cb.tmpl.Execute(w, res)
	if err != nil {
		cb.renderErrCh <- err
	}
	cb.feedbackCh <- res
}

func (cb *callbackServer) getHost() string {
	switch a := cb.arg.(type) {
	case AccountOAuthArgument:
		return a.GetAccountHost()
	case WorkspaceOAuthArgument:
		return a.GetWorkspaceHost()
	default:
		return ""
	}
}

func (cb *callbackServer) awaitResult(authCodeURL string) (oauthResult, error) {
	err := cb.browser(authCodeURL)
	if err != nil {
		fmt.Printf("Please continue the authentication process in your browser:\n%s\n", authCodeURL)
	}
	select {
	case <-cb.ctx.Done():
		return oauthResult{}, cb.ctx.Err()
	case renderErr := <-cb.renderErrCh:
		return oauthResult{}, renderErr
	case res := <-cb.feedbackCh:
		if res.Error != "" {
			return oauthResult{}, fmt.Errorf("%s: %s", res.Error, res.ErrorDescription)
		}
		return res, nil
	}
}

// Handler opens up a browser waits for redirect to come back from the identity provider
func (cb *callbackServer) Handler(authCodeURL string) (string, string, error) {
	res, err := cb.awaitResult(authCodeURL)
	if err != nil {
		return "", "", err
	}
	return res.Code, res.State, nil
}

func (cb *callbackServer) handlerWithIssuer(authCodeURL string) (code, state, issuer string, err error) {
	res, err := cb.awaitResult(authCodeURL)
	if err != nil {
		return "", "", "", err
	}
	return res.Code, res.State, res.Issuer, nil
}
