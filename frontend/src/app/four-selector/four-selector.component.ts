import { Component, OnInit } from '@angular/core';
import { Team } from '../team'
import { TeamsService } from '../teams.service'
import { FinalFourSelection } from './final-four-selection'
import { MatCheckboxChange } from '@angular/material'

@Component({
  selector: 'app-four-selector',
  templateUrl: './four-selector.component.html',
  styleUrls: ['./four-selector.component.css']
})
export class FourSelectorComponent implements OnInit {
  teams: Team[]
  finalFourSelection: FinalFourSelection
  constructor(private teamsService: TeamsService) { }

  ngOnInit() {
    this.getTeams();
    this.finalFourSelection = {email: '', teams: []};
  }

  getTeams(): void {
    this.teamsService.getTeams().subscribe(
      teams => this.teams = teams 
    );
  }

  toggleSelection(team: Team, change: MatCheckboxChange): void {
    if(change.checked == true){
      this.addSelection(team)
    } else {
      this.removeSelection(team)
    }
  }

  addSelection(team: Team): void {
    this.finalFourSelection.teams.push(team);
    if(this.finalFourSelection.teams.length > 4){
      this.finalFourSelection.teams.shift()
    }
  }

  removeSelection(team: Team): void {
    this.finalFourSelection.teams = this.finalFourSelection.teams.filter(t => t != team)
  }
}
