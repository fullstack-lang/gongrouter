// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { ModelPkgDB } from './modelpkg-db';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class ModelPkgService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  ModelPkgServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private modelpkgsUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.modelpkgsUrl = origin + '/api/github.com/fullstack-lang/gong/go/v1/modelpkgs';
  }

  /** GET modelpkgs from the server */
  getModelPkgs(GONG__StackPath: string): Observable<ModelPkgDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<ModelPkgDB[]>(this.modelpkgsUrl, { params: params })
      .pipe(
        tap(),
		// tap(_ => this.log('fetched modelpkgs')),
        catchError(this.handleError<ModelPkgDB[]>('getModelPkgs', []))
      );
  }

  /** GET modelpkg by id. Will 404 if id not found */
  getModelPkg(id: number, GONG__StackPath: string): Observable<ModelPkgDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.modelpkgsUrl}/${id}`;
    return this.http.get<ModelPkgDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched modelpkg id=${id}`)),
      catchError(this.handleError<ModelPkgDB>(`getModelPkg id=${id}`))
    );
  }

  /** POST: add a new modelpkg to the server */
  postModelPkg(modelpkgdb: ModelPkgDB, GONG__StackPath: string): Observable<ModelPkgDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<ModelPkgDB>(this.modelpkgsUrl, modelpkgdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`posted modelpkgdb id=${modelpkgdb.ID}`)
      }),
      catchError(this.handleError<ModelPkgDB>('postModelPkg'))
    );
  }

  /** DELETE: delete the modelpkgdb from the server */
  deleteModelPkg(modelpkgdb: ModelPkgDB | number, GONG__StackPath: string): Observable<ModelPkgDB> {
    const id = typeof modelpkgdb === 'number' ? modelpkgdb : modelpkgdb.ID;
    const url = `${this.modelpkgsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<ModelPkgDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted modelpkgdb id=${id}`)),
      catchError(this.handleError<ModelPkgDB>('deleteModelPkg'))
    );
  }

  /** PUT: update the modelpkgdb on the server */
  updateModelPkg(modelpkgdb: ModelPkgDB, GONG__StackPath: string): Observable<ModelPkgDB> {
    const id = typeof modelpkgdb === 'number' ? modelpkgdb : modelpkgdb.ID;
    const url = `${this.modelpkgsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<ModelPkgDB>(url, modelpkgdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`updated modelpkgdb id=${modelpkgdb.ID}`)
      }),
      catchError(this.handleError<ModelPkgDB>('updateModelPkg'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in ModelPkgService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("ModelPkgService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
      console.log(message)
  }
}
