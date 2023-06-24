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

	"github.com/fullstack-lang/gong/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_GongNote_sql sql.NullBool
var dummy_GongNote_time time.Duration
var dummy_GongNote_sort sort.Float64Slice

// GongNoteAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model gongnoteAPI
type GongNoteAPI struct {
	gorm.Model

	models.GongNote

	// encoding of pointers
	GongNotePointersEnconding
}

// GongNotePointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type GongNotePointersEnconding struct {
	// insertion for pointer fields encoding declaration
}

// GongNoteDB describes a gongnote in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model gongnoteDB
type GongNoteDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field gongnoteDB.Name
	Name_Data sql.NullString

	// Declation for basic field gongnoteDB.Body
	Body_Data sql.NullString

	// Declation for basic field gongnoteDB.BodyHTML
	BodyHTML_Data sql.NullString
	// encoding of pointers
	GongNotePointersEnconding
}

// GongNoteDBs arrays gongnoteDBs
// swagger:response gongnoteDBsResponse
type GongNoteDBs []GongNoteDB

// GongNoteDBResponse provides response
// swagger:response gongnoteDBResponse
type GongNoteDBResponse struct {
	GongNoteDB
}

// GongNoteWOP is a GongNote without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type GongNoteWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Body string `xlsx:"2"`

	BodyHTML string `xlsx:"3"`
	// insertion for WOP pointer fields
}

var GongNote_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Body",
	"BodyHTML",
}

type BackRepoGongNoteStruct struct {
	// stores GongNoteDB according to their gorm ID
	Map_GongNoteDBID_GongNoteDB map[uint]*GongNoteDB

	// stores GongNoteDB ID according to GongNote address
	Map_GongNotePtr_GongNoteDBID map[*models.GongNote]uint

	// stores GongNote according to their gorm ID
	Map_GongNoteDBID_GongNotePtr map[uint]*models.GongNote

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoGongNote *BackRepoGongNoteStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoGongNote.stage
	return
}

func (backRepoGongNote *BackRepoGongNoteStruct) GetDB() *gorm.DB {
	return backRepoGongNote.db
}

// GetGongNoteDBFromGongNotePtr is a handy function to access the back repo instance from the stage instance
func (backRepoGongNote *BackRepoGongNoteStruct) GetGongNoteDBFromGongNotePtr(gongnote *models.GongNote) (gongnoteDB *GongNoteDB) {
	id := backRepoGongNote.Map_GongNotePtr_GongNoteDBID[gongnote]
	gongnoteDB = backRepoGongNote.Map_GongNoteDBID_GongNoteDB[id]
	return
}

