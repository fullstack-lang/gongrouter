import { NgModule } from '@angular/core';
import { GongrouterspecificComponent } from './gongrouterspecific.component';
import { TriageComponent } from './triage/triage.component';

import { Routes, RouterModule } from '@angular/router';


@NgModule({
  declarations: [
    GongrouterspecificComponent,
    TriageComponent
  ],
  imports: [
    RouterModule,
  ],
  exports: [
    GongrouterspecificComponent,
    TriageComponent
  ]
})
export class GongrouterspecificModule { }
