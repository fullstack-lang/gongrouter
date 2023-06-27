import { Component, OnInit } from '@angular/core';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  default = 'Gongrouter Data/Model'
  tableoutlet = "Table Outlet"
  editoroutlet = "Editor Outlet"
  view = this.tableoutlet

  views: string[] = [this.tableoutlet, this.editoroutlet, this.default];

  DataStack = "gongrouter"
  ModelStacks = "github.com/fullstack-lang/gongrouter/go/models"

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
