import { Component, OnInit } from '@angular/core';
import { Team } from '../team'
import { TeamsService } from '../teams.service'
import { FinalFourSelection } from './final-four-selection'
import { MatCheckboxChange } from '@angular/material'
import { MatSnackBar } from '@angular/material'

@Component({
  selector: 'app-four-selector',
  templateUrl: './four-selector.component.html',
  styleUrls: ['./four-selector.component.css']
})
export class FourSelectorComponent implements OnInit {
  teams: Team[]
  finalFourSelection: FinalFourSelection
  constructor(private teamsService: TeamsService, private snackBar: MatSnackBar) { }

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
    if(this.finalFourSelection.teams.length >= 4){
      this.showSnackBar();
    } else {
      this.finalFourSelection.teams.push(team);
    }    
  }

  removeSelection(team: Team): void {
    this.finalFourSelection.teams = this.finalFourSelection.teams.filter(t => t != team)
  }

  showSnackBar(): void {
    this.snackBar.open("You can only select four teams", "ERROR", {
      duration: 2000,
    });
  }

  selectionIsFull(): boolean {
    return this.finalFourSelection.teams.length == 4;
  }

  teamIsSelected(team: Team): boolean {
    return this.finalFourSelection.teams.indexOf(team)!==-1;
  }
}
