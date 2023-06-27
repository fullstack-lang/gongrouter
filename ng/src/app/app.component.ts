import { Component, OnInit } from '@angular/core';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  default = 'Gongrouter Data/Model'
  triage = "Triage"
  view = this.triage

  views: string[] = [this.triage, this.default];

  DataStack = "gongrouter"
  ModelStacks = "github.com/fullstack-lang/gongrouter/go/models"

  myArray = ['Triage']
  tableOutletName: string = ""

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
