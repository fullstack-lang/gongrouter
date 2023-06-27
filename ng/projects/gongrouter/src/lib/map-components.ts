// insertion point sub template for components imports 
  import { TableOutletsTableComponent } from './tableoutlets-table/tableoutlets-table.component'
  import { TableOutletSortingComponent } from './tableoutlet-sorting/tableoutlet-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfTableOutletsComponents: Map<string, any> = new Map([["TableOutletsTableComponent", TableOutletsTableComponent],])
  export const MapOfTableOutletSortingComponents: Map<string, any> = new Map([["TableOutletSortingComponent", TableOutletSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["TableOutlet", MapOfTableOutletsComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["TableOutlet", MapOfTableOutletSortingComponents],
    ]
  )
