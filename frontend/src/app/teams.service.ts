import { Injectable } from '@angular/core';
import { Team } from './team';
import { TEAMS } from './teams';
import { of, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class TeamsService {

  constructor() { }

  getTeams(): Observable<Team[]> {
    return of(TEAMS);
  } 
}
