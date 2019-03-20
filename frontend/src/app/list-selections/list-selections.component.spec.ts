import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ListSelectionsComponent } from './list-selections.component';

describe('ListSelectionsComponent', () => {
  let component: ListSelectionsComponent;
  let fixture: ComponentFixture<ListSelectionsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ListSelectionsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ListSelectionsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
