import { Injectable } from '@angular/core';
import { Route, Router, Routes } from '@angular/router';

// insertion point for imports
import { TriagesTableComponent } from './triages-table/triages-table.component'
import { TriageDetailComponent } from './triage-detail/triage-detail.component'


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
    getTriageTablePath(): string {
        return this.getPathRoot() + '-triages/:GONG__StackPath'
    }
    getTriageTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTriageTablePath(), component: TriagesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getTriageAdderPath(): string {
        return this.getPathRoot() + '-triage-adder/:GONG__StackPath'
    }
    getTriageAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTriageAdderPath(), component: TriageDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getTriageAdderForUsePath(): string {
        return this.getPathRoot() + '-triage-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getTriageAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTriageAdderForUsePath(), component: TriageDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getTriageDetailPath(): string {
        return this.getPathRoot() + '-triage-detail/:id/:GONG__StackPath'
    }
    getTriageDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTriageDetailPath(), component: TriageDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }



    addDataPanelRoutes(stackPath: string) {

        this.addRoutes([
            // insertion point for all routes getter
            this.getTriageTableRoute(stackPath),
            this.getTriageAdderRoute(stackPath),
            this.getTriageAdderForUseRoute(stackPath),
            this.getTriageDetailRoute(stackPath),

        ])
    }
}
