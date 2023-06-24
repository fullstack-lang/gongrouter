// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/gongdoc/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_GongStructShape_sql sql.NullBool
var dummy_GongStructShape_time time.Duration
var dummy_GongStructShape_sort sort.Float64Slice

// GongStructShapeAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model gongstructshapeAPI
type GongStructShapeAPI struct {
	gorm.Model

	models.GongStructShape

	// encoding of pointers
	GongStructShapePointersEnconding
}

// GongStructShapePointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type GongStructShapePointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// field Position is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	PositionID sql.NullInt64

	// Implementation of a reverse ID for field Classdiagram{}.GongStructShapes []*GongStructShape
	Classdiagram_GongStructShapesDBID sql.NullInt64

	// implementation of the index of the withing the slice
	Classdiagram_GongStructShapesDBID_Index sql.NullInt64
}

// GongStructShapeDB describes a gongstructshape in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model gongstructshapeDB
type GongStructShapeDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field gongstructshapeDB.Name
	Name_Data sql.NullString

	// Declation for basic field gongstructshapeDB.Identifier
	Identifier_Data sql.NullString

	// Declation for basic field gongstructshapeDB.ShowNbInstances
	// provide the sql storage for the boolan
	ShowNbInstances_Data sql.NullBool

	// Declation for basic field gongstructshapeDB.NbInstances
	NbInstances_Data sql.NullInt64

	// Declation for basic field gongstructshapeDB.Width
	Width_Data sql.NullFloat64

	// Declation for basic field gongstructshapeDB.Heigth
	Heigth_Data sql.NullFloat64

	// Declation for basic field gongstructshapeDB.IsSelected
	// provide the sql storage for the boolan
	IsSelected_Data sql.NullBool
	// encoding of pointers
	GongStructShapePointersEnconding
}

// GongStructShapeDBs arrays gongstructshapeDBs
// swagger:response gongstructshapeDBsResponse
type GongStructShapeDBs []GongStructShapeDB

// GongStructShapeDBResponse provides response
// swagger:response gongstructshapeDBResponse
type GongStructShapeDBResponse struct {
	GongStructShapeDB
}

// GongStructShapeWOP is a GongStructShape without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type GongStructShapeWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Identifier string `xlsx:"2"`

	ShowNbInstances bool `xlsx:"3"`

	NbInstances int `xlsx:"4"`

	Width float64 `xlsx:"5"`

	Heigth float64 `xlsx:"6"`

	IsSelected bool `xlsx:"7"`
	// insertion for WOP pointer fields
}

var GongStructShape_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Identifier",
	"ShowNbInstances",
	"NbInstances",
	"Width",
	"Heigth",
	"IsSelected",
}

type BackRepoGongStructShapeStruct struct {
	// stores GongStructShapeDB according to their gorm ID
	Map_GongStructShapeDBID_GongStructShapeDB map[uint]*GongStructShapeDB

	// stores GongStructShapeDB ID according to GongStructShape address
	Map_GongStructShapePtr_GongStructShapeDBID map[*models.GongStructShape]uint

	// stores GongStructShape according to their gorm ID
	Map_GongStructShapeDBID_GongStructShapePtr map[uint]*models.GongStructShape

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoGongStructShape *BackRepoGongStructShapeStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoGongStructShape.stage
	return
}

func (backRepoGongStructShape *BackRepoGongStructShapeStruct) GetDB() *gorm.DB {
	return backRepoGongStructShape.db
}

// GetGongStructShapeDBFromGongStructShapePtr is a handy function to access the back repo instance from the stage instance
func (backRepoGongStructShape *BackRepoGongStructShapeStruct) GetGongStructShapeDBFromGongStructShapePtr(gongstructshape *models.GongStructShape) (gongstructshapeDB *GongStructShapeDB) {
	id := backRepoGongStructShape.Map_GongStructShapePtr_GongStructShapeDBID[gongstructshape]
	gongstructshapeDB = backRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapeDB[id]
	return
}

