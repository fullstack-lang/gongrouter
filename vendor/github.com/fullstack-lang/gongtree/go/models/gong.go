// generated by ModelGongFileTemplate
package models

import (
	"errors"
	"fmt"
	"sync"
)

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	GetFields() (res []string)
	GetFieldStringValue(fieldName string) (res string)
}

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct { // insertion point for definition of arrays registering instances
	Buttons           map[*Button]any
	Buttons_mapString map[string]*Button

	OnAfterButtonCreateCallback OnAfterCreateInterface[Button]
	OnAfterButtonUpdateCallback OnAfterUpdateInterface[Button]
	OnAfterButtonDeleteCallback OnAfterDeleteInterface[Button]
	OnAfterButtonReadCallback   OnAfterReadInterface[Button]

	Nodes           map[*Node]any
	Nodes_mapString map[string]*Node

	OnAfterNodeCreateCallback OnAfterCreateInterface[Node]
	OnAfterNodeUpdateCallback OnAfterUpdateInterface[Node]
	OnAfterNodeDeleteCallback OnAfterDeleteInterface[Node]
	OnAfterNodeReadCallback   OnAfterReadInterface[Node]

	Trees           map[*Tree]any
	Trees_mapString map[string]*Tree

	OnAfterTreeCreateCallback OnAfterCreateInterface[Tree]
	OnAfterTreeUpdateCallback OnAfterUpdateInterface[Tree]
	OnAfterTreeDeleteCallback OnAfterDeleteInterface[Tree]
	OnAfterTreeReadCallback   OnAfterReadInterface[Tree]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *StageStruct,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *StageStruct,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *StageStruct, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *StageStruct,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitButton(button *Button)
	CheckoutButton(button *Button)
	CommitNode(node *Node)
	CheckoutNode(node *Node)
	CommitTree(tree *Tree)
	CheckoutTree(tree *Tree)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

var _stage *StageStruct

var once sync.Once

func GetDefaultStage() *StageStruct {
	once.Do(func() {
		_stage = NewStage()
	})
	return _stage
}

func NewStage() (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		Buttons:           make(map[*Button]any),
		Buttons_mapString: make(map[string]*Button),

		Nodes:           make(map[*Node]any),
		Nodes_mapString: make(map[string]*Node),

		Trees:           make(map[*Tree]any),
		Trees_mapString: make(map[string]*Tree),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here
	}

	return
}

func (stage *StageStruct) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *StageStruct) Commit() {
	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Button"] = len(stage.Buttons)
	stage.Map_GongStructName_InstancesNb["Node"] = len(stage.Nodes)
	stage.Map_GongStructName_InstancesNb["Tree"] = len(stage.Trees)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Button"] = len(stage.Buttons)
	stage.Map_GongStructName_InstancesNb["Node"] = len(stage.Nodes)
	stage.Map_GongStructName_InstancesNb["Tree"] = len(stage.Trees)

}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts button to the model stage
func (button *Button) Stage(stage *StageStruct) *Button {
	stage.Buttons[button] = __member
	stage.Buttons_mapString[button.Name] = button

	return button
}

// Unstage removes button off the model stage
func (button *Button) Unstage(stage *StageStruct) *Button {
	delete(stage.Buttons, button)
	delete(stage.Buttons_mapString, button.Name)
	return button
}

// commit button to the back repo (if it is already staged)
func (button *Button) Commit(stage *StageStruct) *Button {
	if _, ok := stage.Buttons[button]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitButton(button)
		}
	}
	return button
}

// Checkout button to the back repo (if it is already staged)
func (button *Button) Checkout(stage *StageStruct) *Button {
	if _, ok := stage.Buttons[button]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutButton(button)
		}
	}
	return button
}

// for satisfaction of GongStruct interface
func (button *Button) GetName() (res string) {
	return button.Name
}

// Stage puts node to the model stage
func (node *Node) Stage(stage *StageStruct) *Node {
	stage.Nodes[node] = __member
	stage.Nodes_mapString[node.Name] = node

	return node
}

// Unstage removes node off the model stage
func (node *Node) Unstage(stage *StageStruct) *Node {
	delete(stage.Nodes, node)
	delete(stage.Nodes_mapString, node.Name)
	return node
}

// commit node to the back repo (if it is already staged)
func (node *Node) Commit(stage *StageStruct) *Node {
	if _, ok := stage.Nodes[node]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitNode(node)
		}
	}
	return node
}

// Checkout node to the back repo (if it is already staged)
func (node *Node) Checkout(stage *StageStruct) *Node {
	if _, ok := stage.Nodes[node]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutNode(node)
		}
	}
	return node
}

// for satisfaction of GongStruct interface
func (node *Node) GetName() (res string) {
	return node.Name
}

// Stage puts tree to the model stage
func (tree *Tree) Stage(stage *StageStruct) *Tree {
	stage.Trees[tree] = __member
	stage.Trees_mapString[tree.Name] = tree

	return tree
}

// Unstage removes tree off the model stage
func (tree *Tree) Unstage(stage *StageStruct) *Tree {
	delete(stage.Trees, tree)
	delete(stage.Trees_mapString, tree.Name)
	return tree
}

// commit tree to the back repo (if it is already staged)
func (tree *Tree) Commit(stage *StageStruct) *Tree {
	if _, ok := stage.Trees[tree]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTree(tree)
		}
	}
	return tree
}

// Checkout tree to the back repo (if it is already staged)
func (tree *Tree) Checkout(stage *StageStruct) *Tree {
	if _, ok := stage.Trees[tree]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTree(tree)
		}
	}
	return tree
}

