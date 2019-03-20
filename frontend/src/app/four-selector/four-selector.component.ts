import { Component, OnInit } from '@angular/core';
import { Team } from '../team';
import { TeamsService } from '../teams.service';
import { FinalFourSelection } from '../final-four-selection';
import { MatCheckboxChange } from '@angular/material';
import { MatSnackBar } from '@angular/material';
import { FormControl } from '@angular/forms';
import { FinalFourService } from '../final-four.service';

@Component({
  selector: 'app-four-selector',
  templateUrl: './four-selector.component.html',
  styleUrls: ['./four-selector.component.css']
})
export class FourSelectorComponent implements OnInit {
  teams: Team[];
  email = new FormControl('');
  finalFourSelection = new FinalFourSelection;

  constructor(
    private teamsService: TeamsService, 
    private snackBar: MatSnackBar, 
    private finalFourService: FinalFourService
  ) { }

  ngOnInit() {
    this.getTeams();
    this.finalFourSelection.email = "";
    this.finalFourSelection.teams = [];
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
    if(this.finalFourSelection.teams.length < 4){
      this.finalFourSelection.teams.push(team);
    }
  }

  removeSelection(team: Team): void {
    this.finalFourSelection.teams = this.finalFourSelection.teams.filter(t => t != team)
  }

  showSnackBar(message: string, action: string): void {
    this.snackBar.open(message, action, {
      duration: 2000,
    });
  }

  selectionIsFull(): boolean {
    return this.finalFourSelection.teams.length == 4;
  }

  teamIsSelected(team: Team): boolean {
    return this.finalFourSelection.teams.indexOf(team)!==-1;
  }

  submit(): void {
    this.finalFourSelection.email = this.email.value;
    let{ valid, err } = this.finalFourSelection.validate();
    if(valid !== true){
      this.showSnackBar(err, "ERROR")
    } else {
      this.finalFourService.submitSelection(this.finalFourSelection).subscribe(
        success => this.showSnackBar("Picks submitted!", "SUCCESS"),
        err => this.showSnackBar("An error occurred submitting the picks", "ERROR")
      );
    }
  }
}
