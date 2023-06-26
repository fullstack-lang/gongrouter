import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import * as gongrouter from 'gongrouter'

@Component({
  selector: 'lib-triage',
  templateUrl: './triage.component.html',
  styleUrls: ['./triage.component.css']
})
export class TriageComponent implements OnInit {

  @Input() DataStack: string = ""

  constructor(
    private router: Router,
    private routeService: gongrouter.RouteService,
  ) {

  }

  ngOnInit(): void {
    this.routeService.addDataPanelRoutes(this.DataStack)
  }

}
