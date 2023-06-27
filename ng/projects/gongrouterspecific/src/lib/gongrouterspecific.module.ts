import { NgModule } from '@angular/core';
import { GongrouterspecificComponent } from './gongrouterspecific.component';
import { TableOutletComponent } from './table-outlet/table-outlet.component'

import { Routes, RouterModule } from '@angular/router';


@NgModule({
  declarations: [
    GongrouterspecificComponent,
    TableOutletComponent
  ],
  imports: [
    RouterModule,
  ],
  exports: [
    GongrouterspecificComponent,
    TableOutletComponent
  ]
})
export class GongrouterspecificModule { }
