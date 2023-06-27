import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EditorOutletComponent } from './editor-outlet.component';

describe('EditorOutletComponent', () => {
  let component: EditorOutletComponent;
  let fixture: ComponentFixture<EditorOutletComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EditorOutletComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(EditorOutletComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
