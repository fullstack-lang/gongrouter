// insertion point sub template for components imports 
  import { TriagesTableComponent } from './triages-table/triages-table.component'
  import { TriageSortingComponent } from './triage-sorting/triage-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfTriagesComponents: Map<string, any> = new Map([["TriagesTableComponent", TriagesTableComponent],])
  export const MapOfTriageSortingComponents: Map<string, any> = new Map([["TriageSortingComponent", TriageSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Triage", MapOfTriagesComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Triage", MapOfTriageSortingComponents],
    ]
  )
