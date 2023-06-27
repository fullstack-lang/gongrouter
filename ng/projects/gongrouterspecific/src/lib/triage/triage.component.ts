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
  @Input() StructNames: string[] = []
  tableOutletName: string = ""

  constructor(
    private router: Router,
    private routeService: gongrouter.RouteService,
  ) {

  }

  ngOnInit(): void {
    this.routeService.addDataPanelRoutes(this.DataStack)

    this.tableOutletName = this.routeService.getTableOutlet(this.DataStack)

    this.setTableRouterOutlet(this.StructNames[0].toLowerCase() + "s")
  }

  /**
 * 
 * @param path for the outlet selection
 */
  setTableRouterOutlet(path: string) {
    let outletName = this.routeService.getTableOutlet(this.DataStack)
    let fullPath = this.routeService.getPathRoot() + "-" + path.toLowerCase()
    let outletConf: any = {}
    outletConf[outletName] = [fullPath, this.DataStack]

    this.router.navigate([{ outlets: outletConf }])
  }

}
