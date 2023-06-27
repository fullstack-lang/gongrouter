import { Injectable } from '@angular/core';
import { Route, Router, Routes } from '@angular/router';

// insertion point for imports
import { TableOutletsTableComponent } from './tableoutlets-table/tableoutlets-table.component'
import { TableOutletDetailComponent } from './tableoutlet-detail/tableoutlet-detail.component'


@Injectable({
    providedIn: 'root'
})
export class RouteService {
    private routes: Routes = []

    constructor(private router: Router) { }

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

    getPathRoot(): string {
        return 'github_com_fullstack_lang_gongrouter_go'
    }
    getTableOutlet(stackPath: string): string {
        return this.getPathRoot() + '_table' + '_' + stackPath
    }
    getEditorOutlet(stackPath: string): string {
        return this.getPathRoot() + '_editor' + '_' + stackPath
    }
    // insertion point for per gongstruct route/path getters
    getTableOutletTablePath(): string {
        return this.getPathRoot() + '-tableoutlets/:GONG__StackPath'
    }
    getTableOutletTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableOutletTablePath(), component: TableOutletsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getTableOutletAdderPath(): string {
        return this.getPathRoot() + '-tableoutlet-adder/:GONG__StackPath'
    }
    getTableOutletAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableOutletAdderPath(), component: TableOutletDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getTableOutletAdderForUsePath(): string {
        return this.getPathRoot() + '-tableoutlet-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getTableOutletAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableOutletAdderForUsePath(), component: TableOutletDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getTableOutletDetailPath(): string {
        return this.getPathRoot() + '-tableoutlet-detail/:id/:GONG__StackPath'
    }
    getTableOutletDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableOutletDetailPath(), component: TableOutletDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }



    addDataPanelRoutes(stackPath: string) {

        this.addRoutes([
            // insertion point for all routes getter
            this.getTableOutletTableRoute(stackPath),
            this.getTableOutletAdderRoute(stackPath),
            this.getTableOutletAdderForUseRoute(stackPath),
            this.getTableOutletDetailRoute(stackPath),

        ])
    }
}
