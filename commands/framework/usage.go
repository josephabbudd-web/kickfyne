package framework

func Usage() (usage string) {
	usage = `🗽 GETTING STARTED WITH kickfyne.

＄ cd ~/projects/«name-of-my-app»
＄ go mod init "example.com/«name-of-my-app»" or "github.com/«my-github-folder»/«name-of-my-app»"
＄ kickfyne framework

🌐 THE FRAMEWORK:
The framework is contained in 3 folders.
1. ./ which contains main.go and FyneApp.toml
3. ./frontend/ which contains the front-end code.
4. ./deps/ which contains deps code.

🔨 BUILDING THE APP:
You can build the app after running the command "＄ kickfyne framework".
The following build example is done in the app's folder.

＄ go mod tidy
＄ go build
＄ ./«name-of-executable»
`
	return
}
