import { Injectable } from '@angular/core';
import { FinalFourSelection } from './final-four-selection'
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

@Injectable({
  providedIn: 'root'
})
export class FinalFourService {
  setSelectionUrl = 'api/v1/setSelection'
  listSelectionUrl = 'api/v1/listSelections'

  constructor(private httpClient: HttpClient) { }

  submitSelection(ffs: FinalFourSelection): Observable<any> {
    return this.httpClient.post<FinalFourSelection>(this.setSelectionUrl, ffs, httpOptions);
  }

  listSelections(): Observable<any> {
    return this.httpClient.get(this.listSelectionUrl, httpOptions);
  }
}
