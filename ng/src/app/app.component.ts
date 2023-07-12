import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import * as gongrouter from 'gongrouter'


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  default = 'Gongrouter Data/Model'
  outletA = "OutletA"
  outletB = "OutletB"
  view = this.outletA

  views: string[] = [this.outletA, this.outletB, this.default];

  DataStack = "gongrouter"
  ModelStacks = "github.com/fullstack-lang/gongrouter/go/models"

  constructor(
    private router: Router,
    private routeService: gongrouter.RouteService,
  ) {

  }

  ngOnInit(): void {
  }
}
