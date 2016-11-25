package main

import (
	"gnd.la/app"
	"gnd.la/apps/users"
	"gnd.la/config"
	_ "gnd.la/frontend/bootstrap3"
	_ "gnd.la/orm/driver/sqlite"
	"gnd.la/social/facebook"
	"gnd.la/social/twitter"
	"gnd.la/util/pathutil"
)

var (
	App *app.App
)

var Config struct {
	// Add facebook-app = <your-app-token> to enabled Facebook integration.
	// See also the User type in models.go
	FacebookApp *facebook.App
	// Add twitter-app = <your-app-token>:<your-app-secret> to enabled Twitter integration.
	// See also the User type in models.go
	TwitterApp *twitter.App

	SiteName string

	// Include your settings here
}

func init() {
	// Initialize the configuration and the App in init, so
	// it's configured correctly when running tests.
	config.Register(&Config)
	config.MustParse()

	App = app.New()
	// Make the config available to templates as @Config
	App.AddTemplateVars(map[string]interface{}{
		"Config": &Config,
	})
	// Asset handling
	App.HandleAssets("/assets/", pathutil.Relative("assets"))

	// You might probably want the following if you're
	// deploying your app behind an upstream proxy.
	//
	// App.SetTrustXHeaders(true)

	// Attach users app
	usersApp := users.New(users.Options{
		SiteName: Config.SiteName,
		UserType: User{},
	})
	usersApp.ContainerTemplateName = "users-base.html"
	usersApp.Attach(App)
	// Site handlers
	App.Handle("^/$", MainHandler, app.NamedHandler("main"))

	// Error handler, for 404
	App.SetErrorHandler(ErrorHandler)
}
