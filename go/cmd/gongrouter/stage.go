package main

import (
	"time"

	"github.com/fullstack-lang/gongrouter/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_stage models.StageStruct
var ___dummy__Time_stage time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_stage map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["stage"] = stageInjection
// }

// stageInjection will stage objects of database "stage"
func stageInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Outlet
	__Outlet__000000_Outlet1 := (&models.Outlet{Name: `Outlet1`}).Stage(stage)
	__Outlet__000001_Outlet2 := (&models.Outlet{Name: `Outlet2`}).Stage(stage)

	// Setup of values

	// Outlet values setup
	__Outlet__000000_Outlet1.Name = `Outlet1`
	__Outlet__000000_Outlet1.Path = `ComponentC`

	// Outlet values setup
	__Outlet__000001_Outlet2.Name = `Outlet2`
	__Outlet__000001_Outlet2.Path = `ComponentB`

	// Setup of pointers
}


