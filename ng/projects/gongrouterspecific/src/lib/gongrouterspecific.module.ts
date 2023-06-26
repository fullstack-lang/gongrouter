import { NgModule } from '@angular/core';
import { GongrouterspecificComponent } from './gongrouterspecific.component';
import { TriageComponent } from './triage/triage.component';



@NgModule({
  declarations: [
    GongrouterspecificComponent,
    TriageComponent
  ],
  imports: [
  ],
  exports: [
    GongrouterspecificComponent,
    TriageComponent
  ]
})
export class GongrouterspecificModule { }
