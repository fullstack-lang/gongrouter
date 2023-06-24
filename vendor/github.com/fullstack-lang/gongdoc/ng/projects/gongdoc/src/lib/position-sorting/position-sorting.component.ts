// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { TypeofExpr } from '@angular/compiler';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { PositionDB } from '../position-db'
import { PositionService } from '../position.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { NullInt64 } from '../null-int64'

@Component({
  selector: 'lib-position-sorting',
  templateUrl: './position-sorting.component.html',
  styleUrls: ['./position-sorting.component.css']
})
export class PositionSortingComponent implements OnInit {

  frontRepo: FrontRepo = new (FrontRepo)

  // array of Position instances that are in the association
  associatedPositions = new Array<PositionDB>();

  constructor(
    private positionService: PositionService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of position instances
    public dialogRef: MatDialogRef<PositionSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getPositions()
  }

  getPositions(): void {
    this.frontRepoService.pull(this.dialogData.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let position of this.frontRepo.Positions_array) {
          let ID = this.dialogData.ID
          let revPointerID = position[this.dialogData.ReversePointer as keyof PositionDB] as unknown as NullInt64
          let revPointerID_Index = position[this.dialogData.ReversePointer + "_Index" as keyof PositionDB] as unknown as NullInt64
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedPositions.push(position)
          }
        }

        // sort associated position according to order
        this.associatedPositions.sort((t1, t2) => {
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
    moveItemInArray(this.associatedPositions, event.previousIndex, event.currentIndex);

    // set the order of Position instances
    let index = 0

    for (let position of this.associatedPositions) {
      let revPointerID_Index = position[this.dialogData.ReversePointer + "_Index" as keyof PositionDB] as unknown as NullInt64
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedPositions.forEach(
      position => {
        this.positionService.updatePosition(position, this.dialogData.GONG__StackPath)
          .subscribe(position => {
            this.positionService.PositionServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of ' + this.dialogData.ReversePointer + ' done');
  }
}
