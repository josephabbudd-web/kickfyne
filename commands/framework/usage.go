package framework

func Usage() (usage string) {
	usage = `🗽 GETTING STARTED WITH kickfyne.

🌐 BUILD AND RUN THE FRAMEWORK:

💲 mkdir «name-of-my-app»
💲 cd «name-of-my-app»
💲 go mod init example.com/«name-of-my-app»
💲 kickfyne framework
💲 go get fyne.io/fyne/v2@latest
💲 go mod tidy
💲 go build
💲 ./«name-of-my-app»

The framework is contained in 3 folders.
1. ./ which contains
  * main.go
  * FyneApp.toml
  * .manifest.yaml
    * This is kickfyne's manifest.
	* kickfyne will not work without the .manifest.yaml file.
2. ./frontend/ which contains:
  * settings.go which you will need to modify.
    * Define your opening screen with the const string openingScreenName.
    * Define if the main menu is used with the const bool usingMainMenu.
    * Define the main menu items with the var []_mainmenu_.MainMenuItem mainMenuItems.
  * deps/ which contains depencies for the frontend only:
    * mainmenu/ is internal main menu code you can ignore.
    * screenmap/ is the screen api maps.
    * types/ is types needed for the frontend.
  * screens/
    * each screen has it's own package folder.
3. ./deps/ which contains framework dependencies.
  * deps.go just starts up this package.
  * metadata/ makes the FyneApp.toml information available.
  * paths/ is path information.
  * thread/ determines if a thread is the main thread or not.
`
	return
}
