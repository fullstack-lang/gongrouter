import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest, BehaviorSubject } from 'rxjs';

// insertion point sub template for services imports 
import { EditorOutletDB } from './editoroutlet-db'
import { EditorOutletService } from './editoroutlet.service'

import { OutletDB } from './outlet-db'
import { OutletService } from './outlet.service'

import { TableOutletDB } from './tableoutlet-db'
import { TableOutletService } from './tableoutlet.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  EditorOutlets_array = new Array<EditorOutletDB>(); // array of repo instances
  EditorOutlets = new Map<number, EditorOutletDB>(); // map of repo instances
  EditorOutlets_batch = new Map<number, EditorOutletDB>(); // same but only in last GET (for finding repo instances to delete)
  Outlets_array = new Array<OutletDB>(); // array of repo instances
  Outlets = new Map<number, OutletDB>(); // map of repo instances
  Outlets_batch = new Map<number, OutletDB>(); // same but only in last GET (for finding repo instances to delete)
  TableOutlets_array = new Array<TableOutletDB>(); // array of repo instances
  TableOutlets = new Map<number, TableOutletDB>(); // map of repo instances
  TableOutlets_batch = new Map<number, TableOutletDB>(); // same but only in last GET (for finding repo instances to delete)
}

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
  ID: number = 0 // ID of the calling instance

  // the reverse pointer is the name of the generated field on the destination
  // struct of the ONE-MANY association
  ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
  OrderingMode: boolean = false // if true, this is for ordering items

  // there are different selection mode : ONE_MANY or MANY_MANY
  SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

  // used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
  //
  // In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
  // 
  // in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
  // at the end of the ONE-MANY association
  SourceStruct: string = ""  // The "Aclass"
  SourceField: string = "" // the "AnarrayofbUse"
  IntermediateStruct: string = "" // the "AclassBclassUse" 
  IntermediateStructField: string = "" // the "Bclass" as field
  NextAssociationStruct: string = "" // the "Bclass"

  GONG__StackPath: string = ""
}

export enum SelectionMode {
  ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
  MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
  providedIn: 'root'
})
export class FrontRepoService {

  GONG__StackPath: string = ""

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  //
  // Store of all instances of the stack
  //
  frontRepo = new (FrontRepo)

  constructor(
    private http: HttpClient, // insertion point sub template 
    private editoroutletService: EditorOutletService,
    private outletService: OutletService,
    private tableoutletService: TableOutletService,
  ) { }

  // postService provides a post function for each struct name
  postService(structName: string, instanceToBePosted: any) {
    let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
    let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

    servicePostFunction(instanceToBePosted).subscribe(
      instance => {
        let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("post")
      }
    );
  }

