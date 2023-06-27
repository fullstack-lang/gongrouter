import { NgModule } from '@angular/core';
import { GongrouterspecificComponent } from './gongrouterspecific.component';
import { TableOutletComponent } from './table-outlet/table-outlet.component'

import { Routes, RouterModule } from '@angular/router';
import { EditorOutletComponent } from './editor-outlet/editor-outlet.component';


@NgModule({
  declarations: [
    GongrouterspecificComponent,
    TableOutletComponent,
    EditorOutletComponent
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