// BackRepoGongStructShape.CommitPhaseOne commits all staged instances of GongStructShape to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongStructShape *BackRepoGongStructShapeStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for gongstructshape := range stage.GongStructShapes {
		backRepoGongStructShape.CommitPhaseOneInstance(gongstructshape)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, gongstructshape := range backRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapePtr {
		if _, ok := stage.GongStructShapes[gongstructshape]; !ok {
			backRepoGongStructShape.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoGongStructShape.CommitDeleteInstance commits deletion of GongStructShape to the BackRepo
func (backRepoGongStructShape *BackRepoGongStructShapeStruct) CommitDeleteInstance(id uint) (Error error) {

	gongstructshape := backRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapePtr[id]

	// gongstructshape is not staged anymore, remove gongstructshapeDB
	gongstructshapeDB := backRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapeDB[id]
	query := backRepoGongStructShape.db.Unscoped().Delete(&gongstructshapeDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete(backRepoGongStructShape.Map_GongStructShapePtr_GongStructShapeDBID, gongstructshape)
	delete(backRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapePtr, id)
	delete(backRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapeDB, id)

	return
}

// BackRepoGongStructShape.CommitPhaseOneInstance commits gongstructshape staged instances of GongStructShape to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongStructShape *BackRepoGongStructShapeStruct) CommitPhaseOneInstance(gongstructshape *models.GongStructShape) (Error error) {

	// check if the gongstructshape is not commited yet
	if _, ok := backRepoGongStructShape.Map_GongStructShapePtr_GongStructShapeDBID[gongstructshape]; ok {
		return
	}

	// initiate gongstructshape
	var gongstructshapeDB GongStructShapeDB
	gongstructshapeDB.CopyBasicFieldsFromGongStructShape(gongstructshape)

	query := backRepoGongStructShape.db.Create(&gongstructshapeDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	backRepoGongStructShape.Map_GongStructShapePtr_GongStructShapeDBID[gongstructshape] = gongstructshapeDB.ID
	backRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapePtr[gongstructshapeDB.ID] = gongstructshape
	backRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapeDB[gongstructshapeDB.ID] = &gongstructshapeDB

	return
}

// BackRepoGongStructShape.CommitPhaseTwo commits all staged instances of GongStructShape to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongStructShape *BackRepoGongStructShapeStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, gongstructshape := range backRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapePtr {
		backRepoGongStructShape.CommitPhaseTwoInstance(backRepo, idx, gongstructshape)
	}

	return
}

// BackRepoGongStructShape.CommitPhaseTwoInstance commits {{structname }} of models.GongStructShape to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongStructShape *BackRepoGongStructShapeStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, gongstructshape *models.GongStructShape) (Error error) {

	// fetch matching gongstructshapeDB
	if gongstructshapeDB, ok := backRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapeDB[idx]; ok {

		gongstructshapeDB.CopyBasicFieldsFromGongStructShape(gongstructshape)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value gongstructshape.Position translates to updating the gongstructshape.PositionID
		gongstructshapeDB.PositionID.Valid = true // allow for a 0 value (nil association)
		if gongstructshape.Position != nil {
			if PositionId, ok := backRepo.BackRepoPosition.Map_PositionPtr_PositionDBID[gongstructshape.Position]; ok {
				gongstructshapeDB.PositionID.Int64 = int64(PositionId)
				gongstructshapeDB.PositionID.Valid = true
			}
		}

		// This loop encodes the slice of pointers gongstructshape.Fields into the back repo.
		// Each back repo instance at the end of the association encode the ID of the association start
		// into a dedicated field for coding the association. The back repo instance is then saved to the db
		for idx, fieldAssocEnd := range gongstructshape.Fields {

			// get the back repo instance at the association end
			fieldAssocEnd_DB :=
				backRepo.BackRepoField.GetFieldDBFromFieldPtr(fieldAssocEnd)

			// encode reverse pointer in the association end back repo instance
			fieldAssocEnd_DB.GongStructShape_FieldsDBID.Int64 = int64(gongstructshapeDB.ID)
			fieldAssocEnd_DB.GongStructShape_FieldsDBID.Valid = true
			fieldAssocEnd_DB.GongStructShape_FieldsDBID_Index.Int64 = int64(idx)
			fieldAssocEnd_DB.GongStructShape_FieldsDBID_Index.Valid = true
			if q := backRepoGongStructShape.db.Save(fieldAssocEnd_DB); q.Error != nil {
				return q.Error
			}
		}

		// This loop encodes the slice of pointers gongstructshape.Links into the back repo.
		// Each back repo instance at the end of the association encode the ID of the association start
		// into a dedicated field for coding the association. The back repo instance is then saved to the db
		for idx, linkAssocEnd := range gongstructshape.Links {

			// get the back repo instance at the association end
			linkAssocEnd_DB :=
				backRepo.BackRepoLink.GetLinkDBFromLinkPtr(linkAssocEnd)

			// encode reverse pointer in the association end back repo instance
			linkAssocEnd_DB.GongStructShape_LinksDBID.Int64 = int64(gongstructshapeDB.ID)
			linkAssocEnd_DB.GongStructShape_LinksDBID.Valid = true
			linkAssocEnd_DB.GongStructShape_LinksDBID_Index.Int64 = int64(idx)
			linkAssocEnd_DB.GongStructShape_LinksDBID_Index.Valid = true
			if q := backRepoGongStructShape.db.Save(linkAssocEnd_DB); q.Error != nil {
				return q.Error
			}
		}

		query := backRepoGongStructShape.db.Save(&gongstructshapeDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown GongStructShape intance %s", gongstructshape.Name))
		return err
	}

	return
}

// BackRepoGongStructShape.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoGongStructShape *BackRepoGongStructShapeStruct) CheckoutPhaseOne() (Error error) {

	gongstructshapeDBArray := make([]GongStructShapeDB, 0)
	query := backRepoGongStructShape.db.Find(&gongstructshapeDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	gongstructshapeInstancesToBeRemovedFromTheStage := make(map[*models.GongStructShape]any)
	for key, value := range backRepoGongStructShape.stage.GongStructShapes {
		gongstructshapeInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, gongstructshapeDB := range gongstructshapeDBArray {
		backRepoGongStructShape.CheckoutPhaseOneInstance(&gongstructshapeDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		gongstructshape, ok := backRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapePtr[gongstructshapeDB.ID]
		if ok {
			delete(gongstructshapeInstancesToBeRemovedFromTheStage, gongstructshape)
		}
	}

	// remove from stage and back repo's 3 maps all gongstructshapes that are not in the checkout
	for gongstructshape := range gongstructshapeInstancesToBeRemovedFromTheStage {
		gongstructshape.Unstage(backRepoGongStructShape.GetStage())

		// remove instance from the back repo 3 maps
		gongstructshapeID := backRepoGongStructShape.Map_GongStructShapePtr_GongStructShapeDBID[gongstructshape]
		delete(backRepoGongStructShape.Map_GongStructShapePtr_GongStructShapeDBID, gongstructshape)
		delete(backRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapeDB, gongstructshapeID)
		delete(backRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapePtr, gongstructshapeID)
	}

	return
}

// CheckoutPhaseOneInstance takes a gongstructshapeDB that has been found in the DB, updates the backRepo and stages the
// models version of the gongstructshapeDB
func (backRepoGongStructShape *BackRepoGongStructShapeStruct) CheckoutPhaseOneInstance(gongstructshapeDB *GongStructShapeDB) (Error error) {

	gongstructshape, ok := backRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapePtr[gongstructshapeDB.ID]
	if !ok {
		gongstructshape = new(models.GongStructShape)

		backRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapePtr[gongstructshapeDB.ID] = gongstructshape
		backRepoGongStructShape.Map_GongStructShapePtr_GongStructShapeDBID[gongstructshape] = gongstructshapeDB.ID

		// append model store with the new element
		gongstructshape.Name = gongstructshapeDB.Name_Data.String
		gongstructshape.Stage(backRepoGongStructShape.GetStage())
	}
	gongstructshapeDB.CopyBasicFieldsToGongStructShape(gongstructshape)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	gongstructshape.Stage(backRepoGongStructShape.GetStage())

	// preserve pointer to gongstructshapeDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_GongStructShapeDBID_GongStructShapeDB)[gongstructshapeDB hold variable pointers
	gongstructshapeDB_Data := *gongstructshapeDB
	preservedPtrToGongStructShape := &gongstructshapeDB_Data
	backRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapeDB[gongstructshapeDB.ID] = preservedPtrToGongStructShape

	return
}

// BackRepoGongStructShape.CheckoutPhaseTwo Checkouts all staged instances of GongStructShape to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongStructShape *BackRepoGongStructShapeStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, gongstructshapeDB := range backRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapeDB {
		backRepoGongStructShape.CheckoutPhaseTwoInstance(backRepo, gongstructshapeDB)
	}
	return
}

// BackRepoGongStructShape.CheckoutPhaseTwoInstance Checkouts staged instances of GongStructShape to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongStructShape *BackRepoGongStructShapeStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, gongstructshapeDB *GongStructShapeDB) (Error error) {

	gongstructshape := backRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapePtr[gongstructshapeDB.ID]
	_ = gongstructshape // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// Position field
	if gongstructshapeDB.PositionID.Int64 != 0 {
		gongstructshape.Position = backRepo.BackRepoPosition.Map_PositionDBID_PositionPtr[uint(gongstructshapeDB.PositionID.Int64)]
	}
	// This loop redeem gongstructshape.Fields in the stage from the encode in the back repo
	// It parses all FieldDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	gongstructshape.Fields = gongstructshape.Fields[:0]
	// 2. loop all instances in the type in the association end
	for _, fieldDB_AssocEnd := range backRepo.BackRepoField.Map_FieldDBID_FieldDB {
		// 3. Does the ID encoding at the end and the ID at the start matches ?
		if fieldDB_AssocEnd.GongStructShape_FieldsDBID.Int64 == int64(gongstructshapeDB.ID) {
			// 4. fetch the associated instance in the stage
			field_AssocEnd := backRepo.BackRepoField.Map_FieldDBID_FieldPtr[fieldDB_AssocEnd.ID]
			// 5. append it the association slice
			gongstructshape.Fields = append(gongstructshape.Fields, field_AssocEnd)
		}
	}

	// sort the array according to the order
	sort.Slice(gongstructshape.Fields, func(i, j int) bool {
		fieldDB_i_ID := backRepo.BackRepoField.Map_FieldPtr_FieldDBID[gongstructshape.Fields[i]]
		fieldDB_j_ID := backRepo.BackRepoField.Map_FieldPtr_FieldDBID[gongstructshape.Fields[j]]

		fieldDB_i := backRepo.BackRepoField.Map_FieldDBID_FieldDB[fieldDB_i_ID]
		fieldDB_j := backRepo.BackRepoField.Map_FieldDBID_FieldDB[fieldDB_j_ID]

		return fieldDB_i.GongStructShape_FieldsDBID_Index.Int64 < fieldDB_j.GongStructShape_FieldsDBID_Index.Int64
	})

	// This loop redeem gongstructshape.Links in the stage from the encode in the back repo
	// It parses all LinkDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	gongstructshape.Links = gongstructshape.Links[:0]
	// 2. loop all instances in the type in the association end
	for _, linkDB_AssocEnd := range backRepo.BackRepoLink.Map_LinkDBID_LinkDB {
		// 3. Does the ID encoding at the end and the ID at the start matches ?
		if linkDB_AssocEnd.GongStructShape_LinksDBID.Int64 == int64(gongstructshapeDB.ID) {
			// 4. fetch the associated instance in the stage
			link_AssocEnd := backRepo.BackRepoLink.Map_LinkDBID_LinkPtr[linkDB_AssocEnd.ID]
			// 5. append it the association slice
			gongstructshape.Links = append(gongstructshape.Links, link_AssocEnd)
		}
	}

	// sort the array according to the order
	sort.Slice(gongstructshape.Links, func(i, j int) bool {
		linkDB_i_ID := backRepo.BackRepoLink.Map_LinkPtr_LinkDBID[gongstructshape.Links[i]]
		linkDB_j_ID := backRepo.BackRepoLink.Map_LinkPtr_LinkDBID[gongstructshape.Links[j]]

		linkDB_i := backRepo.BackRepoLink.Map_LinkDBID_LinkDB[linkDB_i_ID]
		linkDB_j := backRepo.BackRepoLink.Map_LinkDBID_LinkDB[linkDB_j_ID]

		return linkDB_i.GongStructShape_LinksDBID_Index.Int64 < linkDB_j.GongStructShape_LinksDBID_Index.Int64
	})

	return
}

