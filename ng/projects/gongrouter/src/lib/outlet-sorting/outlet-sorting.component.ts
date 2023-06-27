// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { TypeofExpr } from '@angular/compiler';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { OutletDB } from '../outlet-db'
import { OutletService } from '../outlet.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { NullInt64 } from '../null-int64'

@Component({
  selector: 'lib-outlet-sorting',
  templateUrl: './outlet-sorting.component.html',
  styleUrls: ['./outlet-sorting.component.css']
})
export class OutletSortingComponent implements OnInit {

  frontRepo: FrontRepo = new (FrontRepo)

  // array of Outlet instances that are in the association
  associatedOutlets = new Array<OutletDB>();

  constructor(
    private outletService: OutletService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of outlet instances
    public dialogRef: MatDialogRef<OutletSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getOutlets()
  }

  getOutlets(): void {
    this.frontRepoService.pull(this.dialogData.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let outlet of this.frontRepo.Outlets_array) {
          let ID = this.dialogData.ID
          let revPointerID = outlet[this.dialogData.ReversePointer as keyof OutletDB] as unknown as NullInt64
          let revPointerID_Index = outlet[this.dialogData.ReversePointer + "_Index" as keyof OutletDB] as unknown as NullInt64
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedOutlets.push(outlet)
          }
        }

        // sort associated outlet according to order
        this.associatedOutlets.sort((t1, t2) => {
          let t1_revPointerID_Index = t1[this.dialogData.ReversePointer + "_Index" as keyof typeof t1] as unknown as NullInt64
          let t2_revPointerID_Index = t2[this.dialogData.ReversePointer + "_Index" as keyof typeof t2] as unknown as NullInt64
          if (t1_revPointerID_Index && t2_revPointerID_Index) {
            if (t1_revPointerID_Index.Int64 > t2_revPointerID_Index.Int64) {
              return 1;
            }
            if (t1_revPointerID_Index.Int64 < t2_revPointerID_Index.Int64) {
              return -1;
            }
          }
          return 0;
        });
      }
    )
  }

  drop(event: CdkDragDrop<string[]>) {
    moveItemInArray(this.associatedOutlets, event.previousIndex, event.currentIndex);

    // set the order of Outlet instances
    let index = 0

    for (let outlet of this.associatedOutlets) {
      let revPointerID_Index = outlet[this.dialogData.ReversePointer + "_Index" as keyof OutletDB] as unknown as NullInt64
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedOutlets.forEach(
      outlet => {
        this.outletService.updateOutlet(outlet, this.dialogData.GONG__StackPath)
          .subscribe(outlet => {
            this.outletService.OutletServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of ' + this.dialogData.ReversePointer + ' done');
  }
}