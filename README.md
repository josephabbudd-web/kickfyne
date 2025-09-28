# kickfyne "fyne my way"

![kickfyne fyne with a kick.](/images/kick.jpeg)
Image courtesy of https://isorepublic.com/photo/flying-kick/

## Fun with Fyne

[The Fyne Toolkit](https://fyne.io/) is my favorite Go application tool kit. It is really nice. After creating a couple of apps using the fyne toolkit, I decided to make my own tool to build and manage building my fyne apps. So my tool is called kickfyne.

This is my personal project and is not in any way a part of any of the [the Fyne Toolkit projects](https://fyne.io/).

I started this years ago and then stopped while fyne made some critical changes. Now that those are made I'm back to having fun with fyne.

## Sep 28, 2025

This is a work in progress. I'm building an application with kickfyne and making changes to kickfyne according to what I learn building the application.

## To do

1. I need to add a display name to a screen. Right now I'm just using the screen's package name as you will see in the examples below.
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
＄ go get fyne.io/fyne/v2/storage/repository@v2.6.1
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
The screen is named ContactsAT and the tabs are Add Edit and Remove.

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

### Rewrite the app's main menu in the app's FyneApp.toml file.


```toml
MainMenu = "ContactsAT ContactsDT ContactsAC"
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

A panel has it's own folder where there are 2 files. The developer only works with these 2 files.
1. content.go
1. state.go

### content.go lays out the content and handles user input.

Below is the func NewContent in content.go.
1. It initializes the content heading and form members to their widgets.
1. It builds a list of form items.
1. It builds the form with it's submit handler.
1. It lays out the content using a VBox.

```go
// NewContent initializes this panel's content.
// Returns the panel's content and the error.
func NewContent(screen *_misc_.Miscellaneous) (panelContent *Content, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("ConfirmPanel.NewContent: %w", err)
		}
	}()

	planOptions := appstate.PlanOptions()
	plans := make([]string, len(planOptions))
	for i, planOption := range planOptions {
		plans[i] = planOption.Text
	}

	// Create the components of this panel's content.
	panelContent = &Content{
		screen:  screen,
		heading: widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),

		courseName:        widget.NewLabel(""),
		courseDescription: widget.NewLabel(""),
		lessonDescription: widget.NewLabel(""),
		speed:             widget.NewLabel(""),
		plan:              widget.NewLabel(""),
		status:            widget.NewLabel(""),
	}
	// Build the form.
	formItems := []*widget.FormItem{
		{
			Text:   "Name",
			Widget: panelContent.courseName,
		},
		{
			Text:   "Description",
			Widget: panelContent.courseDescription,
		},
		{
			Text:   "Speed",
			Widget: panelContent.speed,
		},
		{
			Text:   "Plan",
			Widget: panelContent.plan,
		},
		{
			Text:   "Current Lesson",
			Widget: panelContent.lessonDescription,
		},
		{
			Text:   "Status",
			Widget: panelContent.status,
		},
	}
	form := &widget.Form{
		Items: formItems,
		OnSubmit: func() {
         // Convert the record.CourseOption to a record.CourseRemove.
			courseRemove := record.ToCourseRemove(panelContent.courseOption)
         // Remove the course.
         // _api_.RemoveCourse will
         // 1. Attempt to remove the course record.
         // 2. Inform the user on the success or failure.
         // 3. Reload all course select lists on success.
			_api_.RemoveCourse(courseRemove)
		},
		SubmitText:  "Yes. Remove this course.",
		OnCancel:    func() { panelContent.screen.Panelers.Select.Show(true) },
		CancelText:  "No. Select another course.",
		Orientation: widget.Horizontal,
	}

	// Layout the components.
	panelContent.content = container.NewVBox(
		panelContent.heading,
		form,
	)
	return
}
```

### state.go mostly sets the content but also gets some content.

The setters are thread safe. Below is the setter for the form shown above. A separate getter func is not required for the form.

```go
// SetForm returns a _types_.Setter that sets the content's form and it's CourseOption record.
func (state *State) SetForm(courseOption *record.CourseOption) (setter _types_.StateSetter) {
	var status string
	if courseOption.Completed {
		status = "This course has been completed."
	} else {
		status = fmt.Sprintf("Currently working on the lesson %s.", courseOption.CurrentLessonDescription)
	}
	setter = func(isMainThread bool) (refreshCanvasObject bool) {
      // The form is part of the panel's content so the canvas object must be refreshed.
		refreshCanvasObject = true
      // Remember this course option record.
		state.content.courseOption = courseOption
      // A func to set the course information for display.
		set := func() {
			state.content.courseName.SetText(courseOption.Name)
			state.content.courseDescription.SetText(courseOption.Description)
			state.content.lessonDescription.SetText(courseOption.CurrentLessonDescription)
			state.content.speed.SetText(courseOption.SpeedDescription)
			state.content.plan.SetText(courseOption.PlanDescription)
			state.content.status.SetText(status)
		}
      // Set the course information for display.
		if isMainThread {
			set()
		} else {
			fyne.Do(set())
		}
		return
	}
	return
}
```

