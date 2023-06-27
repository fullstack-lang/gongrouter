import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TableOutletComponent } from './table-outlet.component';

describe('TableOutletComponent', () => {
  let component: TableOutletComponent;
  let fixture: ComponentFixture<TableOutletComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TableOutletComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TableOutletComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
