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

	"github.com/fullstack-lang/gongrouter/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_EditorOutlet_sql sql.NullBool
var dummy_EditorOutlet_time time.Duration
var dummy_EditorOutlet_sort sort.Float64Slice

// EditorOutletAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model editoroutletAPI
type EditorOutletAPI struct {
	gorm.Model

	models.EditorOutlet

	// encoding of pointers
	EditorOutletPointersEnconding
}

// EditorOutletPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type EditorOutletPointersEnconding struct {
	// insertion for pointer fields encoding declaration
}

// EditorOutletDB describes a editoroutlet in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model editoroutletDB
type EditorOutletDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field editoroutletDB.Name
	Name_Data sql.NullString

	// Declation for basic field editoroutletDB.EditorType
	EditorType_Data sql.NullString

	// Declation for basic field editoroutletDB.UpdatedObjectID
	UpdatedObjectID_Data sql.NullInt64
	// encoding of pointers
	EditorOutletPointersEnconding
}

// EditorOutletDBs arrays editoroutletDBs
// swagger:response editoroutletDBsResponse
type EditorOutletDBs []EditorOutletDB

// EditorOutletDBResponse provides response
// swagger:response editoroutletDBResponse
type EditorOutletDBResponse struct {
	EditorOutletDB
}

// EditorOutletWOP is a EditorOutlet without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type EditorOutletWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	EditorType models.EditorType `xlsx:"2"`

	UpdatedObjectID int `xlsx:"3"`
	// insertion for WOP pointer fields
}

var EditorOutlet_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"EditorType",
	"UpdatedObjectID",
}

type BackRepoEditorOutletStruct struct {
	// stores EditorOutletDB according to their gorm ID
	Map_EditorOutletDBID_EditorOutletDB map[uint]*EditorOutletDB

	// stores EditorOutletDB ID according to EditorOutlet address
	Map_EditorOutletPtr_EditorOutletDBID map[*models.EditorOutlet]uint

	// stores EditorOutlet according to their gorm ID
	Map_EditorOutletDBID_EditorOutletPtr map[uint]*models.EditorOutlet

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoEditorOutlet *BackRepoEditorOutletStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoEditorOutlet.stage
	return
}

func (backRepoEditorOutlet *BackRepoEditorOutletStruct) GetDB() *gorm.DB {
	return backRepoEditorOutlet.db
}

// GetEditorOutletDBFromEditorOutletPtr is a handy function to access the back repo instance from the stage instance
func (backRepoEditorOutlet *BackRepoEditorOutletStruct) GetEditorOutletDBFromEditorOutletPtr(editoroutlet *models.EditorOutlet) (editoroutletDB *EditorOutletDB) {
	id := backRepoEditorOutlet.Map_EditorOutletPtr_EditorOutletDBID[editoroutlet]
	editoroutletDB = backRepoEditorOutlet.Map_EditorOutletDBID_EditorOutletDB[id]
	return
}

