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

  constructor(private httpClient: HttpClient) { }

  submitSelection(ffs: FinalFourSelection): Observable<any> {
    //TODO call out to the backend to submit the finalFour
    console.log("Nice picks!")
    return this.httpClient.post<FinalFourSelection>(this.setSelectionUrl, ffs, httpOptions);
  }
}
