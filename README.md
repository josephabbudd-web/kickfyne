# kickfyne "fyne my way"

![kickfyne fyne with a kick.](/images/kick.jpeg)
Image courtesy of https://isorepublic.com/photo/flying-kick/

## November 18, 2025

Added the Border screen.
Added the system tray menu.

## Fun with Fyne

[The Fyne Toolkit](https://fyne.io/) is my favorite Go application tool kit. It is really nice. After creating a couple of apps using the fyne toolkit, I decided to make my own tool to build and manage building my fyne apps. So my tool is called kickfyne.

This is my personal project and is not in any way a part of any of the [Fyne Toolkit projects](https://fyne.io/).

I started this years ago and then stopped while the fyne team made some critical changes. Now that those are made I'm back to having fun with fyne.

kickfyne is a work in progress. I'm using kickfyne to build my morse code trainer. As I build the trainer I am

* fixing bugs in kickfyne.
* improving documentation in kickfyne.
* removing features I don't want anymore.
* adding features I do want.

## Summary

With kickfyne I create the framework and then add and remove screens in the framework. A screen is a package that lays out panels.

A panel is designed to be simple and do only one thing. A panel has it's own folder where there are 3 files.
1. content.go lays out the content and handles user input.
1. state.go manages the complexities of a updating content and can update content using a preset.
1. preset.go contains the different initial settings for all or part of state. Usually the **Default** preset is all that is needed.

## Example: Code a GUI application with kickfyne

### Create the framework.

The framework always works when modified with kickfyne.

```shell
💲 mkdir mycrud
💲 cd mycrud
💲 go mod init example.com/mycrud
💲 kickfyne framework
💲 go get fyne.io/fyne/v2/storage/repository@v2.6.3
💲 go mod tidy
💲 go build
💲 ./mycrud
```

### HelloWorld is the kickfyne framework's opening screen

The kickfyne framework, when built and run, will open with it's default **HelloWorld** screen. The screen has 2 panels **Hello** and **HelloAgain**.

The **HelloWorld** screen is a Simple screen. A simple screen has 1 or more panels and only displays one panel at a time.

Each panel in a Simple screen uses a Select widget which allows the developer to switch to another panel. That functionality allows the developer see that each of the panels are in fact there. Especially useful after adding and removing panels in a Simple screen.

![The new application's default opening screen.](/images/kickfyne_helloworld_screen.png)

I will remove the **HelloWorld** screen after I create my real opening screen.

### Add a simple screen to Edit a contact.

A simple screen layout has 1 or more panels but only displays one panel at a time.

This Simple screen is named **Edit** and will have 2 panels, **Select** and **Edit**.
The first panel, which is the Select panel is the default panel. In a real application, the Select panel would display a select list and the Edit panel would display the selected record in a form and allow the user to edit and save or cancel and go back to the select form.

```shell
💲 kickfyne screen add-simple Edit Select Edit
💲 go mod tidy
```

### Add a simple screen to Remove a contact.

Again, a simple screen layout has 1 or more panels but only displays one panel at a time.

The Simple screen is named **Remove** and will have 2 panels, **Select** and **Remove**.
The first panel, which is the Select panel is the default panel. In a real application, the Select panel would display a select list and the Remove panel would display the selected record and allow the user to remove or cancel and go back to the select form.

```shell
💲 kickfyne screen add-simple Remove Select Remove
💲 go mod tidy
```

### Add an AppTabs screen.

The AppTabs screen lays out a tabbar of TabItems.
None of the TabItems can be removed by the user.
The tabbar can be positioned at the top, right, bottom or left.
The AppTabs screen's API allows TabItems to be added and removed.

This AppTabs screen is named **ContactsAT** and its TabItems are named **Add** **Edit** and **Remove**.

In the command line
 * The TabItem named Add will gets it's content from it's own Add panel.
 * The TabItem named Edit is prefixed with *. So it's content comes from a new instance of the Edit screen I previously made.
 * The TabItem named Remove is prefixed with *. So it's content comes from a new instance of the Remove screen I previously made.

```shell
💲 kickfyne screen add-apptabs ContactsAT Add *Edit *Remove
💲 go mod tidy
```

### Add a DocTabs screen.

The DocTabs screen lays out a tabbar of TabItems.
Any TabItem can be removed by the user.
The tabbar can be positioned at the top, right, bottom or left.
The DocTabs screen's API also allows TabItems to be added and removed.

This DocTabs screen is named **ContactsDT** and its TabItems are named **Add** **Edit** and **Remove**.

In the command line
 * The TabItem name Add will gets it's content from it's own Add panel.
 * The TabItem name Edit is prefixed with *. So it's content comes from another instance of the Edit screen I previously made.
 * The TabItem name Remove is prefixed with *. So it's content comes from another instance of the Remove screen I previously made.

```shell
💲 kickfyne screen add-doctabs ContactsDT Add *Edit *Remove
💲 go mod tidy
```

### Add an Accordion screen.

The Accordion screen lays out AccordionItems vertically. Each AccordionItem can be opened to present its content.
The Accordion screen's API also allows AccordionItems to be added and removed.

This Accordion screen is named **ContactsAC** and its AccordionItems are named **Add**, **Edit** and **Remove**.

In the command line
 * The label name Add will gets it's content from it's own Add panel.
 * The label name Edit is prefixed with *. So it's content comes from another instance of the Edit screen I previously made.
 * The label name Remove is prefixed with *. So it's content comes from another instance of the Remove screen I previously made.

```shell
💲 kickfyne screen add-accordion ContactsAC Add *Edit *Remove
💲 go mod tidy
```

### Add a Border screen.

A Border Screen lays out content in it's 5 areas, Top, Bottom, Left, Right and Center. The Top, Bottom, Left, and Right areas will only consume the content from their panel with the same name. The Center area can be set to use content from it's panel with the same name or from another screen.

The Border screen's API also allows the 5 areas to be added and removed.

This Border screen is named **ContactsB** and I'll use all 5 areas for demonstration. I'll use the content from the **ContactsAC** screen for the Center area.

In the command line
 * The Top area will gets it's content from it's own Top panel.
 * The Bottom area will gets it's content from it's own Bottom panel.
 * The Left area will gets it's content from it's own Left panel.
 * The Right area will gets it's content from it's own Right panel.
 * The Center area will gets it's content from an implementation of the **ContactsAC** screen.
 
```shell
💲 kickfyne screen add-border ContactsB Top Bottom Left Right Center=*ContactsAC
💲 go mod tidy
```

### Rewrite the app's opening screen and main menu in the file at frontend/settings.go

1. I want the application to open with the **ContactsAT** screen. So I will set the constant `openingScreenName` in frontend/settings.go to "ContactsAT". I will leave `openingScreenPresetName` set to "Default".
1. var mainMenuItems needs to reference the 3 new main screens.

```go
package frontend

import (
	_mainmenu_ "example.com/toolbar/frontend/deps/mainmenu"
)

const (
	// usingMainMenu. Is the application using a main menu.
	usingMainMenu = true

	// This is the screen that the application opens with.
	// It does not have to be referenced in var mainMenuItems.
	openingScreenName       = "ContactsAT"
	openingScreenPresetName = "Default"
)

// mainMenuItems is the list of items for the main menu.
// The following issues will generate an error.
//   - Repeated Label.
//   - Unknown ScreenName.
//   - Unknown PresetName.
var mainMenuItems = []_mainmenu_.MainMenuItem{
	{
		label:  "App Tabs",
		screen: "ContactsAT",
		preset: "Default",
	},
	{
		label:  "Doc Tabs",
		screen: "ContactsDT",
		preset: "Default",
	},
	{
		label:  "Accordion",
		screen: "ContactsAC",
		preset: "Default",
	},
	{
		label:  "Border",
		screen: "ContactsB",
		preset: "Default",
	},
}
```

### Remove that default HelloWorld screen

```shell
💲 kickfyne screen remove HelloWorld
```

## Example: Build and run the app

```shell
💲 go build
💲 ./mycrud
```

### The new opening screen. The ContactsAT screen

The **ContactsAT** screen is an AppTabs screen. The user can not close an AppTabs TabItem. The **ContactsAT** screen is shown below with its **Add** TabItem selected. The **Add** TabItem and displaying the content from it's **Add** panel.

![The AppTabs screen.](/images/kickfyne_contacts_app_tabs_screen.png)

### The ContactsDT screen

The **ContactsDT** screen is an DocTabs screen. The **ContactsDT** screen is shown below with its **Edit** TabItem selected. The **Edit** TabItem is displaying the content from its own instance of the **Edit** screen. The **Edit** screen is displaying its default **Select** panel.

![The AppTabs screen.](/images/kickfyne_contacts_doc_tabs_screen.png)

### The ContactsAC screen

The **ContactsAC** screen is an Accordion screen. The **ContactsAC** screen is shown below with its **Remove** AccordionItem selected. The **Remove** AccordionItem is displaying the content from its own instance of the **Remove** screen. The **Remove** screen is displaying its default **Select** panel.

![The Accordion screen.](/images/kickfyne_contacts_accordion_screen.png)

### The ContactsB screen

The **ContactsB** screen is a Border screen. Notice that the Top, Left, Right and Bottom areas also show buttons by default. That's to make them easier to identify. Notice that the Center area is displaying content from the **ContactsAC** screen which has its **Edit** AccordionItem selected.

![The Border screen.](/images/kickfyne_contacts_border_screen.png)

### The main menu

The main menu is at the top left of the app and it is shown open.

![The main menu opened showing the names of the screens shown above.](images/kickfyne_main_menu.png)


## Screen modification

1. AccordionItems can be appended to and removed from an Accordion screen.
1. TabItems can be appended to and removed from an AppTab screen.
1. TabItems can be appended to and removed from a DocTab screen.
1. Areas can be added to and removed from a Border screen.
1. Panels can be added to and removed from a simple screen.

```shell
💲 kickfyne screen add-item «screen-package-name» «[*]item name» ...
💲 kickfyne screen remove-item «screen-package-name» «[*]item name» ...
```

## User configuration with AppTabs and DocTabs screens

With user configuration

 * The first TabItem following the settings TabItem is selected at startup.
 * After the user selects the settings TabItem, makes a change and clicks the dismiss button, the settings TabItem gets unselected and the previously selected TabItem is selected. 

By adding a **+** to the verbs `add-apptabs` and `add-apptabs` a tabbar screen is created with a settings TabItem that allows the user to set where the tabs are located.

```shell
💲 kickfyne screen add-apptabs+ «screen-package-name» «[*]tab-item-name, ...»
💲 kickfyne screen add-doctabs+ «screen-package-name» «[*]tab-item-name, ...»
```

Below is an example showing the selected config TabItem with its content.

![The new configurable AppTabs screen.](/images/kickfyne_apptabs_config.png)

## To do

1. Adding the ability to safely undo changes when an error occurs.
1. Review my API for adding and removing tabs and accordion items.
1. Review my inline and output documentation.
1. Write tests.