// BackRepoGongNote.CommitPhaseOne commits all staged instances of GongNote to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongNote *BackRepoGongNoteStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for gongnote := range stage.GongNotes {
		backRepoGongNote.CommitPhaseOneInstance(gongnote)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, gongnote := range backRepoGongNote.Map_GongNoteDBID_GongNotePtr {
		if _, ok := stage.GongNotes[gongnote]; !ok {
			backRepoGongNote.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoGongNote.CommitDeleteInstance commits deletion of GongNote to the BackRepo
func (backRepoGongNote *BackRepoGongNoteStruct) CommitDeleteInstance(id uint) (Error error) {

	gongnote := backRepoGongNote.Map_GongNoteDBID_GongNotePtr[id]

	// gongnote is not staged anymore, remove gongnoteDB
	gongnoteDB := backRepoGongNote.Map_GongNoteDBID_GongNoteDB[id]
	query := backRepoGongNote.db.Unscoped().Delete(&gongnoteDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete(backRepoGongNote.Map_GongNotePtr_GongNoteDBID, gongnote)
	delete(backRepoGongNote.Map_GongNoteDBID_GongNotePtr, id)
	delete(backRepoGongNote.Map_GongNoteDBID_GongNoteDB, id)

	return
}

// BackRepoGongNote.CommitPhaseOneInstance commits gongnote staged instances of GongNote to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongNote *BackRepoGongNoteStruct) CommitPhaseOneInstance(gongnote *models.GongNote) (Error error) {

	// check if the gongnote is not commited yet
	if _, ok := backRepoGongNote.Map_GongNotePtr_GongNoteDBID[gongnote]; ok {
		return
	}

	// initiate gongnote
	var gongnoteDB GongNoteDB
	gongnoteDB.CopyBasicFieldsFromGongNote(gongnote)

	query := backRepoGongNote.db.Create(&gongnoteDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	backRepoGongNote.Map_GongNotePtr_GongNoteDBID[gongnote] = gongnoteDB.ID
	backRepoGongNote.Map_GongNoteDBID_GongNotePtr[gongnoteDB.ID] = gongnote
	backRepoGongNote.Map_GongNoteDBID_GongNoteDB[gongnoteDB.ID] = &gongnoteDB

	return
}

// BackRepoGongNote.CommitPhaseTwo commits all staged instances of GongNote to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongNote *BackRepoGongNoteStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, gongnote := range backRepoGongNote.Map_GongNoteDBID_GongNotePtr {
		backRepoGongNote.CommitPhaseTwoInstance(backRepo, idx, gongnote)
	}

	return
}

// BackRepoGongNote.CommitPhaseTwoInstance commits {{structname }} of models.GongNote to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongNote *BackRepoGongNoteStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, gongnote *models.GongNote) (Error error) {

	// fetch matching gongnoteDB
	if gongnoteDB, ok := backRepoGongNote.Map_GongNoteDBID_GongNoteDB[idx]; ok {

		gongnoteDB.CopyBasicFieldsFromGongNote(gongnote)

		// insertion point for translating pointers encodings into actual pointers
		// This loop encodes the slice of pointers gongnote.Links into the back repo.
		// Each back repo instance at the end of the association encode the ID of the association start
		// into a dedicated field for coding the association. The back repo instance is then saved to the db
		for idx, gonglinkAssocEnd := range gongnote.Links {

			// get the back repo instance at the association end
			gonglinkAssocEnd_DB :=
				backRepo.BackRepoGongLink.GetGongLinkDBFromGongLinkPtr(gonglinkAssocEnd)

			// encode reverse pointer in the association end back repo instance
			gonglinkAssocEnd_DB.GongNote_LinksDBID.Int64 = int64(gongnoteDB.ID)
			gonglinkAssocEnd_DB.GongNote_LinksDBID.Valid = true
			gonglinkAssocEnd_DB.GongNote_LinksDBID_Index.Int64 = int64(idx)
			gonglinkAssocEnd_DB.GongNote_LinksDBID_Index.Valid = true
			if q := backRepoGongNote.db.Save(gonglinkAssocEnd_DB); q.Error != nil {
				return q.Error
			}
		}

		query := backRepoGongNote.db.Save(&gongnoteDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown GongNote intance %s", gongnote.Name))
		return err
	}

	return
}

// BackRepoGongNote.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoGongNote *BackRepoGongNoteStruct) CheckoutPhaseOne() (Error error) {

	gongnoteDBArray := make([]GongNoteDB, 0)
	query := backRepoGongNote.db.Find(&gongnoteDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	gongnoteInstancesToBeRemovedFromTheStage := make(map[*models.GongNote]any)
	for key, value := range backRepoGongNote.stage.GongNotes {
		gongnoteInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, gongnoteDB := range gongnoteDBArray {
		backRepoGongNote.CheckoutPhaseOneInstance(&gongnoteDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		gongnote, ok := backRepoGongNote.Map_GongNoteDBID_GongNotePtr[gongnoteDB.ID]
		if ok {
			delete(gongnoteInstancesToBeRemovedFromTheStage, gongnote)
		}
	}

	// remove from stage and back repo's 3 maps all gongnotes that are not in the checkout
	for gongnote := range gongnoteInstancesToBeRemovedFromTheStage {
		gongnote.Unstage(backRepoGongNote.GetStage())

		// remove instance from the back repo 3 maps
		gongnoteID := backRepoGongNote.Map_GongNotePtr_GongNoteDBID[gongnote]
		delete(backRepoGongNote.Map_GongNotePtr_GongNoteDBID, gongnote)
		delete(backRepoGongNote.Map_GongNoteDBID_GongNoteDB, gongnoteID)
		delete(backRepoGongNote.Map_GongNoteDBID_GongNotePtr, gongnoteID)
	}

	return
}

// CheckoutPhaseOneInstance takes a gongnoteDB that has been found in the DB, updates the backRepo and stages the
// models version of the gongnoteDB
func (backRepoGongNote *BackRepoGongNoteStruct) CheckoutPhaseOneInstance(gongnoteDB *GongNoteDB) (Error error) {

	gongnote, ok := backRepoGongNote.Map_GongNoteDBID_GongNotePtr[gongnoteDB.ID]
	if !ok {
		gongnote = new(models.GongNote)

		backRepoGongNote.Map_GongNoteDBID_GongNotePtr[gongnoteDB.ID] = gongnote
		backRepoGongNote.Map_GongNotePtr_GongNoteDBID[gongnote] = gongnoteDB.ID

		// append model store with the new element
		gongnote.Name = gongnoteDB.Name_Data.String
		gongnote.Stage(backRepoGongNote.GetStage())
	}
	gongnoteDB.CopyBasicFieldsToGongNote(gongnote)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	gongnote.Stage(backRepoGongNote.GetStage())

	// preserve pointer to gongnoteDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_GongNoteDBID_GongNoteDB)[gongnoteDB hold variable pointers
	gongnoteDB_Data := *gongnoteDB
	preservedPtrToGongNote := &gongnoteDB_Data
	backRepoGongNote.Map_GongNoteDBID_GongNoteDB[gongnoteDB.ID] = preservedPtrToGongNote

	return
}

// BackRepoGongNote.CheckoutPhaseTwo Checkouts all staged instances of GongNote to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongNote *BackRepoGongNoteStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, gongnoteDB := range backRepoGongNote.Map_GongNoteDBID_GongNoteDB {
		backRepoGongNote.CheckoutPhaseTwoInstance(backRepo, gongnoteDB)
	}
	return
}

// BackRepoGongNote.CheckoutPhaseTwoInstance Checkouts staged instances of GongNote to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongNote *BackRepoGongNoteStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, gongnoteDB *GongNoteDB) (Error error) {

	gongnote := backRepoGongNote.Map_GongNoteDBID_GongNotePtr[gongnoteDB.ID]
	_ = gongnote // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// This loop redeem gongnote.Links in the stage from the encode in the back repo
	// It parses all GongLinkDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	gongnote.Links = gongnote.Links[:0]
	// 2. loop all instances in the type in the association end
	for _, gonglinkDB_AssocEnd := range backRepo.BackRepoGongLink.Map_GongLinkDBID_GongLinkDB {
		// 3. Does the ID encoding at the end and the ID at the start matches ?
		if gonglinkDB_AssocEnd.GongNote_LinksDBID.Int64 == int64(gongnoteDB.ID) {
			// 4. fetch the associated instance in the stage
			gonglink_AssocEnd := backRepo.BackRepoGongLink.Map_GongLinkDBID_GongLinkPtr[gonglinkDB_AssocEnd.ID]
			// 5. append it the association slice
			gongnote.Links = append(gongnote.Links, gonglink_AssocEnd)
		}
	}

	// sort the array according to the order
	sort.Slice(gongnote.Links, func(i, j int) bool {
		gonglinkDB_i_ID := backRepo.BackRepoGongLink.Map_GongLinkPtr_GongLinkDBID[gongnote.Links[i]]
		gonglinkDB_j_ID := backRepo.BackRepoGongLink.Map_GongLinkPtr_GongLinkDBID[gongnote.Links[j]]

		gonglinkDB_i := backRepo.BackRepoGongLink.Map_GongLinkDBID_GongLinkDB[gonglinkDB_i_ID]
		gonglinkDB_j := backRepo.BackRepoGongLink.Map_GongLinkDBID_GongLinkDB[gonglinkDB_j_ID]

		return gonglinkDB_i.GongNote_LinksDBID_Index.Int64 < gonglinkDB_j.GongNote_LinksDBID_Index.Int64
	})

	return
}

