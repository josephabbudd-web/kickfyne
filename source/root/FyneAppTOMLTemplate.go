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

# Development in this toml file = app.Metadata().Custom in golang.
# All members must be a string.
# Invalid: MyBool = false
# Valid:   MyBool = "false"
[Development]
`
