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

import { LinkAnchoredTextDB } from './linkanchoredtext-db';

// insertion point for imports
import { LinkDB } from './link-db'

@Injectable({
  providedIn: 'root'
})
export class LinkAnchoredTextService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  LinkAnchoredTextServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private linkanchoredtextsUrl: string

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
    this.linkanchoredtextsUrl = origin + '/api/github.com/fullstack-lang/gongsvg/go/v1/linkanchoredtexts';
  }

  /** GET linkanchoredtexts from the server */
  getLinkAnchoredTexts(GONG__StackPath: string): Observable<LinkAnchoredTextDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<LinkAnchoredTextDB[]>(this.linkanchoredtextsUrl, { params: params })
      .pipe(
        tap(),
		// tap(_ => this.log('fetched linkanchoredtexts')),
        catchError(this.handleError<LinkAnchoredTextDB[]>('getLinkAnchoredTexts', []))
      );
  }

  /** GET linkanchoredtext by id. Will 404 if id not found */
  getLinkAnchoredText(id: number, GONG__StackPath: string): Observable<LinkAnchoredTextDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.linkanchoredtextsUrl}/${id}`;
    return this.http.get<LinkAnchoredTextDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched linkanchoredtext id=${id}`)),
      catchError(this.handleError<LinkAnchoredTextDB>(`getLinkAnchoredText id=${id}`))
    );
  }

  /** POST: add a new linkanchoredtext to the server */
  postLinkAnchoredText(linkanchoredtextdb: LinkAnchoredTextDB, GONG__StackPath: string): Observable<LinkAnchoredTextDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    linkanchoredtextdb.Animates = []
    let _Link_TextAtArrowEnd_reverse = linkanchoredtextdb.Link_TextAtArrowEnd_reverse
    linkanchoredtextdb.Link_TextAtArrowEnd_reverse = new LinkDB
    let _Link_TextAtArrowStart_reverse = linkanchoredtextdb.Link_TextAtArrowStart_reverse
    linkanchoredtextdb.Link_TextAtArrowStart_reverse = new LinkDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<LinkAnchoredTextDB>(this.linkanchoredtextsUrl, linkanchoredtextdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        linkanchoredtextdb.Link_TextAtArrowEnd_reverse = _Link_TextAtArrowEnd_reverse
        linkanchoredtextdb.Link_TextAtArrowStart_reverse = _Link_TextAtArrowStart_reverse
        // this.log(`posted linkanchoredtextdb id=${linkanchoredtextdb.ID}`)
      }),
      catchError(this.handleError<LinkAnchoredTextDB>('postLinkAnchoredText'))
    );
  }

  /** DELETE: delete the linkanchoredtextdb from the server */
  deleteLinkAnchoredText(linkanchoredtextdb: LinkAnchoredTextDB | number, GONG__StackPath: string): Observable<LinkAnchoredTextDB> {
    const id = typeof linkanchoredtextdb === 'number' ? linkanchoredtextdb : linkanchoredtextdb.ID;
    const url = `${this.linkanchoredtextsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<LinkAnchoredTextDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted linkanchoredtextdb id=${id}`)),
      catchError(this.handleError<LinkAnchoredTextDB>('deleteLinkAnchoredText'))
    );
  }

  /** PUT: update the linkanchoredtextdb on the server */
  updateLinkAnchoredText(linkanchoredtextdb: LinkAnchoredTextDB, GONG__StackPath: string): Observable<LinkAnchoredTextDB> {
    const id = typeof linkanchoredtextdb === 'number' ? linkanchoredtextdb : linkanchoredtextdb.ID;
    const url = `${this.linkanchoredtextsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    linkanchoredtextdb.Animates = []
    let _Link_TextAtArrowEnd_reverse = linkanchoredtextdb.Link_TextAtArrowEnd_reverse
    linkanchoredtextdb.Link_TextAtArrowEnd_reverse = new LinkDB
    let _Link_TextAtArrowStart_reverse = linkanchoredtextdb.Link_TextAtArrowStart_reverse
    linkanchoredtextdb.Link_TextAtArrowStart_reverse = new LinkDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<LinkAnchoredTextDB>(url, linkanchoredtextdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        linkanchoredtextdb.Link_TextAtArrowEnd_reverse = _Link_TextAtArrowEnd_reverse
        linkanchoredtextdb.Link_TextAtArrowStart_reverse = _Link_TextAtArrowStart_reverse
        // this.log(`updated linkanchoredtextdb id=${linkanchoredtextdb.ID}`)
      }),
      catchError(this.handleError<LinkAnchoredTextDB>('updateLinkAnchoredText'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in LinkAnchoredTextService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("LinkAnchoredTextService" + error); // log to console instead

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