// BackRepoEditorOutlet.CommitPhaseOne commits all staged instances of EditorOutlet to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoEditorOutlet *BackRepoEditorOutletStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for editoroutlet := range stage.EditorOutlets {
		backRepoEditorOutlet.CommitPhaseOneInstance(editoroutlet)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, editoroutlet := range backRepoEditorOutlet.Map_EditorOutletDBID_EditorOutletPtr {
		if _, ok := stage.EditorOutlets[editoroutlet]; !ok {
			backRepoEditorOutlet.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoEditorOutlet.CommitDeleteInstance commits deletion of EditorOutlet to the BackRepo
func (backRepoEditorOutlet *BackRepoEditorOutletStruct) CommitDeleteInstance(id uint) (Error error) {

	editoroutlet := backRepoEditorOutlet.Map_EditorOutletDBID_EditorOutletPtr[id]

	// editoroutlet is not staged anymore, remove editoroutletDB
	editoroutletDB := backRepoEditorOutlet.Map_EditorOutletDBID_EditorOutletDB[id]
	query := backRepoEditorOutlet.db.Unscoped().Delete(&editoroutletDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete(backRepoEditorOutlet.Map_EditorOutletPtr_EditorOutletDBID, editoroutlet)
	delete(backRepoEditorOutlet.Map_EditorOutletDBID_EditorOutletPtr, id)
	delete(backRepoEditorOutlet.Map_EditorOutletDBID_EditorOutletDB, id)

	return
}

// BackRepoEditorOutlet.CommitPhaseOneInstance commits editoroutlet staged instances of EditorOutlet to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoEditorOutlet *BackRepoEditorOutletStruct) CommitPhaseOneInstance(editoroutlet *models.EditorOutlet) (Error error) {

	// check if the editoroutlet is not commited yet
	if _, ok := backRepoEditorOutlet.Map_EditorOutletPtr_EditorOutletDBID[editoroutlet]; ok {
		return
	}

	// initiate editoroutlet
	var editoroutletDB EditorOutletDB
	editoroutletDB.CopyBasicFieldsFromEditorOutlet(editoroutlet)

	query := backRepoEditorOutlet.db.Create(&editoroutletDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	backRepoEditorOutlet.Map_EditorOutletPtr_EditorOutletDBID[editoroutlet] = editoroutletDB.ID
	backRepoEditorOutlet.Map_EditorOutletDBID_EditorOutletPtr[editoroutletDB.ID] = editoroutlet
	backRepoEditorOutlet.Map_EditorOutletDBID_EditorOutletDB[editoroutletDB.ID] = &editoroutletDB

	return
}

// BackRepoEditorOutlet.CommitPhaseTwo commits all staged instances of EditorOutlet to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoEditorOutlet *BackRepoEditorOutletStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, editoroutlet := range backRepoEditorOutlet.Map_EditorOutletDBID_EditorOutletPtr {
		backRepoEditorOutlet.CommitPhaseTwoInstance(backRepo, idx, editoroutlet)
	}

	return
}

// BackRepoEditorOutlet.CommitPhaseTwoInstance commits {{structname }} of models.EditorOutlet to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoEditorOutlet *BackRepoEditorOutletStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, editoroutlet *models.EditorOutlet) (Error error) {

	// fetch matching editoroutletDB
	if editoroutletDB, ok := backRepoEditorOutlet.Map_EditorOutletDBID_EditorOutletDB[idx]; ok {

		editoroutletDB.CopyBasicFieldsFromEditorOutlet(editoroutlet)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoEditorOutlet.db.Save(&editoroutletDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown EditorOutlet intance %s", editoroutlet.Name))
		return err
	}

	return
}

// BackRepoEditorOutlet.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoEditorOutlet *BackRepoEditorOutletStruct) CheckoutPhaseOne() (Error error) {

	editoroutletDBArray := make([]EditorOutletDB, 0)
	query := backRepoEditorOutlet.db.Find(&editoroutletDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	editoroutletInstancesToBeRemovedFromTheStage := make(map[*models.EditorOutlet]any)
	for key, value := range backRepoEditorOutlet.stage.EditorOutlets {
		editoroutletInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, editoroutletDB := range editoroutletDBArray {
		backRepoEditorOutlet.CheckoutPhaseOneInstance(&editoroutletDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		editoroutlet, ok := backRepoEditorOutlet.Map_EditorOutletDBID_EditorOutletPtr[editoroutletDB.ID]
		if ok {
			delete(editoroutletInstancesToBeRemovedFromTheStage, editoroutlet)
		}
	}

	// remove from stage and back repo's 3 maps all editoroutlets that are not in the checkout
	for editoroutlet := range editoroutletInstancesToBeRemovedFromTheStage {
		editoroutlet.Unstage(backRepoEditorOutlet.GetStage())

		// remove instance from the back repo 3 maps
		editoroutletID := backRepoEditorOutlet.Map_EditorOutletPtr_EditorOutletDBID[editoroutlet]
		delete(backRepoEditorOutlet.Map_EditorOutletPtr_EditorOutletDBID, editoroutlet)
		delete(backRepoEditorOutlet.Map_EditorOutletDBID_EditorOutletDB, editoroutletID)
		delete(backRepoEditorOutlet.Map_EditorOutletDBID_EditorOutletPtr, editoroutletID)
	}

	return
}

// CheckoutPhaseOneInstance takes a editoroutletDB that has been found in the DB, updates the backRepo and stages the
// models version of the editoroutletDB
func (backRepoEditorOutlet *BackRepoEditorOutletStruct) CheckoutPhaseOneInstance(editoroutletDB *EditorOutletDB) (Error error) {

	editoroutlet, ok := backRepoEditorOutlet.Map_EditorOutletDBID_EditorOutletPtr[editoroutletDB.ID]
	if !ok {
		editoroutlet = new(models.EditorOutlet)

		backRepoEditorOutlet.Map_EditorOutletDBID_EditorOutletPtr[editoroutletDB.ID] = editoroutlet
		backRepoEditorOutlet.Map_EditorOutletPtr_EditorOutletDBID[editoroutlet] = editoroutletDB.ID

		// append model store with the new element
		editoroutlet.Name = editoroutletDB.Name_Data.String
		editoroutlet.Stage(backRepoEditorOutlet.GetStage())
	}
	editoroutletDB.CopyBasicFieldsToEditorOutlet(editoroutlet)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	editoroutlet.Stage(backRepoEditorOutlet.GetStage())

	// preserve pointer to editoroutletDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_EditorOutletDBID_EditorOutletDB)[editoroutletDB hold variable pointers
	editoroutletDB_Data := *editoroutletDB
	preservedPtrToEditorOutlet := &editoroutletDB_Data
	backRepoEditorOutlet.Map_EditorOutletDBID_EditorOutletDB[editoroutletDB.ID] = preservedPtrToEditorOutlet

	return
}

// BackRepoEditorOutlet.CheckoutPhaseTwo Checkouts all staged instances of EditorOutlet to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoEditorOutlet *BackRepoEditorOutletStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, editoroutletDB := range backRepoEditorOutlet.Map_EditorOutletDBID_EditorOutletDB {
		backRepoEditorOutlet.CheckoutPhaseTwoInstance(backRepo, editoroutletDB)
	}
	return
}

// BackRepoEditorOutlet.CheckoutPhaseTwoInstance Checkouts staged instances of EditorOutlet to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoEditorOutlet *BackRepoEditorOutletStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, editoroutletDB *EditorOutletDB) (Error error) {

	editoroutlet := backRepoEditorOutlet.Map_EditorOutletDBID_EditorOutletPtr[editoroutletDB.ID]
	_ = editoroutlet // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitEditorOutlet allows commit of a single editoroutlet (if already staged)
func (backRepo *BackRepoStruct) CommitEditorOutlet(editoroutlet *models.EditorOutlet) {
	backRepo.BackRepoEditorOutlet.CommitPhaseOneInstance(editoroutlet)
	if id, ok := backRepo.BackRepoEditorOutlet.Map_EditorOutletPtr_EditorOutletDBID[editoroutlet]; ok {
		backRepo.BackRepoEditorOutlet.CommitPhaseTwoInstance(backRepo, id, editoroutlet)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitEditorOutlet allows checkout of a single editoroutlet (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutEditorOutlet(editoroutlet *models.EditorOutlet) {
	// check if the editoroutlet is staged
	if _, ok := backRepo.BackRepoEditorOutlet.Map_EditorOutletPtr_EditorOutletDBID[editoroutlet]; ok {

		if id, ok := backRepo.BackRepoEditorOutlet.Map_EditorOutletPtr_EditorOutletDBID[editoroutlet]; ok {
			var editoroutletDB EditorOutletDB
			editoroutletDB.ID = id

			if err := backRepo.BackRepoEditorOutlet.db.First(&editoroutletDB, id).Error; err != nil {
				log.Panicln("CheckoutEditorOutlet : Problem with getting object with id:", id)
			}
			backRepo.BackRepoEditorOutlet.CheckoutPhaseOneInstance(&editoroutletDB)
			backRepo.BackRepoEditorOutlet.CheckoutPhaseTwoInstance(backRepo, &editoroutletDB)
		}
	}
}

// CopyBasicFieldsFromEditorOutlet
func (editoroutletDB *EditorOutletDB) CopyBasicFieldsFromEditorOutlet(editoroutlet *models.EditorOutlet) {
	// insertion point for fields commit

	editoroutletDB.Name_Data.String = editoroutlet.Name
	editoroutletDB.Name_Data.Valid = true

	editoroutletDB.EditorType_Data.String = editoroutlet.EditorType.ToString()
	editoroutletDB.EditorType_Data.Valid = true

	editoroutletDB.UpdatedObjectID_Data.Int64 = int64(editoroutlet.UpdatedObjectID)
	editoroutletDB.UpdatedObjectID_Data.Valid = true
}

// CopyBasicFieldsFromEditorOutletWOP
func (editoroutletDB *EditorOutletDB) CopyBasicFieldsFromEditorOutletWOP(editoroutlet *EditorOutletWOP) {
	// insertion point for fields commit

	editoroutletDB.Name_Data.String = editoroutlet.Name
	editoroutletDB.Name_Data.Valid = true

	editoroutletDB.EditorType_Data.String = editoroutlet.EditorType.ToString()
	editoroutletDB.EditorType_Data.Valid = true

	editoroutletDB.UpdatedObjectID_Data.Int64 = int64(editoroutlet.UpdatedObjectID)
	editoroutletDB.UpdatedObjectID_Data.Valid = true
}

// CopyBasicFieldsToEditorOutlet
func (editoroutletDB *EditorOutletDB) CopyBasicFieldsToEditorOutlet(editoroutlet *models.EditorOutlet) {
	// insertion point for checkout of basic fields (back repo to stage)
	editoroutlet.Name = editoroutletDB.Name_Data.String
	editoroutlet.EditorType.FromString(editoroutletDB.EditorType_Data.String)
	editoroutlet.UpdatedObjectID = int(editoroutletDB.UpdatedObjectID_Data.Int64)
}

// CopyBasicFieldsToEditorOutletWOP
func (editoroutletDB *EditorOutletDB) CopyBasicFieldsToEditorOutletWOP(editoroutlet *EditorOutletWOP) {
	editoroutlet.ID = int(editoroutletDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	editoroutlet.Name = editoroutletDB.Name_Data.String
	editoroutlet.EditorType.FromString(editoroutletDB.EditorType_Data.String)
	editoroutlet.UpdatedObjectID = int(editoroutletDB.UpdatedObjectID_Data.Int64)
}

// Backup generates a json file from a slice of all EditorOutletDB instances in the backrepo
func (backRepoEditorOutlet *BackRepoEditorOutletStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "EditorOutletDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*EditorOutletDB, 0)
	for _, editoroutletDB := range backRepoEditorOutlet.Map_EditorOutletDBID_EditorOutletDB {
		forBackup = append(forBackup, editoroutletDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json EditorOutlet ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json EditorOutlet file", err.Error())
	}
}

// Backup generates a json file from a slice of all EditorOutletDB instances in the backrepo
func (backRepoEditorOutlet *BackRepoEditorOutletStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*EditorOutletDB, 0)
	for _, editoroutletDB := range backRepoEditorOutlet.Map_EditorOutletDBID_EditorOutletDB {
		forBackup = append(forBackup, editoroutletDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("EditorOutlet")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&EditorOutlet_Fields, -1)
	for _, editoroutletDB := range forBackup {

		var editoroutletWOP EditorOutletWOP
		editoroutletDB.CopyBasicFieldsToEditorOutletWOP(&editoroutletWOP)

		row := sh.AddRow()
		row.WriteStruct(&editoroutletWOP, -1)
	}
}

// RestoreXL from the "EditorOutlet" sheet all EditorOutletDB instances
func (backRepoEditorOutlet *BackRepoEditorOutletStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoEditorOutletid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["EditorOutlet"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoEditorOutlet.rowVisitorEditorOutlet)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoEditorOutlet *BackRepoEditorOutletStruct) rowVisitorEditorOutlet(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var editoroutletWOP EditorOutletWOP
		row.ReadStruct(&editoroutletWOP)

		// add the unmarshalled struct to the stage
		editoroutletDB := new(EditorOutletDB)
		editoroutletDB.CopyBasicFieldsFromEditorOutletWOP(&editoroutletWOP)

		editoroutletDB_ID_atBackupTime := editoroutletDB.ID
		editoroutletDB.ID = 0
		query := backRepoEditorOutlet.db.Create(editoroutletDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoEditorOutlet.Map_EditorOutletDBID_EditorOutletDB[editoroutletDB.ID] = editoroutletDB
		BackRepoEditorOutletid_atBckpTime_newID[editoroutletDB_ID_atBackupTime] = editoroutletDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "EditorOutletDB.json" in dirPath that stores an array
// of EditorOutletDB and stores it in the database
// the map BackRepoEditorOutletid_atBckpTime_newID is updated accordingly
func (backRepoEditorOutlet *BackRepoEditorOutletStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoEditorOutletid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "EditorOutletDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json EditorOutlet file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*EditorOutletDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_EditorOutletDBID_EditorOutletDB
	for _, editoroutletDB := range forRestore {

		editoroutletDB_ID_atBackupTime := editoroutletDB.ID
		editoroutletDB.ID = 0
		query := backRepoEditorOutlet.db.Create(editoroutletDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoEditorOutlet.Map_EditorOutletDBID_EditorOutletDB[editoroutletDB.ID] = editoroutletDB
		BackRepoEditorOutletid_atBckpTime_newID[editoroutletDB_ID_atBackupTime] = editoroutletDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json EditorOutlet file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<EditorOutlet>id_atBckpTime_newID
// to compute new index
func (backRepoEditorOutlet *BackRepoEditorOutletStruct) RestorePhaseTwo() {

	for _, editoroutletDB := range backRepoEditorOutlet.Map_EditorOutletDBID_EditorOutletDB {

		// next line of code is to avert unused variable compilation error
		_ = editoroutletDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoEditorOutlet.db.Model(editoroutletDB).Updates(*editoroutletDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoEditorOutletid_atBckpTime_newID map[uint]uint
