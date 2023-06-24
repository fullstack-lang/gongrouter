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

import { PathDB } from './path-db';

// insertion point for imports
import { LayerDB } from './layer-db'

@Injectable({
  providedIn: 'root'
})
export class PathService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  PathServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private pathsUrl: string

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
    this.pathsUrl = origin + '/api/github.com/fullstack-lang/gongsvg/go/v1/paths';
  }

  /** GET paths from the server */
  getPaths(GONG__StackPath: string): Observable<PathDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<PathDB[]>(this.pathsUrl, { params: params })
      .pipe(
        tap(),
		// tap(_ => this.log('fetched paths')),
        catchError(this.handleError<PathDB[]>('getPaths', []))
      );
  }

  /** GET path by id. Will 404 if id not found */
  getPath(id: number, GONG__StackPath: string): Observable<PathDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.pathsUrl}/${id}`;
    return this.http.get<PathDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched path id=${id}`)),
      catchError(this.handleError<PathDB>(`getPath id=${id}`))
    );
  }

  /** POST: add a new path to the server */
  postPath(pathdb: PathDB, GONG__StackPath: string): Observable<PathDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    pathdb.Animates = []
    let _Layer_Paths_reverse = pathdb.Layer_Paths_reverse
    pathdb.Layer_Paths_reverse = new LayerDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<PathDB>(this.pathsUrl, pathdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        pathdb.Layer_Paths_reverse = _Layer_Paths_reverse
        // this.log(`posted pathdb id=${pathdb.ID}`)
      }),
      catchError(this.handleError<PathDB>('postPath'))
    );
  }

  /** DELETE: delete the pathdb from the server */
  deletePath(pathdb: PathDB | number, GONG__StackPath: string): Observable<PathDB> {
    const id = typeof pathdb === 'number' ? pathdb : pathdb.ID;
    const url = `${this.pathsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<PathDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted pathdb id=${id}`)),
      catchError(this.handleError<PathDB>('deletePath'))
    );
  }

  /** PUT: update the pathdb on the server */
  updatePath(pathdb: PathDB, GONG__StackPath: string): Observable<PathDB> {
    const id = typeof pathdb === 'number' ? pathdb : pathdb.ID;
    const url = `${this.pathsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    pathdb.Animates = []
    let _Layer_Paths_reverse = pathdb.Layer_Paths_reverse
    pathdb.Layer_Paths_reverse = new LayerDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<PathDB>(url, pathdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        pathdb.Layer_Paths_reverse = _Layer_Paths_reverse
        // this.log(`updated pathdb id=${pathdb.ID}`)
      }),
      catchError(this.handleError<PathDB>('updatePath'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in PathService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("PathService" + error); // log to console instead

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