// for satisfaction of GongStruct interface
func (tree *Tree) GetName() (res string) {
	return tree.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMButton(Button *Button)
	CreateORMNode(Node *Node)
	CreateORMTree(Tree *Tree)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMButton(Button *Button)
	DeleteORMNode(Node *Node)
	DeleteORMTree(Tree *Tree)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.Buttons = make(map[*Button]any)
	stage.Buttons_mapString = make(map[string]*Button)

	stage.Nodes = make(map[*Node]any)
	stage.Nodes_mapString = make(map[string]*Node)

	stage.Trees = make(map[*Tree]any)
	stage.Trees_mapString = make(map[string]*Tree)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.Buttons = nil
	stage.Buttons_mapString = nil

	stage.Nodes = nil
	stage.Nodes_mapString = nil

	stage.Trees = nil
	stage.Trees_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for button := range stage.Buttons {
		button.Unstage(stage)
	}

	for node := range stage.Nodes {
		node.Unstage(stage)
	}

	for tree := range stage.Trees {
		tree.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
	// insertion point for generic types
	Button | Node | Tree
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type PointerToGongstruct interface {
	// insertion point for generic types
	*Button | *Node | *Tree
	GetName() string
}

type GongstructSet interface {
	map[any]any |
		// insertion point for generic types
		map[*Button]any |
		map[*Node]any |
		map[*Tree]any |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

type GongstructMapString interface {
	map[any]any |
		// insertion point for generic types
		map[string]*Button |
		map[string]*Node |
		map[string]*Tree |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*Button]any:
		return any(&stage.Buttons).(*Type)
	case map[*Node]any:
		return any(&stage.Nodes).(*Type)
	case map[*Tree]any:
		return any(&stage.Trees).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*Button:
		return any(&stage.Buttons_mapString).(*Type)
	case map[string]*Node:
		return any(&stage.Nodes_mapString).(*Type)
	case map[string]*Tree:
		return any(&stage.Trees_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *StageStruct) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Button:
		return any(&stage.Buttons).(*map[*Type]any)
	case Node:
		return any(&stage.Nodes).(*map[*Type]any)
	case Tree:
		return any(&stage.Trees).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *StageStruct) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Button:
		return any(&stage.Buttons_mapString).(*map[string]*Type)
	case Node:
		return any(&stage.Nodes_mapString).(*map[string]*Type)
	case Tree:
		return any(&stage.Trees_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case Button:
		return any(&Button{
			// Initialisation of associations
		}).(*Type)
	case Node:
		return any(&Node{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			Children: []*Node{{Name: "Children"}},
			// field is initialized with an instance of Button with the name of the field
			Buttons: []*Button{{Name: "Buttons"}},
		}).(*Type)
	case Tree:
		return any(&Tree{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			RootNodes: []*Node{{Name: "RootNodes"}},
		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Button
	case Button:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Node
	case Node:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tree
	case Tree:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Button
	case Button:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Node
	case Node:
		switch fieldname {
		// insertion point for per direct association field
		case "Children":
			res := make(map[*Node]*Node)
			for node := range stage.Nodes {
				for _, node_ := range node.Children {
					res[node_] = node
				}
			}
			return any(res).(map[*End]*Start)
		case "Buttons":
			res := make(map[*Button]*Node)
			for node := range stage.Nodes {
				for _, button_ := range node.Buttons {
					res[button_] = node
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Tree
	case Tree:
		switch fieldname {
		// insertion point for per direct association field
		case "RootNodes":
			res := make(map[*Node]*Tree)
			for tree := range stage.Trees {
				for _, node_ := range tree.RootNodes {
					res[node_] = tree
				}
			}
			return any(res).(map[*End]*Start)
		}
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Button:
		res = "Button"
	case Node:
		res = "Node"
	case Tree:
		res = "Tree"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Button:
		res = []string{"Name", "Icon"}
	case Node:
		res = []string{"Name", "IsExpanded", "HasCheckboxButton", "IsChecked", "IsCheckboxDisabled", "IsInEditMode", "Children", "Buttons"}
	case Tree:
		res = []string{"Name", "RootNodes"}
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct field value
	case Button:
		switch fieldName {
		// string value of fields
		case "Name":
			res = any(instance).(Button).Name
		case "Icon":
			res = any(instance).(Button).Icon
		}
	case Node:
		switch fieldName {
		// string value of fields
		case "Name":
			res = any(instance).(Node).Name
		case "IsExpanded":
			res = fmt.Sprintf("%t", any(instance).(Node).IsExpanded)
		case "HasCheckboxButton":
			res = fmt.Sprintf("%t", any(instance).(Node).HasCheckboxButton)
		case "IsChecked":
			res = fmt.Sprintf("%t", any(instance).(Node).IsChecked)
		case "IsCheckboxDisabled":
			res = fmt.Sprintf("%t", any(instance).(Node).IsCheckboxDisabled)
		case "IsInEditMode":
			res = fmt.Sprintf("%t", any(instance).(Node).IsInEditMode)
		case "Children":
			for idx, __instance__ := range any(instance).(Node).Children {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Buttons":
			for idx, __instance__ := range any(instance).(Node).Buttons {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case Tree:
		switch fieldName {
		// string value of fields
		case "Name":
			res = any(instance).(Tree).Name
		case "RootNodes":
			for idx, __instance__ := range any(instance).(Tree).RootNodes {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	}
	return
}

// Last line of the template
