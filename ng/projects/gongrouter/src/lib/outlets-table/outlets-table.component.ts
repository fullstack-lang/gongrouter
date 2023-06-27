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
import { OutletDB } from '../outlet-db'
import { OutletService } from '../outlet.service'

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
  selector: 'app-outletstable',
  templateUrl: './outlets-table.component.html',
  styleUrls: ['./outlets-table.component.css'],
})
export class OutletsTableComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of Outlet instances
  selection: SelectionModel<OutletDB> = new (SelectionModel)
  initialSelection = new Array<OutletDB>()

  // the data source for the table
  outlets: OutletDB[] = []
  matTableDataSource: MatTableDataSource<OutletDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.outlets
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
    this.matTableDataSource.sortingDataAccessor = (outletDB: OutletDB, property: string) => {
      switch (property) {
        case 'ID':
          return outletDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return outletDB.Name;

        case 'Path':
          return outletDB.Path;

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (outletDB: OutletDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the outletDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += outletDB.Name.toLowerCase()
      mergedContent += outletDB.Path.toLowerCase()

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
    private outletService: OutletService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of outlet instances
    public dialogRef: MatDialogRef<OutletsTableComponent>,
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
    this.outletService.OutletServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getOutlets()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Delete', // insertion point for columns to display
        "Name",
        "Path",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Path",
      ]
      this.selection = new SelectionModel<OutletDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    let stackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')
    if (stackPath != undefined) {
      this.GONG__StackPath = stackPath
    }

    this.getOutlets()

    this.matTableDataSource = new MatTableDataSource(this.outlets)
  }

  getOutlets(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.outlets = this.frontRepo.Outlets_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let outlet of this.outlets) {
            let ID = this.dialogData.ID
            let revPointer = outlet[this.dialogData.ReversePointer as keyof OutletDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(outlet)
            }
            this.selection = new SelectionModel<OutletDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, OutletDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          // we associates on sourceInstance of type SourceStruct with a MANY MANY associations to OutletDB
          // the field name is sourceField
          let sourceFieldArray = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as OutletDB[]
          if (sourceFieldArray != null) {
            for (let associationInstance of sourceFieldArray) {
              let outlet = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as OutletDB
              this.initialSelection.push(outlet)
            }
          }

          this.selection = new SelectionModel<OutletDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.outlets
      }
    )
  }

  // newOutlet initiate a new outlet
  // create a new Outlet objet
  newOutlet() {
  }

  deleteOutlet(outletID: number, outlet: OutletDB) {
    // list of outlets is truncated of outlet before the delete
    this.outlets = this.outlets.filter(h => h !== outlet);

    this.outletService.deleteOutlet(outletID, this.GONG__StackPath).subscribe(
      outlet => {
        this.outletService.OutletServiceChanged.next("delete")
      }
    );
  }

  editOutlet(outletID: number, outlet: OutletDB) {

  }

  // set editor outlet
  setEditorRouterOutlet(outletID: number) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + "outlet" + "-detail"

    let outletConf: any = {}
    outletConf[outletName] = [fullPath, outletID, this.GONG__StackPath]

    this.router.navigate([{ outlets: outletConf }])
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.outlets.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.outlets.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<OutletDB>()

      // reset all initial selection of outlet that belong to outlet
      for (let outlet of this.initialSelection) {
        let index = outlet[this.dialogData.ReversePointer as keyof OutletDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(outlet)

      }

      // from selection, set outlet that belong to outlet
      for (let outlet of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = outlet[this.dialogData.ReversePointer as keyof OutletDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(outlet)
      }


      // update all outlet (only update selection & initial selection)
      for (let outlet of toUpdate) {
        this.outletService.updateOutlet(outlet, this.GONG__StackPath)
          .subscribe(outlet => {
            this.outletService.OutletServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, OutletDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedOutlet = new Set<number>()
      for (let outlet of this.initialSelection) {
        if (this.selection.selected.includes(outlet)) {
          // console.log("outlet " + outlet.Name + " is still selected")
        } else {
          console.log("outlet " + outlet.Name + " has been unselected")
          unselectedOutlet.add(outlet.ID)
          console.log("is unselected " + unselectedOutlet.has(outlet.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let outlet = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as OutletDB
      if (unselectedOutlet.has(outlet.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<OutletDB>) = new Array<OutletDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          outlet => {
            if (!this.initialSelection.includes(outlet)) {
              // console.log("outlet " + outlet.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + outlet.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = outlet.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = outlet.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("outlet " + outlet.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<OutletDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}