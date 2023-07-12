import { Component, OnInit } from '@angular/core';
import { Route, Router, Routes } from '@angular/router';

import * as gongrouter from 'gongrouter'
import { ComponentAComponent } from './component-a/component-a.component';
import { ComponentBComponent } from './component-b/component-b.component';
import { ComponentCComponent } from './component-c/component-c.component';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  default = 'Gongrouter Data/Model'
  outlet1 = "Outlet1"
  outlet2 = "Outlet2"
  view = this.outlet1

  views: string[] = [this.outlet1, this.outlet2, this.default];

  DataStack = "gongrouter"
  ModelStacks = "github.com/fullstack-lang/gongrouter/go/models"

  private routes: Routes = []

  constructor(
    private router: Router,
    private routeService: gongrouter.RouteService,
  ) {

  }

  ngOnInit(): void {
    this.addRoutes([
      this.getComponentARoute(this.DataStack),
      this.getComponentBRoute(this.DataStack),
      this.getComponentCRoute(this.DataStack),
    ]


    )
  }

  getComponentARoute(stackPath: string): Route {
    let route: Route =
      { path: "ComponentA", component: ComponentAComponent, outlet: this.outlet1 }
    return route
  }

  getComponentCRoute(stackPath: string): Route {
    let route: Route =
      { path: "ComponentC", component: ComponentCComponent, outlet: this.outlet1 }
    return route
  }

  getComponentBRoute(stackPath: string): Route {
    let route: Route =
      { path: "ComponentB", component: ComponentBComponent, outlet: this.outlet2 }
    return route
  }

  public addRoutes(newRoutes: Routes): void {
    const existingRoutes = this.router.config
    this.routes = this.router.config

    for (let newRoute of newRoutes) {
      if (!existingRoutes.includes(newRoute)) {
        this.routes.push(newRoute)
      }
    }
    this.router.resetConfig(this.routes)
  }
}
