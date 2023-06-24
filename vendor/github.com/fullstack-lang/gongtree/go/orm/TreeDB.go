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

	"github.com/fullstack-lang/gongtree/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Tree_sql sql.NullBool
var dummy_Tree_time time.Duration
var dummy_Tree_sort sort.Float64Slice

// TreeAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model treeAPI
type TreeAPI struct {
	gorm.Model

	models.Tree

	// encoding of pointers
	TreePointersEnconding
}

// TreePointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type TreePointersEnconding struct {
	// insertion for pointer fields encoding declaration
}

// TreeDB describes a tree in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model treeDB
type TreeDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field treeDB.Name
	Name_Data sql.NullString
	// encoding of pointers
	TreePointersEnconding
}

// TreeDBs arrays treeDBs
// swagger:response treeDBsResponse
type TreeDBs []TreeDB

// TreeDBResponse provides response
// swagger:response treeDBResponse
type TreeDBResponse struct {
	TreeDB
}

// TreeWOP is a Tree without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type TreeWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Tree_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoTreeStruct struct {
	// stores TreeDB according to their gorm ID
	Map_TreeDBID_TreeDB map[uint]*TreeDB

	// stores TreeDB ID according to Tree address
	Map_TreePtr_TreeDBID map[*models.Tree]uint

	// stores Tree according to their gorm ID
	Map_TreeDBID_TreePtr map[uint]*models.Tree

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoTree *BackRepoTreeStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoTree.stage
	return
}

func (backRepoTree *BackRepoTreeStruct) GetDB() *gorm.DB {
	return backRepoTree.db
}

// GetTreeDBFromTreePtr is a handy function to access the back repo instance from the stage instance
func (backRepoTree *BackRepoTreeStruct) GetTreeDBFromTreePtr(tree *models.Tree) (treeDB *TreeDB) {
	id := backRepoTree.Map_TreePtr_TreeDBID[tree]
	treeDB = backRepoTree.Map_TreeDBID_TreeDB[id]
	return
}

