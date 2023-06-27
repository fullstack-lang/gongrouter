package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *EditorOutlet:
		if stage.OnAfterEditorOutletCreateCallback != nil {
			stage.OnAfterEditorOutletCreateCallback.OnAfterCreate(stage, target)
		}
	case *Outlet:
		if stage.OnAfterOutletCreateCallback != nil {
			stage.OnAfterOutletCreateCallback.OnAfterCreate(stage, target)
		}
	case *TableOutlet:
		if stage.OnAfterTableOutletCreateCallback != nil {
			stage.OnAfterTableOutletCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *EditorOutlet:
		newTarget := any(new).(*EditorOutlet)
		if stage.OnAfterEditorOutletUpdateCallback != nil {
			stage.OnAfterEditorOutletUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Outlet:
		newTarget := any(new).(*Outlet)
		if stage.OnAfterOutletUpdateCallback != nil {
			stage.OnAfterOutletUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TableOutlet:
		newTarget := any(new).(*TableOutlet)
		if stage.OnAfterTableOutletUpdateCallback != nil {
			stage.OnAfterTableOutletUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *EditorOutlet:
		if stage.OnAfterEditorOutletDeleteCallback != nil {
			staged := any(staged).(*EditorOutlet)
			stage.OnAfterEditorOutletDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Outlet:
		if stage.OnAfterOutletDeleteCallback != nil {
			staged := any(staged).(*Outlet)
			stage.OnAfterOutletDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TableOutlet:
		if stage.OnAfterTableOutletDeleteCallback != nil {
			staged := any(staged).(*TableOutlet)
			stage.OnAfterTableOutletDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *EditorOutlet:
		if stage.OnAfterEditorOutletReadCallback != nil {
			stage.OnAfterEditorOutletReadCallback.OnAfterRead(stage, target)
		}
	case *Outlet:
		if stage.OnAfterOutletReadCallback != nil {
			stage.OnAfterOutletReadCallback.OnAfterRead(stage, target)
		}
	case *TableOutlet:
		if stage.OnAfterTableOutletReadCallback != nil {
			stage.OnAfterTableOutletReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *EditorOutlet:
		stage.OnAfterEditorOutletUpdateCallback = any(callback).(OnAfterUpdateInterface[EditorOutlet])
	
	case *Outlet:
		stage.OnAfterOutletUpdateCallback = any(callback).(OnAfterUpdateInterface[Outlet])
	
	case *TableOutlet:
		stage.OnAfterTableOutletUpdateCallback = any(callback).(OnAfterUpdateInterface[TableOutlet])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *EditorOutlet:
		stage.OnAfterEditorOutletCreateCallback = any(callback).(OnAfterCreateInterface[EditorOutlet])
	
	case *Outlet:
		stage.OnAfterOutletCreateCallback = any(callback).(OnAfterCreateInterface[Outlet])
	
	case *TableOutlet:
		stage.OnAfterTableOutletCreateCallback = any(callback).(OnAfterCreateInterface[TableOutlet])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *EditorOutlet:
		stage.OnAfterEditorOutletDeleteCallback = any(callback).(OnAfterDeleteInterface[EditorOutlet])
	
	case *Outlet:
		stage.OnAfterOutletDeleteCallback = any(callback).(OnAfterDeleteInterface[Outlet])
	
	case *TableOutlet:
		stage.OnAfterTableOutletDeleteCallback = any(callback).(OnAfterDeleteInterface[TableOutlet])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *EditorOutlet:
		stage.OnAfterEditorOutletReadCallback = any(callback).(OnAfterReadInterface[EditorOutlet])
	
	case *Outlet:
		stage.OnAfterOutletReadCallback = any(callback).(OnAfterReadInterface[Outlet])
	
	case *TableOutlet:
		stage.OnAfterTableOutletReadCallback = any(callback).(OnAfterReadInterface[TableOutlet])
	
	}
}
