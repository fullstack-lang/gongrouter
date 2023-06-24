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
import { VerticeDB } from '../vertice-db'
import { VerticeService } from '../vertice.service'

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
  selector: 'app-verticestable',
  templateUrl: './vertices-table.component.html',
  styleUrls: ['./vertices-table.component.css'],
})
export class VerticesTableComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of Vertice instances
  selection: SelectionModel<VerticeDB> = new (SelectionModel)
  initialSelection = new Array<VerticeDB>()

  // the data source for the table
  vertices: VerticeDB[] = []
  matTableDataSource: MatTableDataSource<VerticeDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.vertices
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
    this.matTableDataSource.sortingDataAccessor = (verticeDB: VerticeDB, property: string) => {
      switch (property) {
        case 'ID':
          return verticeDB.ID

        // insertion point for specific sorting accessor
        case 'X':
          return verticeDB.X;

        case 'Y':
          return verticeDB.Y;

        case 'Name':
          return verticeDB.Name;

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (verticeDB: VerticeDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the verticeDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += verticeDB.X.toString()
      mergedContent += verticeDB.Y.toString()
      mergedContent += verticeDB.Name.toLowerCase()

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
    private verticeService: VerticeService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of vertice instances
    public dialogRef: MatDialogRef<VerticesTableComponent>,
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
    this.verticeService.VerticeServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getVertices()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Delete', // insertion point for columns to display
        "X",
        "Y",
        "Name",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "X",
        "Y",
        "Name",
      ]
      this.selection = new SelectionModel<VerticeDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    let stackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')
    if (stackPath != undefined) {
      this.GONG__StackPath = stackPath
    }

    this.getVertices()

    this.matTableDataSource = new MatTableDataSource(this.vertices)
  }

  getVertices(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.vertices = this.frontRepo.Vertices_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let vertice of this.vertices) {
            let ID = this.dialogData.ID
            let revPointer = vertice[this.dialogData.ReversePointer as keyof VerticeDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(vertice)
            }
            this.selection = new SelectionModel<VerticeDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, VerticeDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          // we associates on sourceInstance of type SourceStruct with a MANY MANY associations to VerticeDB
          // the field name is sourceField
          let sourceFieldArray = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as VerticeDB[]
          if (sourceFieldArray != null) {
            for (let associationInstance of sourceFieldArray) {
              let vertice = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as VerticeDB
              this.initialSelection.push(vertice)
            }
          }

          this.selection = new SelectionModel<VerticeDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.vertices
      }
    )
  }

  // newVertice initiate a new vertice
  // create a new Vertice objet
  newVertice() {
  }

  deleteVertice(verticeID: number, vertice: VerticeDB) {
    // list of vertices is truncated of vertice before the delete
    this.vertices = this.vertices.filter(h => h !== vertice);

    this.verticeService.deleteVertice(verticeID, this.GONG__StackPath).subscribe(
      vertice => {
        this.verticeService.VerticeServiceChanged.next("delete")
      }
    );
  }

  editVertice(verticeID: number, vertice: VerticeDB) {

  }

  // set editor outlet
  setEditorRouterOutlet(verticeID: number) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + "vertice" + "-detail"

    let outletConf: any = {}
    outletConf[outletName] = [fullPath, verticeID, this.GONG__StackPath]

    this.router.navigate([{ outlets: outletConf }])
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.vertices.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.vertices.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<VerticeDB>()

      // reset all initial selection of vertice that belong to vertice
      for (let vertice of this.initialSelection) {
        let index = vertice[this.dialogData.ReversePointer as keyof VerticeDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(vertice)

      }

      // from selection, set vertice that belong to vertice
      for (let vertice of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = vertice[this.dialogData.ReversePointer as keyof VerticeDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(vertice)
      }


      // update all vertice (only update selection & initial selection)
      for (let vertice of toUpdate) {
        this.verticeService.updateVertice(vertice, this.GONG__StackPath)
          .subscribe(vertice => {
            this.verticeService.VerticeServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, VerticeDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedVertice = new Set<number>()
      for (let vertice of this.initialSelection) {
        if (this.selection.selected.includes(vertice)) {
          // console.log("vertice " + vertice.Name + " is still selected")
        } else {
          console.log("vertice " + vertice.Name + " has been unselected")
          unselectedVertice.add(vertice.ID)
          console.log("is unselected " + unselectedVertice.has(vertice.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let vertice = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as VerticeDB
      if (unselectedVertice.has(vertice.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<VerticeDB>) = new Array<VerticeDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          vertice => {
            if (!this.initialSelection.includes(vertice)) {
              // console.log("vertice " + vertice.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + vertice.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = vertice.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = vertice.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("vertice " + vertice.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<VerticeDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
