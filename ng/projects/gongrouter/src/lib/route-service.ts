import { Injectable } from '@angular/core';
import { Route, Router, Routes } from '@angular/router';

// insertion point for imports
import { EditorOutletsTableComponent } from './editoroutlets-table/editoroutlets-table.component'
import { EditorOutletDetailComponent } from './editoroutlet-detail/editoroutlet-detail.component'

import { OutletsTableComponent } from './outlets-table/outlets-table.component'
import { OutletDetailComponent } from './outlet-detail/outlet-detail.component'

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
    getEditorOutletTablePath(): string {
        return this.getPathRoot() + '-editoroutlets/:GONG__StackPath'
    }
    getEditorOutletTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getEditorOutletTablePath(), component: EditorOutletsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getEditorOutletAdderPath(): string {
        return this.getPathRoot() + '-editoroutlet-adder/:GONG__StackPath'
    }
    getEditorOutletAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getEditorOutletAdderPath(), component: EditorOutletDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getEditorOutletAdderForUsePath(): string {
        return this.getPathRoot() + '-editoroutlet-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getEditorOutletAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getEditorOutletAdderForUsePath(), component: EditorOutletDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getEditorOutletDetailPath(): string {
        return this.getPathRoot() + '-editoroutlet-detail/:id/:GONG__StackPath'
    }
    getEditorOutletDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getEditorOutletDetailPath(), component: EditorOutletDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getOutletTablePath(): string {
        return this.getPathRoot() + '-outlets/:GONG__StackPath'
    }
    getOutletTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getOutletTablePath(), component: OutletsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getOutletAdderPath(): string {
        return this.getPathRoot() + '-outlet-adder/:GONG__StackPath'
    }
    getOutletAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getOutletAdderPath(), component: OutletDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getOutletAdderForUsePath(): string {
        return this.getPathRoot() + '-outlet-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getOutletAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getOutletAdderForUsePath(), component: OutletDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getOutletDetailPath(): string {
        return this.getPathRoot() + '-outlet-detail/:id/:GONG__StackPath'
    }
    getOutletDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getOutletDetailPath(), component: OutletDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

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
            this.getEditorOutletTableRoute(stackPath),
            this.getEditorOutletAdderRoute(stackPath),
            this.getEditorOutletAdderForUseRoute(stackPath),
            this.getEditorOutletDetailRoute(stackPath),

            this.getOutletTableRoute(stackPath),
            this.getOutletAdderRoute(stackPath),
            this.getOutletAdderForUseRoute(stackPath),
            this.getOutletDetailRoute(stackPath),

            this.getTableOutletTableRoute(stackPath),
            this.getTableOutletAdderRoute(stackPath),
            this.getTableOutletAdderForUseRoute(stackPath),
            this.getTableOutletDetailRoute(stackPath),

        ])
    }
}