  // deleteService provides a delete function for each struct name
  deleteService(structName: string, instanceToBeDeleted: any) {
    let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
    let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

    serviceDeleteFunction(instanceToBeDeleted).subscribe(
      instance => {
        let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("delete")
      }
    );
  }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [ // insertion point sub template 
    Observable<EditorOutletDB[]>,
    Observable<OutletDB[]>,
    Observable<TableOutletDB[]>,
  ] = [ // insertion point sub template
      this.editoroutletService.getEditorOutlets(this.GONG__StackPath),
      this.outletService.getOutlets(this.GONG__StackPath),
      this.tableoutletService.getTableOutlets(this.GONG__StackPath),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(GONG__StackPath: string = ""): Observable<FrontRepo> {

    this.GONG__StackPath = GONG__StackPath

    this.observableFrontRepo = [ // insertion point sub template
      this.editoroutletService.getEditorOutlets(this.GONG__StackPath),
      this.outletService.getOutlets(this.GONG__StackPath),
      this.tableoutletService.getTableOutlets(this.GONG__StackPath),
    ]

    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([ // insertion point sub template for declarations 
            editoroutlets_,
            outlets_,
            tableoutlets_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var editoroutlets: EditorOutletDB[]
            editoroutlets = editoroutlets_ as EditorOutletDB[]
            var outlets: OutletDB[]
            outlets = outlets_ as OutletDB[]
            var tableoutlets: TableOutletDB[]
            tableoutlets = tableoutlets_ as TableOutletDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            this.frontRepo.EditorOutlets_array = editoroutlets

            // clear the map that counts EditorOutlet in the GET
            this.frontRepo.EditorOutlets_batch.clear()

            editoroutlets.forEach(
              editoroutlet => {
                this.frontRepo.EditorOutlets.set(editoroutlet.ID, editoroutlet)
                this.frontRepo.EditorOutlets_batch.set(editoroutlet.ID, editoroutlet)
              }
            )

            // clear editoroutlets that are absent from the batch
            this.frontRepo.EditorOutlets.forEach(
              editoroutlet => {
                if (this.frontRepo.EditorOutlets_batch.get(editoroutlet.ID) == undefined) {
                  this.frontRepo.EditorOutlets.delete(editoroutlet.ID)
                }
              }
            )

            // sort EditorOutlets_array array
            this.frontRepo.EditorOutlets_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Outlets_array = outlets

            // clear the map that counts Outlet in the GET
            this.frontRepo.Outlets_batch.clear()

            outlets.forEach(
              outlet => {
                this.frontRepo.Outlets.set(outlet.ID, outlet)
                this.frontRepo.Outlets_batch.set(outlet.ID, outlet)
              }
            )

            // clear outlets that are absent from the batch
            this.frontRepo.Outlets.forEach(
              outlet => {
                if (this.frontRepo.Outlets_batch.get(outlet.ID) == undefined) {
                  this.frontRepo.Outlets.delete(outlet.ID)
                }
              }
            )

            // sort Outlets_array array
            this.frontRepo.Outlets_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.TableOutlets_array = tableoutlets

            // clear the map that counts TableOutlet in the GET
            this.frontRepo.TableOutlets_batch.clear()

            tableoutlets.forEach(
              tableoutlet => {
                this.frontRepo.TableOutlets.set(tableoutlet.ID, tableoutlet)
                this.frontRepo.TableOutlets_batch.set(tableoutlet.ID, tableoutlet)
              }
            )

            // clear tableoutlets that are absent from the batch
            this.frontRepo.TableOutlets.forEach(
              tableoutlet => {
                if (this.frontRepo.TableOutlets_batch.get(tableoutlet.ID) == undefined) {
                  this.frontRepo.TableOutlets.delete(tableoutlet.ID)
                }
              }
            )

            // sort TableOutlets_array array
            this.frontRepo.TableOutlets_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });


            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem 
            editoroutlets.forEach(
              editoroutlet => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            outlets.forEach(
              outlet => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            tableoutlets.forEach(
              tableoutlet => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // EditorOutletPull performs a GET on EditorOutlet of the stack and redeem association pointers 
  EditorOutletPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.editoroutletService.getEditorOutlets(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            editoroutlets,
          ]) => {
            // init the array
            this.frontRepo.EditorOutlets_array = editoroutlets

            // clear the map that counts EditorOutlet in the GET
            this.frontRepo.EditorOutlets_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            editoroutlets.forEach(
              editoroutlet => {
                this.frontRepo.EditorOutlets.set(editoroutlet.ID, editoroutlet)
                this.frontRepo.EditorOutlets_batch.set(editoroutlet.ID, editoroutlet)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear editoroutlets that are absent from the GET
            this.frontRepo.EditorOutlets.forEach(
              editoroutlet => {
                if (this.frontRepo.EditorOutlets_batch.get(editoroutlet.ID) == undefined) {
                  this.frontRepo.EditorOutlets.delete(editoroutlet.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // OutletPull performs a GET on Outlet of the stack and redeem association pointers 
  OutletPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.outletService.getOutlets(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            outlets,
          ]) => {
            // init the array
            this.frontRepo.Outlets_array = outlets

            // clear the map that counts Outlet in the GET
            this.frontRepo.Outlets_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            outlets.forEach(
              outlet => {
                this.frontRepo.Outlets.set(outlet.ID, outlet)
                this.frontRepo.Outlets_batch.set(outlet.ID, outlet)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear outlets that are absent from the GET
            this.frontRepo.Outlets.forEach(
              outlet => {
                if (this.frontRepo.Outlets_batch.get(outlet.ID) == undefined) {
                  this.frontRepo.Outlets.delete(outlet.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // TableOutletPull performs a GET on TableOutlet of the stack and redeem association pointers 
  TableOutletPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.tableoutletService.getTableOutlets(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            tableoutlets,
          ]) => {
            // init the array
            this.frontRepo.TableOutlets_array = tableoutlets

            // clear the map that counts TableOutlet in the GET
            this.frontRepo.TableOutlets_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            tableoutlets.forEach(
              tableoutlet => {
                this.frontRepo.TableOutlets.set(tableoutlet.ID, tableoutlet)
                this.frontRepo.TableOutlets_batch.set(tableoutlet.ID, tableoutlet)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear tableoutlets that are absent from the GET
            this.frontRepo.TableOutlets.forEach(
              tableoutlet => {
                if (this.frontRepo.TableOutlets_batch.get(tableoutlet.ID) == undefined) {
                  this.frontRepo.TableOutlets.delete(tableoutlet.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }
}

// insertion point for get unique ID per struct 
export function getEditorOutletUniqueID(id: number): number {
  return 31 * id
}
export function getOutletUniqueID(id: number): number {
  return 37 * id
}
export function getTableOutletUniqueID(id: number): number {
  return 41 * id
}
