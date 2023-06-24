// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { TypeofExpr } from '@angular/compiler';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { GongEnumShapeDB } from '../gongenumshape-db'
import { GongEnumShapeService } from '../gongenumshape.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { NullInt64 } from '../null-int64'

@Component({
  selector: 'lib-gongenumshape-sorting',
  templateUrl: './gongenumshape-sorting.component.html',
  styleUrls: ['./gongenumshape-sorting.component.css']
})
export class GongEnumShapeSortingComponent implements OnInit {

  frontRepo: FrontRepo = new (FrontRepo)

  // array of GongEnumShape instances that are in the association
  associatedGongEnumShapes = new Array<GongEnumShapeDB>();

  constructor(
    private gongenumshapeService: GongEnumShapeService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of gongenumshape instances
    public dialogRef: MatDialogRef<GongEnumShapeSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getGongEnumShapes()
  }

  getGongEnumShapes(): void {
    this.frontRepoService.pull(this.dialogData.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let gongenumshape of this.frontRepo.GongEnumShapes_array) {
          let ID = this.dialogData.ID
          let revPointerID = gongenumshape[this.dialogData.ReversePointer as keyof GongEnumShapeDB] as unknown as NullInt64
          let revPointerID_Index = gongenumshape[this.dialogData.ReversePointer + "_Index" as keyof GongEnumShapeDB] as unknown as NullInt64
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedGongEnumShapes.push(gongenumshape)
          }
        }

        // sort associated gongenumshape according to order
        this.associatedGongEnumShapes.sort((t1, t2) => {
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
    moveItemInArray(this.associatedGongEnumShapes, event.previousIndex, event.currentIndex);

    // set the order of GongEnumShape instances
    let index = 0

    for (let gongenumshape of this.associatedGongEnumShapes) {
      let revPointerID_Index = gongenumshape[this.dialogData.ReversePointer + "_Index" as keyof GongEnumShapeDB] as unknown as NullInt64
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedGongEnumShapes.forEach(
      gongenumshape => {
        this.gongenumshapeService.updateGongEnumShape(gongenumshape, this.dialogData.GONG__StackPath)
          .subscribe(gongenumshape => {
            this.gongenumshapeService.GongEnumShapeServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of ' + this.dialogData.ReversePointer + ' done');
  }
}
