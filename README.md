# kickfyne "fyne my way"

![kickfyne fyne with a kick.](/images/kick.jpeg)
Image courtesy of https://isorepublic.com/photo/flying-kick/

## Fun with Fyne

[The Fyne Toolkit](https://fyne.io/) is my favorite Go application tool kit. It is really nice. After creating a couple of apps using the fyne toolkit, I decided to make my own tool to build and manage building my fyne apps. So my tool is called kickfyne.

This is my personal project and is not in any way a part of any of the [the Fyne Toolkit projects](https://fyne.io/).

I started this years ago and then stopped while fyne made some critical changes. Now that those are made I'm back to having fun with fyne.

## October 7, 2025

This is a work in progress. I'm building an application with kickfyne and making changes to kickfyne according to what I learn building the application.

The main menu is more flexible. Each unique menuitem in the main menu can open the same screen to display different content.

Additional changes regarding presets.
Fixed a bug in panel state func Set.

## To do

1. Review my API for adding and removing tabs and accordion items.
1. Review my inline and output documentation.

## Summary 
kickfyne allows a developer to create a running framework and then

1. Add and remove screens to the GUI. Screens are layouts of panels. A panel lays out content and controls user input. Currently there is
   1. An AppTabs layout. A tabbar layout in which the user can not close any tabs.
   1. A DocTabs layout. A tabbar layout in which the user can close any and all tabs.
   1. An Accordion layout.
   1. A simple layout. A layout where only one of the panels is displayed at a time.

## Example

### Create the framework.

The framework always works when modified with kickfyne.

```shell
＄ mkdir mycrud
＄ cd mycrud
＄ go mod init example.com/mycrud
＄ kickfyne framework
＄ go get fyne.io/fyne/v2/storage/repository@v2.6.3
＄ go mod tidy
＄ go build
＄ ./mycrud
```

### HelloWorld is the kickfyne framwork's opening screen

The kickfyne framwork, when built and run, will open with it's default **HelloWorld** screen. The screen has 2 panels with over size buttons for switching back and forth. The **HelloWorld** screen is a simple screen. A simple screen has 1 or more panels and only displays one panel at a time.

The two panels are **Hello** and **HelloAgain**. 
![The new application's default opening screen.](/images/kickfyne_helloworld_screen.png)

I will remove the **HelloWorld** screen when I create my real opening screen. Then I will also edit the framework's main menu, removing "HelloWorld" from the main menu and adding my real opening screen's name.

### Add a simple screen to Edit a contact.

A simple screen layout has 1 or more panels but only displays one panel at a time.

The screen is named Edit and will have 2 panels, Select and Edit.
The Select panel is the default panel. In a real application, the Select panel would display a select list and the Edit panel would display the selected record in a form and allow the user to edit and save or cancel and go back to the select form.

```shell
＄ kickfyne screen add-simple Edit Select Edit
＄ go mod tidy
```

### Add a simple screen to Remove a contact.

Again, a simple screen layout has 1 or more panels but only displays one panel at a time.

The screen is named Remove and will have 2 panels, Select and Remove.
The Select panel is the default panel. In a real application, the Select panel would display a select list and the Remove panel would display the selected record and allow the user to remove or cancel and go back to the select form.

```shell
＄ kickfyne screen add-simple Remove Select Remove
＄ go mod tidy
```

### Add an AppTabs tabbar screen.

The AppTabs tabbar layout uses tabs that the user can not close.
The API allows you the developer, to open (add) and close (remove) tabs.

An AppTabs screen lays out an AppTabs tab bar.
The screen is named ContactsAT and the tabs are Add, Edit and Remove.

In the command line
 * The tab named Add will gets it's content from it's own Add panel.
 * The tab named Edit is prefixed with *. So it's content comes from a new instance of the Edit screen I previously made.
 * The tab named Remove is prefixed with *. So it's content comes from a new instance of the Remove screen I previously made.

```shell
＄ kickfyne screen add-apptabs ContactsAT Add *Edit *Remove
＄ go mod tidy
```

### Add a DocTabs tabbar screen.

The DocTabs tabbar layout uses tabs that the user can close.
The API allows you the developer, to open (add) and close (remove) tabs.

A DocTabs screen lays out a DocTabs tab bar.
The screen is named ContactsDT and the tabs are Add Edit and Remove.

In the command line
 * The tab name Add will gets it's content from it's own Add panel.
 * The tab name Edit is prefixed with *. So it's content comes from another instance of the Edit screen I previously made.
 * The tab name Remove is prefixed with *. So it's content comes from another instance of the Remove screen I previously made.

```shell
＄ kickfyne screen add-doctabs ContactsDT Add *Edit *Remove
＄ go mod tidy
```

### Add an Accordion screen.

The Accordion uses labels listed vertically. Each label unfolds to present its content.
The API allows you the developer, to open (add) and close (remove) labels.

An Accordion screen lays out a vertical list of folded content labels.
The screen is named ContactsAC and the labels are Add Edit and Remove.

In the command line
 * The label name Add will gets it's content from it's own Add panel.
 * The label name Edit is prefixed with *. So it's content comes from another instance of the Edit screen I previously made.
 * The label name Remove is prefixed with *. So it's content comes from another instance of the Remove screen I previously made.

```shell
＄ kickfyne screen add-accordion ContactsAC Add *Edit *Remove
＄ go mod tidy
```

### Rewrite the app's main menu in the file at frontend/mainmenu/mainmenu.go.

I remove the mainMenuItem for HellowWorld and add the mainMenuItem for the 3 new main screens.

```go
// The first mainMenuItem is also the opening screen.
// The following items in mainMenuItems are ignored and logged without an error.
//   - Repeated labels.
//   - Invalid screen package names.
//   - Invalid presets names.
var mainMenuItems = []mainMenuItem{
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
}
```

Because ContactsAT is the first screen named in MainMenu, the application will open with it.

### Remove that default HelloWorld screen and build the app.

```shell
＄ kickfyne screen remove HelloWorld
＄ go build
＄ ./mycrud
```

### The ContactsAT screen

The **ContactsAT** screen is an AppTabs screen. The user can not close an AppTabs tab. The **ContactsAT** screen is shown below with its default panel **Add** displayed.

![The AppTabs screen.](/images/kickfyne_contacts_app_tabs_screen.png)

### The ContactsDT screen

The **ContactsDT** screen is an DocTabs screen. The **ContactsDT** screen is shown below with its **Edit** tab displayed. The **Edit** tab is using its own instance of the **Edit** screen for content. The **Edit** screen is displaying its default **Select** panel.

![The AppTabs screen.](/images/kickfyne_contacts_doc_tabs_screen.png)

### The ContactsAC screen

The **ContactsAC** screen is an Accordion screen. The **ContactsAC** screen is shown below with its **Remove** label opened showing its content. The **Remove** label is using its own instance of the **Remove** screen for content. The **Remove** screen is displaying its default **Select** panel.

![The Accordion screen.](/images/kickfyne_contacts_accordion_screen.png)

### The main menu

The main menu is at the top left of the app and it is shown open.

![The main menu opened showing the names of the screens shown above.](images/kickfyne_main_menu.png)

## So what is a panel?

A panel has it's own folder where there are 2 files.
1. content.go lays out the content and handles user input.
1. state.go manages the complexities of a changing state and can set state using a preset.

## Presets indicate how a screen should initialize itself.

### Each screen package

 * returns it's presets in func Presets() (presets map[string]any) in api.go.
 * defines its presets in its presets/presets.go.

### Each screen preset

 * is identified by its name. The name for the default preset is **Default**.
 * contains presets for each package panel.

 ### AppTab, DocTab and Accordion screen packages

 * also load the preset of each separate screen package that is used for content. In the example above those separate screens are **Edit** and **Remove**.
