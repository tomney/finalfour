import { Component, OnInit } from '@angular/core';
import { FinalFourSelection } from '../final-four-selection';
import { FinalFourService } from '../final-four.service';
import { Team } from '../team';

@Component({
  selector: 'app-list-selections',
  templateUrl: './list-selections.component.html',
  styleUrls: ['./list-selections.component.css']
})
export class ListSelectionsComponent implements OnInit {
  selections: FinalFourSelection[]

  constructor(private finalFourService: FinalFourService) { }

  ngOnInit() {
    this.selections = []
    this.listSelections()
  }

  listSelections(): void {
    this.finalFourService.listSelections().subscribe(
      response => this.convertSelectionsResponse(response)
    )
    console.log(this.selections);
  }

  //TODO move this to a more angular-approps locash
  private convertSelectionsResponse(selections: any[]): void {
    selections.forEach(selectionResponse => {
        let selection: FinalFourSelection = new FinalFourSelection;
        selection.email = selectionResponse.Email;
        selection.teams = this.convertTeamsResponse(selectionResponse.Teams)
        this.selections.push(selection)
    });
  }

  //TODO move this to a more angular-approps locash
  private convertTeamsResponse(teamsResponse: any[]): Team[] {
    let teams: Team[] = [];
    teamsResponse.forEach(teamResponse => {
      let team: Team = new Team;
      team.id = teamResponse.ID;
      team.imageUrl = teamResponse.ImageURL;
      team.name = teamResponse.Name;
      teams.push(team)
    });
  
    return teams;
  }
}
