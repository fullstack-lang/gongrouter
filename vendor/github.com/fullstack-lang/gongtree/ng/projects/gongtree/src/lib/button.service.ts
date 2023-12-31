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

import { ButtonDB } from './button-db';

// insertion point for imports
import { NodeDB } from './node-db'

@Injectable({
  providedIn: 'root'
})
export class ButtonService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  ButtonServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private buttonsUrl: string

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
    this.buttonsUrl = origin + '/api/github.com/fullstack-lang/gongtree/go/v1/buttons';
  }

  /** GET buttons from the server */
  getButtons(GONG__StackPath: string): Observable<ButtonDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<ButtonDB[]>(this.buttonsUrl, { params: params })
      .pipe(
        tap(),
		// tap(_ => this.log('fetched buttons')),
        catchError(this.handleError<ButtonDB[]>('getButtons', []))
      );
  }

  /** GET button by id. Will 404 if id not found */
  getButton(id: number, GONG__StackPath: string): Observable<ButtonDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.buttonsUrl}/${id}`;
    return this.http.get<ButtonDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched button id=${id}`)),
      catchError(this.handleError<ButtonDB>(`getButton id=${id}`))
    );
  }

  /** POST: add a new button to the server */
  postButton(buttondb: ButtonDB, GONG__StackPath: string): Observable<ButtonDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let _Node_Buttons_reverse = buttondb.Node_Buttons_reverse
    buttondb.Node_Buttons_reverse = new NodeDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<ButtonDB>(this.buttonsUrl, buttondb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        buttondb.Node_Buttons_reverse = _Node_Buttons_reverse
        // this.log(`posted buttondb id=${buttondb.ID}`)
      }),
      catchError(this.handleError<ButtonDB>('postButton'))
    );
  }

  /** DELETE: delete the buttondb from the server */
  deleteButton(buttondb: ButtonDB | number, GONG__StackPath: string): Observable<ButtonDB> {
    const id = typeof buttondb === 'number' ? buttondb : buttondb.ID;
    const url = `${this.buttonsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<ButtonDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted buttondb id=${id}`)),
      catchError(this.handleError<ButtonDB>('deleteButton'))
    );
  }

  /** PUT: update the buttondb on the server */
  updateButton(buttondb: ButtonDB, GONG__StackPath: string): Observable<ButtonDB> {
    const id = typeof buttondb === 'number' ? buttondb : buttondb.ID;
    const url = `${this.buttonsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let _Node_Buttons_reverse = buttondb.Node_Buttons_reverse
    buttondb.Node_Buttons_reverse = new NodeDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<ButtonDB>(url, buttondb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        buttondb.Node_Buttons_reverse = _Node_Buttons_reverse
        // this.log(`updated buttondb id=${buttondb.ID}`)
      }),
      catchError(this.handleError<ButtonDB>('updateButton'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in ButtonService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("ButtonService" + error); // log to console instead

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
