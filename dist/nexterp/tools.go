package main

import "net/http"

func (a *application) Get(s string, h http.HandlerFunc) {
	a.App.Routes.Get(s, h)
}

func (a *application) Post(s string, h http.HandlerFunc) {
	a.App.Routes.Post(s, h)
}
