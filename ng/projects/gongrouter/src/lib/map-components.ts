// insertion point sub template for components imports 
  import { EditorOutletsTableComponent } from './editoroutlets-table/editoroutlets-table.component'
  import { EditorOutletSortingComponent } from './editoroutlet-sorting/editoroutlet-sorting.component'
  import { OutletsTableComponent } from './outlets-table/outlets-table.component'
  import { OutletSortingComponent } from './outlet-sorting/outlet-sorting.component'
  import { TableOutletsTableComponent } from './tableoutlets-table/tableoutlets-table.component'
  import { TableOutletSortingComponent } from './tableoutlet-sorting/tableoutlet-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfEditorOutletsComponents: Map<string, any> = new Map([["EditorOutletsTableComponent", EditorOutletsTableComponent],])
  export const MapOfEditorOutletSortingComponents: Map<string, any> = new Map([["EditorOutletSortingComponent", EditorOutletSortingComponent],])
  export const MapOfOutletsComponents: Map<string, any> = new Map([["OutletsTableComponent", OutletsTableComponent],])
  export const MapOfOutletSortingComponents: Map<string, any> = new Map([["OutletSortingComponent", OutletSortingComponent],])
  export const MapOfTableOutletsComponents: Map<string, any> = new Map([["TableOutletsTableComponent", TableOutletsTableComponent],])
  export const MapOfTableOutletSortingComponents: Map<string, any> = new Map([["TableOutletSortingComponent", TableOutletSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["EditorOutlet", MapOfEditorOutletsComponents],
      ["Outlet", MapOfOutletsComponents],
      ["TableOutlet", MapOfTableOutletsComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["EditorOutlet", MapOfEditorOutletSortingComponents],
      ["Outlet", MapOfOutletSortingComponents],
      ["TableOutlet", MapOfTableOutletSortingComponents],
    ]
  )
