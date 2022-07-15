package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Contact struct {
	firstName   string
	lastName    string
	email       string
	phoneNumber string
	state       string
	business    bool
}

var contacts []Contact
var states = []string{
	"AK", "AL", "AR", "AZ", "CA", "CO", "CT", "DC", "DE", "FL", "GA",
	"HI", "IA", "ID", "IL", "IN", "KS", "KY", "LA", "MA", "MD", "ME",
	"MI", "MN", "MO", "MS", "MT", "NC", "ND", "NE", "NH", "NJ", "NM",
	"NV", "NY", "OH", "OK", "OR", "PA", "RI", "SC", "SD", "TN", "TX",
	"UT", "VA", "VT", "WA", "WI", "WV", "WY"}

var app = tview.NewApplication()
var text = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText("(a) to add a new contact \n(q) to quit")
var form = tview.NewForm()
var pages = tview.NewPages()
var contactsList = tview.NewList().ShowSecondaryText(false)
var contactText = tview.NewTextView()
var flex = tview.NewFlex()

func main() {
	contactsList.SetSelectedFunc(func(index int, name string, second_name string, shortcut rune) {
		ContactText(&contacts[index])
	})

	flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().
			AddItem(contactsList, 0, 1, true).
			AddItem(contactText, 0, 4, false), 0, 6, false).
		AddItem(text, 0, 1, false)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			app.Stop()
		} else if event.Rune() == 97 {
			form.Clear(true)
			addContactForm()
			pages.SwitchToPage("Add Contact")
		}
		return event
	})

	pages.AddPage("Menu", flex, true, true)
	pages.AddPage("Add Contact", form, true, false)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func addContactList() {
	contactsList.Clear()
	for i, c := range contacts {
		contactsList.AddItem(c.firstName+" "+c.lastName, "", rune('1'+i), nil)
	}
}

func addContactForm() {
	c := Contact{}
	form.AddInputField("First Name", "", 20, nil, func(s string) { c.firstName = s })
	form.AddInputField("Last Name", "", 20, nil, func(s string) { c.lastName = s })
	form.AddInputField("Email", "", 20, nil, func(s string) { c.email = s })
	form.AddInputField("Phone", "", 20, nil, func(s string) { c.phoneNumber = s })
	form.AddDropDown("State", states, 0, func(s string, i int) { c.state = s })
	form.AddCheckbox("Business", false, func(b bool) { c.business = b })

	form.AddButton("Save", func() {
		contacts = append(contacts, c)
		addContactList()
		pages.SwitchToPage("Menu")
	})
}

func ContactText(c *Contact) {
	contactText.Clear()
	contactText.SetText(fmt.Sprintf("%s %s\n%s\n%s\n",
		c.firstName, c.lastName, c.email, c.phoneNumber))
}
