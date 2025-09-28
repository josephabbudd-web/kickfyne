package root

type fyneAppTOMLData struct {
	WebSiteURL string // "https://github.com/josephbudd/okp"
	AppName    string // "OKP"
	AppID      string // "com.github.josephbudd.okp"
}

var dyneAppTOMLTemplate = `# Website = "{{ .WebSiteURL }}"

[Details]
# ID is the unique ID of this application, used by many distribution platforms.
ID = "{{ .AppID }}"
# Name is the human friendly name of this app.
Name = "{{ .AppName }}"
# Version represents the version of this application, normally following semantic versioning.
Version = "0.1.0"
# Build is the build number of this app, some times appended to the version number.
Build = 1
# Icon contains, if present, a resource of the icon that was bundled at build time.
# Icon = Resource "Icon.png"
# Release if true this binary was build in release mode
Release = false

# Custom is now Development. But it's still app.Metadata().Custom.
[Development]
# MainMenu is a space separated sorted list of screen package names that you want in the main menu.
# The first screen package name is the opening screen.
#   The screen package names are
#     the screen package names in frontend/screen/,
#     the keys of the map screenmap.Map in frontend/screenmap/screenmap.go.
# Names in MainMenu that are not found are ignored and logged.
MainMenu = "HelloWorld"
`