// CommitGongNote allows commit of a single gongnote (if already staged)
func (backRepo *BackRepoStruct) CommitGongNote(gongnote *models.GongNote) {
	backRepo.BackRepoGongNote.CommitPhaseOneInstance(gongnote)
	if id, ok := backRepo.BackRepoGongNote.Map_GongNotePtr_GongNoteDBID[gongnote]; ok {
		backRepo.BackRepoGongNote.CommitPhaseTwoInstance(backRepo, id, gongnote)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitGongNote allows checkout of a single gongnote (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutGongNote(gongnote *models.GongNote) {
	// check if the gongnote is staged
	if _, ok := backRepo.BackRepoGongNote.Map_GongNotePtr_GongNoteDBID[gongnote]; ok {

		if id, ok := backRepo.BackRepoGongNote.Map_GongNotePtr_GongNoteDBID[gongnote]; ok {
			var gongnoteDB GongNoteDB
			gongnoteDB.ID = id

			if err := backRepo.BackRepoGongNote.db.First(&gongnoteDB, id).Error; err != nil {
				log.Panicln("CheckoutGongNote : Problem with getting object with id:", id)
			}
			backRepo.BackRepoGongNote.CheckoutPhaseOneInstance(&gongnoteDB)
			backRepo.BackRepoGongNote.CheckoutPhaseTwoInstance(backRepo, &gongnoteDB)
		}
	}
}

// CopyBasicFieldsFromGongNote
func (gongnoteDB *GongNoteDB) CopyBasicFieldsFromGongNote(gongnote *models.GongNote) {
	// insertion point for fields commit

	gongnoteDB.Name_Data.String = gongnote.Name
	gongnoteDB.Name_Data.Valid = true

	gongnoteDB.Body_Data.String = gongnote.Body
	gongnoteDB.Body_Data.Valid = true

	gongnoteDB.BodyHTML_Data.String = gongnote.BodyHTML
	gongnoteDB.BodyHTML_Data.Valid = true
}

// CopyBasicFieldsFromGongNoteWOP
func (gongnoteDB *GongNoteDB) CopyBasicFieldsFromGongNoteWOP(gongnote *GongNoteWOP) {
	// insertion point for fields commit

	gongnoteDB.Name_Data.String = gongnote.Name
	gongnoteDB.Name_Data.Valid = true

	gongnoteDB.Body_Data.String = gongnote.Body
	gongnoteDB.Body_Data.Valid = true

	gongnoteDB.BodyHTML_Data.String = gongnote.BodyHTML
	gongnoteDB.BodyHTML_Data.Valid = true
}

// CopyBasicFieldsToGongNote
func (gongnoteDB *GongNoteDB) CopyBasicFieldsToGongNote(gongnote *models.GongNote) {
	// insertion point for checkout of basic fields (back repo to stage)
	gongnote.Name = gongnoteDB.Name_Data.String
	gongnote.Body = gongnoteDB.Body_Data.String
	gongnote.BodyHTML = gongnoteDB.BodyHTML_Data.String
}

// CopyBasicFieldsToGongNoteWOP
func (gongnoteDB *GongNoteDB) CopyBasicFieldsToGongNoteWOP(gongnote *GongNoteWOP) {
	gongnote.ID = int(gongnoteDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	gongnote.Name = gongnoteDB.Name_Data.String
	gongnote.Body = gongnoteDB.Body_Data.String
	gongnote.BodyHTML = gongnoteDB.BodyHTML_Data.String
}

// Backup generates a json file from a slice of all GongNoteDB instances in the backrepo
func (backRepoGongNote *BackRepoGongNoteStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "GongNoteDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongNoteDB, 0)
	for _, gongnoteDB := range backRepoGongNote.Map_GongNoteDBID_GongNoteDB {
		forBackup = append(forBackup, gongnoteDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json GongNote ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json GongNote file", err.Error())
	}
}

// Backup generates a json file from a slice of all GongNoteDB instances in the backrepo
func (backRepoGongNote *BackRepoGongNoteStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongNoteDB, 0)
	for _, gongnoteDB := range backRepoGongNote.Map_GongNoteDBID_GongNoteDB {
		forBackup = append(forBackup, gongnoteDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("GongNote")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&GongNote_Fields, -1)
	for _, gongnoteDB := range forBackup {

		var gongnoteWOP GongNoteWOP
		gongnoteDB.CopyBasicFieldsToGongNoteWOP(&gongnoteWOP)

		row := sh.AddRow()
		row.WriteStruct(&gongnoteWOP, -1)
	}
}

// RestoreXL from the "GongNote" sheet all GongNoteDB instances
func (backRepoGongNote *BackRepoGongNoteStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoGongNoteid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["GongNote"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoGongNote.rowVisitorGongNote)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoGongNote *BackRepoGongNoteStruct) rowVisitorGongNote(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var gongnoteWOP GongNoteWOP
		row.ReadStruct(&gongnoteWOP)

		// add the unmarshalled struct to the stage
		gongnoteDB := new(GongNoteDB)
		gongnoteDB.CopyBasicFieldsFromGongNoteWOP(&gongnoteWOP)

		gongnoteDB_ID_atBackupTime := gongnoteDB.ID
		gongnoteDB.ID = 0
		query := backRepoGongNote.db.Create(gongnoteDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoGongNote.Map_GongNoteDBID_GongNoteDB[gongnoteDB.ID] = gongnoteDB
		BackRepoGongNoteid_atBckpTime_newID[gongnoteDB_ID_atBackupTime] = gongnoteDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "GongNoteDB.json" in dirPath that stores an array
// of GongNoteDB and stores it in the database
// the map BackRepoGongNoteid_atBckpTime_newID is updated accordingly
func (backRepoGongNote *BackRepoGongNoteStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoGongNoteid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "GongNoteDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json GongNote file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*GongNoteDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_GongNoteDBID_GongNoteDB
	for _, gongnoteDB := range forRestore {

		gongnoteDB_ID_atBackupTime := gongnoteDB.ID
		gongnoteDB.ID = 0
		query := backRepoGongNote.db.Create(gongnoteDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoGongNote.Map_GongNoteDBID_GongNoteDB[gongnoteDB.ID] = gongnoteDB
		BackRepoGongNoteid_atBckpTime_newID[gongnoteDB_ID_atBackupTime] = gongnoteDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json GongNote file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<GongNote>id_atBckpTime_newID
// to compute new index
func (backRepoGongNote *BackRepoGongNoteStruct) RestorePhaseTwo() {

	for _, gongnoteDB := range backRepoGongNote.Map_GongNoteDBID_GongNoteDB {

		// next line of code is to avert unused variable compilation error
		_ = gongnoteDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoGongNote.db.Model(gongnoteDB).Updates(*gongnoteDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoGongNoteid_atBckpTime_newID map[uint]uint
