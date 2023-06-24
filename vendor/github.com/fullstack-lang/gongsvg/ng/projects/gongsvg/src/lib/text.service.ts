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

import { TextDB } from './text-db';

// insertion point for imports
import { LayerDB } from './layer-db'

@Injectable({
  providedIn: 'root'
})
export class TextService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  TextServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private textsUrl: string

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
    this.textsUrl = origin + '/api/github.com/fullstack-lang/gongsvg/go/v1/texts';
  }

  /** GET texts from the server */
  getTexts(GONG__StackPath: string): Observable<TextDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<TextDB[]>(this.textsUrl, { params: params })
      .pipe(
        tap(),
		// tap(_ => this.log('fetched texts')),
        catchError(this.handleError<TextDB[]>('getTexts', []))
      );
  }

  /** GET text by id. Will 404 if id not found */
  getText(id: number, GONG__StackPath: string): Observable<TextDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.textsUrl}/${id}`;
    return this.http.get<TextDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched text id=${id}`)),
      catchError(this.handleError<TextDB>(`getText id=${id}`))
    );
  }

  /** POST: add a new text to the server */
  postText(textdb: TextDB, GONG__StackPath: string): Observable<TextDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    textdb.Animates = []
    let _Layer_Texts_reverse = textdb.Layer_Texts_reverse
    textdb.Layer_Texts_reverse = new LayerDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<TextDB>(this.textsUrl, textdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        textdb.Layer_Texts_reverse = _Layer_Texts_reverse
        // this.log(`posted textdb id=${textdb.ID}`)
      }),
      catchError(this.handleError<TextDB>('postText'))
    );
  }

  /** DELETE: delete the textdb from the server */
  deleteText(textdb: TextDB | number, GONG__StackPath: string): Observable<TextDB> {
    const id = typeof textdb === 'number' ? textdb : textdb.ID;
    const url = `${this.textsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<TextDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted textdb id=${id}`)),
      catchError(this.handleError<TextDB>('deleteText'))
    );
  }

  /** PUT: update the textdb on the server */
  updateText(textdb: TextDB, GONG__StackPath: string): Observable<TextDB> {
    const id = typeof textdb === 'number' ? textdb : textdb.ID;
    const url = `${this.textsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    textdb.Animates = []
    let _Layer_Texts_reverse = textdb.Layer_Texts_reverse
    textdb.Layer_Texts_reverse = new LayerDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<TextDB>(url, textdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        textdb.Layer_Texts_reverse = _Layer_Texts_reverse
        // this.log(`updated textdb id=${textdb.ID}`)
      }),
      catchError(this.handleError<TextDB>('updateText'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in TextService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("TextService" + error); // log to console instead

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
