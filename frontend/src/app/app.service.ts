import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { logging } from 'protractor';

@Injectable()
export class AppService {
  constructor(private http: HttpClient) { }


  getGreeting() {
    return this.http.get("api/v1/hello");
  } 
}