// BackRepoTree.CommitPhaseOne commits all staged instances of Tree to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoTree *BackRepoTreeStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for tree := range stage.Trees {
		backRepoTree.CommitPhaseOneInstance(tree)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, tree := range backRepoTree.Map_TreeDBID_TreePtr {
		if _, ok := stage.Trees[tree]; !ok {
			backRepoTree.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoTree.CommitDeleteInstance commits deletion of Tree to the BackRepo
func (backRepoTree *BackRepoTreeStruct) CommitDeleteInstance(id uint) (Error error) {

	tree := backRepoTree.Map_TreeDBID_TreePtr[id]

	// tree is not staged anymore, remove treeDB
	treeDB := backRepoTree.Map_TreeDBID_TreeDB[id]
	query := backRepoTree.db.Unscoped().Delete(&treeDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete(backRepoTree.Map_TreePtr_TreeDBID, tree)
	delete(backRepoTree.Map_TreeDBID_TreePtr, id)
	delete(backRepoTree.Map_TreeDBID_TreeDB, id)

	return
}

// BackRepoTree.CommitPhaseOneInstance commits tree staged instances of Tree to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoTree *BackRepoTreeStruct) CommitPhaseOneInstance(tree *models.Tree) (Error error) {

	// check if the tree is not commited yet
	if _, ok := backRepoTree.Map_TreePtr_TreeDBID[tree]; ok {
		return
	}

	// initiate tree
	var treeDB TreeDB
	treeDB.CopyBasicFieldsFromTree(tree)

	query := backRepoTree.db.Create(&treeDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	backRepoTree.Map_TreePtr_TreeDBID[tree] = treeDB.ID
	backRepoTree.Map_TreeDBID_TreePtr[treeDB.ID] = tree
	backRepoTree.Map_TreeDBID_TreeDB[treeDB.ID] = &treeDB

	return
}

// BackRepoTree.CommitPhaseTwo commits all staged instances of Tree to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoTree *BackRepoTreeStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, tree := range backRepoTree.Map_TreeDBID_TreePtr {
		backRepoTree.CommitPhaseTwoInstance(backRepo, idx, tree)
	}

	return
}

// BackRepoTree.CommitPhaseTwoInstance commits {{structname }} of models.Tree to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoTree *BackRepoTreeStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, tree *models.Tree) (Error error) {

	// fetch matching treeDB
	if treeDB, ok := backRepoTree.Map_TreeDBID_TreeDB[idx]; ok {

		treeDB.CopyBasicFieldsFromTree(tree)

		// insertion point for translating pointers encodings into actual pointers
		// This loop encodes the slice of pointers tree.RootNodes into the back repo.
		// Each back repo instance at the end of the association encode the ID of the association start
		// into a dedicated field for coding the association. The back repo instance is then saved to the db
		for idx, nodeAssocEnd := range tree.RootNodes {

			// get the back repo instance at the association end
			nodeAssocEnd_DB :=
				backRepo.BackRepoNode.GetNodeDBFromNodePtr(nodeAssocEnd)

			// encode reverse pointer in the association end back repo instance
			nodeAssocEnd_DB.Tree_RootNodesDBID.Int64 = int64(treeDB.ID)
			nodeAssocEnd_DB.Tree_RootNodesDBID.Valid = true
			nodeAssocEnd_DB.Tree_RootNodesDBID_Index.Int64 = int64(idx)
			nodeAssocEnd_DB.Tree_RootNodesDBID_Index.Valid = true
			if q := backRepoTree.db.Save(nodeAssocEnd_DB); q.Error != nil {
				return q.Error
			}
		}

		query := backRepoTree.db.Save(&treeDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Tree intance %s", tree.Name))
		return err
	}

	return
}

// BackRepoTree.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoTree *BackRepoTreeStruct) CheckoutPhaseOne() (Error error) {

	treeDBArray := make([]TreeDB, 0)
	query := backRepoTree.db.Find(&treeDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	treeInstancesToBeRemovedFromTheStage := make(map[*models.Tree]any)
	for key, value := range backRepoTree.stage.Trees {
		treeInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, treeDB := range treeDBArray {
		backRepoTree.CheckoutPhaseOneInstance(&treeDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		tree, ok := backRepoTree.Map_TreeDBID_TreePtr[treeDB.ID]
		if ok {
			delete(treeInstancesToBeRemovedFromTheStage, tree)
		}
	}

	// remove from stage and back repo's 3 maps all trees that are not in the checkout
	for tree := range treeInstancesToBeRemovedFromTheStage {
		tree.Unstage(backRepoTree.GetStage())

		// remove instance from the back repo 3 maps
		treeID := backRepoTree.Map_TreePtr_TreeDBID[tree]
		delete(backRepoTree.Map_TreePtr_TreeDBID, tree)
		delete(backRepoTree.Map_TreeDBID_TreeDB, treeID)
		delete(backRepoTree.Map_TreeDBID_TreePtr, treeID)
	}

	return
}

// CheckoutPhaseOneInstance takes a treeDB that has been found in the DB, updates the backRepo and stages the
// models version of the treeDB
func (backRepoTree *BackRepoTreeStruct) CheckoutPhaseOneInstance(treeDB *TreeDB) (Error error) {

	tree, ok := backRepoTree.Map_TreeDBID_TreePtr[treeDB.ID]
	if !ok {
		tree = new(models.Tree)

		backRepoTree.Map_TreeDBID_TreePtr[treeDB.ID] = tree
		backRepoTree.Map_TreePtr_TreeDBID[tree] = treeDB.ID

		// append model store with the new element
		tree.Name = treeDB.Name_Data.String
		tree.Stage(backRepoTree.GetStage())
	}
	treeDB.CopyBasicFieldsToTree(tree)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	tree.Stage(backRepoTree.GetStage())

	// preserve pointer to treeDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_TreeDBID_TreeDB)[treeDB hold variable pointers
	treeDB_Data := *treeDB
	preservedPtrToTree := &treeDB_Data
	backRepoTree.Map_TreeDBID_TreeDB[treeDB.ID] = preservedPtrToTree

	return
}

// BackRepoTree.CheckoutPhaseTwo Checkouts all staged instances of Tree to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoTree *BackRepoTreeStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, treeDB := range backRepoTree.Map_TreeDBID_TreeDB {
		backRepoTree.CheckoutPhaseTwoInstance(backRepo, treeDB)
	}
	return
}

// BackRepoTree.CheckoutPhaseTwoInstance Checkouts staged instances of Tree to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoTree *BackRepoTreeStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, treeDB *TreeDB) (Error error) {

	tree := backRepoTree.Map_TreeDBID_TreePtr[treeDB.ID]
	_ = tree // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// This loop redeem tree.RootNodes in the stage from the encode in the back repo
	// It parses all NodeDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	tree.RootNodes = tree.RootNodes[:0]
	// 2. loop all instances in the type in the association end
	for _, nodeDB_AssocEnd := range backRepo.BackRepoNode.Map_NodeDBID_NodeDB {
		// 3. Does the ID encoding at the end and the ID at the start matches ?
		if nodeDB_AssocEnd.Tree_RootNodesDBID.Int64 == int64(treeDB.ID) {
			// 4. fetch the associated instance in the stage
			node_AssocEnd := backRepo.BackRepoNode.Map_NodeDBID_NodePtr[nodeDB_AssocEnd.ID]
			// 5. append it the association slice
			tree.RootNodes = append(tree.RootNodes, node_AssocEnd)
		}
	}

	// sort the array according to the order
	sort.Slice(tree.RootNodes, func(i, j int) bool {
		nodeDB_i_ID := backRepo.BackRepoNode.Map_NodePtr_NodeDBID[tree.RootNodes[i]]
		nodeDB_j_ID := backRepo.BackRepoNode.Map_NodePtr_NodeDBID[tree.RootNodes[j]]

		nodeDB_i := backRepo.BackRepoNode.Map_NodeDBID_NodeDB[nodeDB_i_ID]
		nodeDB_j := backRepo.BackRepoNode.Map_NodeDBID_NodeDB[nodeDB_j_ID]

		return nodeDB_i.Tree_RootNodesDBID_Index.Int64 < nodeDB_j.Tree_RootNodesDBID_Index.Int64
	})

	return
}

