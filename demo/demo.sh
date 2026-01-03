# How To Run this kickfyne demonstration.
# $ cp -r ./demo ~/my_projects
# $ cd ~/my_projects/demo
# $ .demo.sh

clear
set -e # break on error.

if [[ -e logs ]] then
	rm -r ./logs
fi
mkdir logs

if [[ -e bin ]] then
	rm -r ./bin
fi
mkdir bin
mkdir bin/images
cp ./app-icon.jpeg ./bin/images/app-icon.jpeg

if [[ -e demo ]] then
	rm -r ./demo
fi
mkdir demo
cd demo

go mod init example.com/demo
kickfyne framework >&../logs/Framework.log

kickfyne screen add-accordion Accordion Add Subtract >&../logs/Add-Accordion.log
kickfyne screen add-apptabs+ AppTabs Add Subtract >&../logs/Add-AppTabs.log
kickfyne screen add-border Border Top Bottom Left Right Center >&../logs/Add-Border.log
kickfyne screen add-doctabs+ DocTabs Add Subtract >&../logs/Add-DocTabs.log
kickfyne screen add-split Split Leading Trailing >&../logs/Add-Split.log

kickfyne screen add-split SplitAccordionBorder Leading=*Accordion Trailing=*Border >&../logs/Add-SplitAccordionBorder.log
kickfyne screen add-split SplitAppTabsDocTabs Leading=*AppTabs Trailing=*DocTabs >&../logs/Add-SplitAppTabsDocTabs.log
kickfyne screen add-split SplitHelloWorld Leading Trailing=*HelloWorld >&../logs/Add-SplitHelloWorld.log
kickfyne screen add-split SplitSplit Leading=*Split Trailing >&../logs/Add-SplitSplit.log

kickfyne screen add-border BorderAccordion Top Bottom Left Right Center=*Accordion >&../logs/Add-BorderAccordion.log
kickfyne screen add-border BorderAppTabs Top Bottom Left Right Center=*AppTabs >&../logs/Add-BorderAppTabs.log
kickfyne screen add-border BorderBorder Top Bottom Left Right Center=*Border >&../logs/Add-BorderBorder.log
kickfyne screen add-border BorderDocTabs Top Bottom Left Right Center=*DocTabs >&../logs/Add-BorderDocTabs.log
kickfyne screen add-border BorderSplit Top Bottom Left Right Center=*Split >&../logs/Add-BorderSplit.log

kickfyne screen add-apptabs+ AppTabsFull *HelloWorld *Accordion *AppTabs *Border *DocTabs *Split >&../logs/Add-AppTabsFull.log
kickfyne screen add-doctabs+ DocTabsFull *HelloWorld *Accordion *AppTabs *Border *DocTabs *Split >&../logs/Add-DocTabsFull.log

go mod tidy

# settings.go contains the main menu information.
cp ../settings_go_demo.txt ./frontend/settings.go
cp ./FyneApp.toml ../bin

go build -o ../bin/demo >&../logs/build.log

# Clean up.
cd ..
rm -r ./demo
rm -r ./logs

# Run the demo.
cd ./bin
./demo

