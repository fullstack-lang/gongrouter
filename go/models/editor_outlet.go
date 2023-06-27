package models

// EditorOutlet because router cannot work (cause in front collides)
type EditorOutlet struct {
	Name string

	EditorType EditorType

	// in case it is an updater editor
	UpdatedObjectID int
}