// CommitTree allows commit of a single tree (if already staged)
func (backRepo *BackRepoStruct) CommitTree(tree *models.Tree) {
	backRepo.BackRepoTree.CommitPhaseOneInstance(tree)
	if id, ok := backRepo.BackRepoTree.Map_TreePtr_TreeDBID[tree]; ok {
		backRepo.BackRepoTree.CommitPhaseTwoInstance(backRepo, id, tree)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitTree allows checkout of a single tree (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutTree(tree *models.Tree) {
	// check if the tree is staged
	if _, ok := backRepo.BackRepoTree.Map_TreePtr_TreeDBID[tree]; ok {

		if id, ok := backRepo.BackRepoTree.Map_TreePtr_TreeDBID[tree]; ok {
			var treeDB TreeDB
			treeDB.ID = id

			if err := backRepo.BackRepoTree.db.First(&treeDB, id).Error; err != nil {
				log.Panicln("CheckoutTree : Problem with getting object with id:", id)
			}
			backRepo.BackRepoTree.CheckoutPhaseOneInstance(&treeDB)
			backRepo.BackRepoTree.CheckoutPhaseTwoInstance(backRepo, &treeDB)
		}
	}
}

// CopyBasicFieldsFromTree
func (treeDB *TreeDB) CopyBasicFieldsFromTree(tree *models.Tree) {
	// insertion point for fields commit

	treeDB.Name_Data.String = tree.Name
	treeDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromTreeWOP
func (treeDB *TreeDB) CopyBasicFieldsFromTreeWOP(tree *TreeWOP) {
	// insertion point for fields commit

	treeDB.Name_Data.String = tree.Name
	treeDB.Name_Data.Valid = true
}

// CopyBasicFieldsToTree
func (treeDB *TreeDB) CopyBasicFieldsToTree(tree *models.Tree) {
	// insertion point for checkout of basic fields (back repo to stage)
	tree.Name = treeDB.Name_Data.String
}

// CopyBasicFieldsToTreeWOP
func (treeDB *TreeDB) CopyBasicFieldsToTreeWOP(tree *TreeWOP) {
	tree.ID = int(treeDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	tree.Name = treeDB.Name_Data.String
}

// Backup generates a json file from a slice of all TreeDB instances in the backrepo
func (backRepoTree *BackRepoTreeStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "TreeDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*TreeDB, 0)
	for _, treeDB := range backRepoTree.Map_TreeDBID_TreeDB {
		forBackup = append(forBackup, treeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json Tree ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json Tree file", err.Error())
	}
}

// Backup generates a json file from a slice of all TreeDB instances in the backrepo
func (backRepoTree *BackRepoTreeStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*TreeDB, 0)
	for _, treeDB := range backRepoTree.Map_TreeDBID_TreeDB {
		forBackup = append(forBackup, treeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Tree")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Tree_Fields, -1)
	for _, treeDB := range forBackup {

		var treeWOP TreeWOP
		treeDB.CopyBasicFieldsToTreeWOP(&treeWOP)

		row := sh.AddRow()
		row.WriteStruct(&treeWOP, -1)
	}
}

// RestoreXL from the "Tree" sheet all TreeDB instances
func (backRepoTree *BackRepoTreeStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoTreeid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Tree"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoTree.rowVisitorTree)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoTree *BackRepoTreeStruct) rowVisitorTree(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var treeWOP TreeWOP
		row.ReadStruct(&treeWOP)

		// add the unmarshalled struct to the stage
		treeDB := new(TreeDB)
		treeDB.CopyBasicFieldsFromTreeWOP(&treeWOP)

		treeDB_ID_atBackupTime := treeDB.ID
		treeDB.ID = 0
		query := backRepoTree.db.Create(treeDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoTree.Map_TreeDBID_TreeDB[treeDB.ID] = treeDB
		BackRepoTreeid_atBckpTime_newID[treeDB_ID_atBackupTime] = treeDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "TreeDB.json" in dirPath that stores an array
// of TreeDB and stores it in the database
// the map BackRepoTreeid_atBckpTime_newID is updated accordingly
func (backRepoTree *BackRepoTreeStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoTreeid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "TreeDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json Tree file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*TreeDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_TreeDBID_TreeDB
	for _, treeDB := range forRestore {

		treeDB_ID_atBackupTime := treeDB.ID
		treeDB.ID = 0
		query := backRepoTree.db.Create(treeDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoTree.Map_TreeDBID_TreeDB[treeDB.ID] = treeDB
		BackRepoTreeid_atBckpTime_newID[treeDB_ID_atBackupTime] = treeDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json Tree file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Tree>id_atBckpTime_newID
// to compute new index
func (backRepoTree *BackRepoTreeStruct) RestorePhaseTwo() {

	for _, treeDB := range backRepoTree.Map_TreeDBID_TreeDB {

		// next line of code is to avert unused variable compilation error
		_ = treeDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoTree.db.Model(treeDB).Updates(*treeDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoTreeid_atBckpTime_newID map[uint]uint
