import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as gongrouter from 'gongrouter'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  default = 'Gongrouter Data/Model'
  router = "Router"
  view = this.router

  views: string[] = [this.router, this.default];

  DataStack = "gongrouter"
  ModelStacks = "github.com/fullstack-lang/gongrouter/go/models"

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
