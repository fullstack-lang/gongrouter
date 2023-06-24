// generated by gong
import { Component, OnInit, AfterViewInit, ViewChild, Inject, Optional, Input } from '@angular/core';
import { BehaviorSubject } from 'rxjs'
import { MatSort } from '@angular/material/sort';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData, FrontRepoService, FrontRepo, SelectionMode } from '../front-repo.service'
import { NullInt64 } from '../null-int64'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { ActivatedRoute, Router, RouterState } from '@angular/router';
import { NoteShapeDB } from '../noteshape-db'
import { NoteShapeService } from '../noteshape.service'

// insertion point for additional imports

import { RouteService } from '../route-service';

// TableComponent is initilizaed from different routes
// TableComponentMode detail different cases 
enum TableComponentMode {
  DISPLAY_MODE,
  ONE_MANY_ASSOCIATION_MODE,
  MANY_MANY_ASSOCIATION_MODE,
}

// generated table component
@Component({
  selector: 'app-noteshapestable',
  templateUrl: './noteshapes-table.component.html',
  styleUrls: ['./noteshapes-table.component.css'],
})
export class NoteShapesTableComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of NoteShape instances
  selection: SelectionModel<NoteShapeDB> = new (SelectionModel)
  initialSelection = new Array<NoteShapeDB>()

  // the data source for the table
  noteshapes: NoteShapeDB[] = []
  matTableDataSource: MatTableDataSource<NoteShapeDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.noteshapes
  frontRepo: FrontRepo = new (FrontRepo)

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  // for sorting & pagination
  @ViewChild(MatSort)
  sort: MatSort | undefined
  @ViewChild(MatPaginator)
  paginator: MatPaginator | undefined;

  ngAfterViewInit() {

    // enable sorting on all fields (including pointers and reverse pointer)
    this.matTableDataSource.sortingDataAccessor = (noteshapeDB: NoteShapeDB, property: string) => {
      switch (property) {
        case 'ID':
          return noteshapeDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return noteshapeDB.Name;

        case 'Identifier':
          return noteshapeDB.Identifier;

        case 'Body':
          return noteshapeDB.Body;

        case 'BodyHTML':
          return noteshapeDB.BodyHTML;

        case 'X':
          return noteshapeDB.X;

        case 'Y':
          return noteshapeDB.Y;

        case 'Width':
          return noteshapeDB.Width;

        case 'Heigth':
          return noteshapeDB.Heigth;

        case 'Matched':
          return noteshapeDB.Matched ? "true" : "false";

        case 'Classdiagram_NoteShapes':
          if (this.frontRepo.Classdiagrams.get(noteshapeDB.Classdiagram_NoteShapesDBID.Int64) != undefined) {
            return this.frontRepo.Classdiagrams.get(noteshapeDB.Classdiagram_NoteShapesDBID.Int64)!.Name
          } else {
            return ""
          }

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (noteshapeDB: NoteShapeDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the noteshapeDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += noteshapeDB.Name.toLowerCase()
      mergedContent += noteshapeDB.Identifier.toLowerCase()
      mergedContent += noteshapeDB.Body.toLowerCase()
      mergedContent += noteshapeDB.BodyHTML.toLowerCase()
      mergedContent += noteshapeDB.X.toString()
      mergedContent += noteshapeDB.Y.toString()
      mergedContent += noteshapeDB.Width.toString()
      mergedContent += noteshapeDB.Heigth.toString()
      if (noteshapeDB.Classdiagram_NoteShapesDBID.Int64 != 0) {
        mergedContent += this.frontRepo.Classdiagrams.get(noteshapeDB.Classdiagram_NoteShapesDBID.Int64)!.Name.toLowerCase()
      }


      let isSelected = mergedContent.includes(filter.toLowerCase())
      return isSelected
    };

    this.matTableDataSource.sort = this.sort!
    this.matTableDataSource.paginator = this.paginator!
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.matTableDataSource.filter = filterValue.trim().toLowerCase();
  }

  constructor(
    private noteshapeService: NoteShapeService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of noteshape instances
    public dialogRef: MatDialogRef<NoteShapesTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
    private activatedRoute: ActivatedRoute,

    private routeService: RouteService,
  ) {

    // compute mode
    if (dialogData == undefined) {
      this.mode = TableComponentMode.DISPLAY_MODE
    } else {
      this.GONG__StackPath = dialogData.GONG__StackPath
      switch (dialogData.SelectionMode) {
        case SelectionMode.ONE_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.ONE_MANY_ASSOCIATION_MODE
          break
        case SelectionMode.MANY_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.MANY_MANY_ASSOCIATION_MODE
          break
        default:
      }
    }

    // observable for changes in structs
    this.noteshapeService.NoteShapeServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getNoteShapes()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Delete', // insertion point for columns to display
        "Name",
        "Identifier",
        "Body",
        "BodyHTML",
        "X",
        "Y",
        "Width",
        "Heigth",
        "Matched",
        "Classdiagram_NoteShapes",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Identifier",
        "Body",
        "BodyHTML",
        "X",
        "Y",
        "Width",
        "Heigth",
        "Matched",
        "Classdiagram_NoteShapes",
      ]
      this.selection = new SelectionModel<NoteShapeDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    let stackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')
    if (stackPath != undefined) {
      this.GONG__StackPath = stackPath
    }

    this.getNoteShapes()

    this.matTableDataSource = new MatTableDataSource(this.noteshapes)
  }

  getNoteShapes(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.noteshapes = this.frontRepo.NoteShapes_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let noteshape of this.noteshapes) {
            let ID = this.dialogData.ID
            let revPointer = noteshape[this.dialogData.ReversePointer as keyof NoteShapeDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(noteshape)
            }
            this.selection = new SelectionModel<NoteShapeDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, NoteShapeDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          // we associates on sourceInstance of type SourceStruct with a MANY MANY associations to NoteShapeDB
          // the field name is sourceField
          let sourceFieldArray = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as NoteShapeDB[]
          if (sourceFieldArray != null) {
            for (let associationInstance of sourceFieldArray) {
              let noteshape = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as NoteShapeDB
              this.initialSelection.push(noteshape)
            }
          }

          this.selection = new SelectionModel<NoteShapeDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.noteshapes
      }
    )
  }

  // newNoteShape initiate a new noteshape
  // create a new NoteShape objet
  newNoteShape() {
  }

  deleteNoteShape(noteshapeID: number, noteshape: NoteShapeDB) {
    // list of noteshapes is truncated of noteshape before the delete
    this.noteshapes = this.noteshapes.filter(h => h !== noteshape);

    this.noteshapeService.deleteNoteShape(noteshapeID, this.GONG__StackPath).subscribe(
      noteshape => {
        this.noteshapeService.NoteShapeServiceChanged.next("delete")
      }
    );
  }

  editNoteShape(noteshapeID: number, noteshape: NoteShapeDB) {

  }

  // set editor outlet
  setEditorRouterOutlet(noteshapeID: number) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + "noteshape" + "-detail"

    let outletConf: any = {}
    outletConf[outletName] = [fullPath, noteshapeID, this.GONG__StackPath]

    this.router.navigate([{ outlets: outletConf }])
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.noteshapes.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.noteshapes.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<NoteShapeDB>()

      // reset all initial selection of noteshape that belong to noteshape
      for (let noteshape of this.initialSelection) {
        let index = noteshape[this.dialogData.ReversePointer as keyof NoteShapeDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(noteshape)

      }

      // from selection, set noteshape that belong to noteshape
      for (let noteshape of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = noteshape[this.dialogData.ReversePointer as keyof NoteShapeDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(noteshape)
      }


      // update all noteshape (only update selection & initial selection)
      for (let noteshape of toUpdate) {
        this.noteshapeService.updateNoteShape(noteshape, this.GONG__StackPath)
          .subscribe(noteshape => {
            this.noteshapeService.NoteShapeServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, NoteShapeDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedNoteShape = new Set<number>()
      for (let noteshape of this.initialSelection) {
        if (this.selection.selected.includes(noteshape)) {
          // console.log("noteshape " + noteshape.Name + " is still selected")
        } else {
          console.log("noteshape " + noteshape.Name + " has been unselected")
          unselectedNoteShape.add(noteshape.ID)
          console.log("is unselected " + unselectedNoteShape.has(noteshape.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let noteshape = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as NoteShapeDB
      if (unselectedNoteShape.has(noteshape.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<NoteShapeDB>) = new Array<NoteShapeDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          noteshape => {
            if (!this.initialSelection.includes(noteshape)) {
              // console.log("noteshape " + noteshape.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + noteshape.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = noteshape.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = noteshape.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("noteshape " + noteshape.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<NoteShapeDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
