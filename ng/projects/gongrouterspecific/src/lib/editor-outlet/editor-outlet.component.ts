import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import * as gongrouter from 'gongrouter'
import { Subscription } from 'rxjs';

@Component({
  selector: 'lib-editor-outlet',
  templateUrl: './editor-outlet.component.html',
  styleUrls: ['./editor-outlet.component.css']
})
export class EditorOutletComponent implements OnInit {

  @Input() DataStack: string = ""
  editorOutletName: string = ""

  name: string = ""


  // the component is refreshed when modification are performed in the back repo 
  // 
  // the checkCommitNbFromBackTimer polls the commit number of the back repo
  // if the commit number has increased, it pulls the front repo and redraw the diagram
  private commutNbFromBackSubscription: Subscription = new Subscription
  lastCommitNbFromBack = -1
  lastPushFromFrontNb = -1
  currTime: number = 0
  dateOfLastTimerEmission: Date = new Date


  public gongrouterFrontRepo?: gongrouter.FrontRepo

  constructor(
    private router: Router,
    private routeService: gongrouter.RouteService,
    private gongrouterFrontRepoService: gongrouter.FrontRepoService,
    private gongrouterCommitNbFromBackService: gongrouter.CommitNbFromBackService,
  ) {

  }

  ngOnInit(): void {
    this.routeService.addDataPanelRoutes(this.DataStack)

    this.editorOutletName = this.routeService.getEditorOutlet(this.DataStack)

    this.startAutoRefresh(500); // Refresh every 500 ms (half second)
  }

  ngOnDestroy(): void {
    this.stopAutoRefresh();
  }


  stopAutoRefresh(): void {
    if (this.commutNbFromBackSubscription) {
      this.commutNbFromBackSubscription.unsubscribe();
    }
  }

  startAutoRefresh(intervalMs: number): void {
    this.commutNbFromBackSubscription = this.gongrouterCommitNbFromBackService
      .getCommitNbFromBack(intervalMs, this.DataStack)
      .subscribe((commitNbFromBack: number) => {
        // console.log("EditorOutletComponent, last commit nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack)

        if (this.lastCommitNbFromBack < commitNbFromBack) {
          const d = new Date()
          console.log("EditorOutletComponent, ", this.DataStack, " name ", d.toLocaleTimeString() + `.${d.getMilliseconds()}` +
            ", last commit increased nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack)
          this.lastCommitNbFromBack = commitNbFromBack
          this.refresh()
        }
      }
      )
  }

  refresh(): void {

    this.gongrouterFrontRepoService.pull(this.DataStack).subscribe(
      gongroutersFrontRepo => {
        this.gongrouterFrontRepo = gongroutersFrontRepo

        var editoroutletSingloton: gongrouter.EditorOutletDB = new (gongrouter.EditorOutletDB)
        var selected: boolean = false
        for (var editoroutlet of this.gongrouterFrontRepo.EditorOutlets_array) {
          editoroutletSingloton = editoroutlet
          selected = true
        }
        if (!selected) {
          console.log("no editoroutlet present")
          return
        }

        this.setEditorRouterOutlet(editoroutletSingloton.Name.toLowerCase() + "-adder")

      }
    )
  }

  setEditorRouterOutlet(path: string) {
    let outletName = this.routeService.getEditorOutlet(this.DataStack)
    let fullPath = this.routeService.getPathRoot() + "-" + path.toLowerCase()
    let outletConf: any = {}
    outletConf[outletName] = [fullPath, this.DataStack]
    this.router.navigate([{ outlets: outletConf }]);
  }

}