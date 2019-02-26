import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FourSelectorComponent } from './four-selector.component';

describe('FourSelectorComponent', () => {
  let component: FourSelectorComponent;
  let fixture: ComponentFixture<FourSelectorComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FourSelectorComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FourSelectorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
