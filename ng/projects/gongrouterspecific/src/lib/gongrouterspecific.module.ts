import { NgModule } from '@angular/core';
import { GongrouterspecificComponent } from './gongrouterspecific.component';
import { TableOutletComponent } from './table-outlet/table-outlet.component'

import { Routes, RouterModule } from '@angular/router';
import { EditorOutletComponent } from './editor-outlet/editor-outlet.component';
import { OutletComponent } from './outlet/outlet.component';
import { GongrouterOutletComponent } from './gongrouter-outlet/gongrouter-outlet.component';


@NgModule({
  declarations: [
    GongrouterspecificComponent,
    TableOutletComponent,
    EditorOutletComponent,
    OutletComponent,
    GongrouterOutletComponent
  ],
  imports: [
    RouterModule,
  ],
  exports: [
    GongrouterspecificComponent,
    TableOutletComponent,
    EditorOutletComponent
  ]
})
export class GongrouterspecificModule { }
