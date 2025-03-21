//our data model

package main

import "github.com/charmbracelet/bubbles/filepicker"

type Model struct {
	filepicker   filepicker.Model
	selectedFile string
	quitting     bool
	err          error
	current      int
	selected     map[int]struct{}
	choices      []string
}