// CommitGongStructShape allows commit of a single gongstructshape (if already staged)
func (backRepo *BackRepoStruct) CommitGongStructShape(gongstructshape *models.GongStructShape) {
	backRepo.BackRepoGongStructShape.CommitPhaseOneInstance(gongstructshape)
	if id, ok := backRepo.BackRepoGongStructShape.Map_GongStructShapePtr_GongStructShapeDBID[gongstructshape]; ok {
		backRepo.BackRepoGongStructShape.CommitPhaseTwoInstance(backRepo, id, gongstructshape)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitGongStructShape allows checkout of a single gongstructshape (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutGongStructShape(gongstructshape *models.GongStructShape) {
	// check if the gongstructshape is staged
	if _, ok := backRepo.BackRepoGongStructShape.Map_GongStructShapePtr_GongStructShapeDBID[gongstructshape]; ok {

		if id, ok := backRepo.BackRepoGongStructShape.Map_GongStructShapePtr_GongStructShapeDBID[gongstructshape]; ok {
			var gongstructshapeDB GongStructShapeDB
			gongstructshapeDB.ID = id

			if err := backRepo.BackRepoGongStructShape.db.First(&gongstructshapeDB, id).Error; err != nil {
				log.Panicln("CheckoutGongStructShape : Problem with getting object with id:", id)
			}
			backRepo.BackRepoGongStructShape.CheckoutPhaseOneInstance(&gongstructshapeDB)
			backRepo.BackRepoGongStructShape.CheckoutPhaseTwoInstance(backRepo, &gongstructshapeDB)
		}
	}
}

// CopyBasicFieldsFromGongStructShape
func (gongstructshapeDB *GongStructShapeDB) CopyBasicFieldsFromGongStructShape(gongstructshape *models.GongStructShape) {
	// insertion point for fields commit

	gongstructshapeDB.Name_Data.String = gongstructshape.Name
	gongstructshapeDB.Name_Data.Valid = true

	gongstructshapeDB.Identifier_Data.String = gongstructshape.Identifier
	gongstructshapeDB.Identifier_Data.Valid = true

	gongstructshapeDB.ShowNbInstances_Data.Bool = gongstructshape.ShowNbInstances
	gongstructshapeDB.ShowNbInstances_Data.Valid = true

	gongstructshapeDB.NbInstances_Data.Int64 = int64(gongstructshape.NbInstances)
	gongstructshapeDB.NbInstances_Data.Valid = true

	gongstructshapeDB.Width_Data.Float64 = gongstructshape.Width
	gongstructshapeDB.Width_Data.Valid = true

	gongstructshapeDB.Heigth_Data.Float64 = gongstructshape.Heigth
	gongstructshapeDB.Heigth_Data.Valid = true

	gongstructshapeDB.IsSelected_Data.Bool = gongstructshape.IsSelected
	gongstructshapeDB.IsSelected_Data.Valid = true
}

// CopyBasicFieldsFromGongStructShapeWOP
func (gongstructshapeDB *GongStructShapeDB) CopyBasicFieldsFromGongStructShapeWOP(gongstructshape *GongStructShapeWOP) {
	// insertion point for fields commit

	gongstructshapeDB.Name_Data.String = gongstructshape.Name
	gongstructshapeDB.Name_Data.Valid = true

	gongstructshapeDB.Identifier_Data.String = gongstructshape.Identifier
	gongstructshapeDB.Identifier_Data.Valid = true

	gongstructshapeDB.ShowNbInstances_Data.Bool = gongstructshape.ShowNbInstances
	gongstructshapeDB.ShowNbInstances_Data.Valid = true

	gongstructshapeDB.NbInstances_Data.Int64 = int64(gongstructshape.NbInstances)
	gongstructshapeDB.NbInstances_Data.Valid = true

	gongstructshapeDB.Width_Data.Float64 = gongstructshape.Width
	gongstructshapeDB.Width_Data.Valid = true

	gongstructshapeDB.Heigth_Data.Float64 = gongstructshape.Heigth
	gongstructshapeDB.Heigth_Data.Valid = true

	gongstructshapeDB.IsSelected_Data.Bool = gongstructshape.IsSelected
	gongstructshapeDB.IsSelected_Data.Valid = true
}

// CopyBasicFieldsToGongStructShape
func (gongstructshapeDB *GongStructShapeDB) CopyBasicFieldsToGongStructShape(gongstructshape *models.GongStructShape) {
	// insertion point for checkout of basic fields (back repo to stage)
	gongstructshape.Name = gongstructshapeDB.Name_Data.String
	gongstructshape.Identifier = gongstructshapeDB.Identifier_Data.String
	gongstructshape.ShowNbInstances = gongstructshapeDB.ShowNbInstances_Data.Bool
	gongstructshape.NbInstances = int(gongstructshapeDB.NbInstances_Data.Int64)
	gongstructshape.Width = gongstructshapeDB.Width_Data.Float64
	gongstructshape.Heigth = gongstructshapeDB.Heigth_Data.Float64
	gongstructshape.IsSelected = gongstructshapeDB.IsSelected_Data.Bool
}

// CopyBasicFieldsToGongStructShapeWOP
func (gongstructshapeDB *GongStructShapeDB) CopyBasicFieldsToGongStructShapeWOP(gongstructshape *GongStructShapeWOP) {
	gongstructshape.ID = int(gongstructshapeDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	gongstructshape.Name = gongstructshapeDB.Name_Data.String
	gongstructshape.Identifier = gongstructshapeDB.Identifier_Data.String
	gongstructshape.ShowNbInstances = gongstructshapeDB.ShowNbInstances_Data.Bool
	gongstructshape.NbInstances = int(gongstructshapeDB.NbInstances_Data.Int64)
	gongstructshape.Width = gongstructshapeDB.Width_Data.Float64
	gongstructshape.Heigth = gongstructshapeDB.Heigth_Data.Float64
	gongstructshape.IsSelected = gongstructshapeDB.IsSelected_Data.Bool
}

// Backup generates a json file from a slice of all GongStructShapeDB instances in the backrepo
func (backRepoGongStructShape *BackRepoGongStructShapeStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "GongStructShapeDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongStructShapeDB, 0)
	for _, gongstructshapeDB := range backRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapeDB {
		forBackup = append(forBackup, gongstructshapeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json GongStructShape ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json GongStructShape file", err.Error())
	}
}

// Backup generates a json file from a slice of all GongStructShapeDB instances in the backrepo
func (backRepoGongStructShape *BackRepoGongStructShapeStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongStructShapeDB, 0)
	for _, gongstructshapeDB := range backRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapeDB {
		forBackup = append(forBackup, gongstructshapeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("GongStructShape")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&GongStructShape_Fields, -1)
	for _, gongstructshapeDB := range forBackup {

		var gongstructshapeWOP GongStructShapeWOP
		gongstructshapeDB.CopyBasicFieldsToGongStructShapeWOP(&gongstructshapeWOP)

		row := sh.AddRow()
		row.WriteStruct(&gongstructshapeWOP, -1)
	}
}

// RestoreXL from the "GongStructShape" sheet all GongStructShapeDB instances
func (backRepoGongStructShape *BackRepoGongStructShapeStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoGongStructShapeid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["GongStructShape"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoGongStructShape.rowVisitorGongStructShape)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoGongStructShape *BackRepoGongStructShapeStruct) rowVisitorGongStructShape(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var gongstructshapeWOP GongStructShapeWOP
		row.ReadStruct(&gongstructshapeWOP)

		// add the unmarshalled struct to the stage
		gongstructshapeDB := new(GongStructShapeDB)
		gongstructshapeDB.CopyBasicFieldsFromGongStructShapeWOP(&gongstructshapeWOP)

		gongstructshapeDB_ID_atBackupTime := gongstructshapeDB.ID
		gongstructshapeDB.ID = 0
		query := backRepoGongStructShape.db.Create(gongstructshapeDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapeDB[gongstructshapeDB.ID] = gongstructshapeDB
		BackRepoGongStructShapeid_atBckpTime_newID[gongstructshapeDB_ID_atBackupTime] = gongstructshapeDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "GongStructShapeDB.json" in dirPath that stores an array
// of GongStructShapeDB and stores it in the database
// the map BackRepoGongStructShapeid_atBckpTime_newID is updated accordingly
func (backRepoGongStructShape *BackRepoGongStructShapeStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoGongStructShapeid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "GongStructShapeDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json GongStructShape file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*GongStructShapeDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_GongStructShapeDBID_GongStructShapeDB
	for _, gongstructshapeDB := range forRestore {

		gongstructshapeDB_ID_atBackupTime := gongstructshapeDB.ID
		gongstructshapeDB.ID = 0
		query := backRepoGongStructShape.db.Create(gongstructshapeDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapeDB[gongstructshapeDB.ID] = gongstructshapeDB
		BackRepoGongStructShapeid_atBckpTime_newID[gongstructshapeDB_ID_atBackupTime] = gongstructshapeDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json GongStructShape file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<GongStructShape>id_atBckpTime_newID
// to compute new index
func (backRepoGongStructShape *BackRepoGongStructShapeStruct) RestorePhaseTwo() {

	for _, gongstructshapeDB := range backRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapeDB {

		// next line of code is to avert unused variable compilation error
		_ = gongstructshapeDB

		// insertion point for reindexing pointers encoding
		// reindexing Position field
		if gongstructshapeDB.PositionID.Int64 != 0 {
			gongstructshapeDB.PositionID.Int64 = int64(BackRepoPositionid_atBckpTime_newID[uint(gongstructshapeDB.PositionID.Int64)])
			gongstructshapeDB.PositionID.Valid = true
		}

		// This reindex gongstructshape.GongStructShapes
		if gongstructshapeDB.Classdiagram_GongStructShapesDBID.Int64 != 0 {
			gongstructshapeDB.Classdiagram_GongStructShapesDBID.Int64 =
				int64(BackRepoClassdiagramid_atBckpTime_newID[uint(gongstructshapeDB.Classdiagram_GongStructShapesDBID.Int64)])
		}

		// update databse with new index encoding
		query := backRepoGongStructShape.db.Model(gongstructshapeDB).Updates(*gongstructshapeDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoGongStructShapeid_atBckpTime_newID map[uint]uint
