import { TestBed } from '@angular/core/testing';

import { FinalFourService } from './final-four.service';

describe('FinalFourService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: FinalFourService = TestBed.get(FinalFourService);
    expect(service).toBeTruthy();
  });
});
