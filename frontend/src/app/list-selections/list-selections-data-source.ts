import { CollectionViewer, DataSource } from "@angular/cdk/collections";
import { Selection } from "./selection";
import { Team } from "../team";

import { BehaviorSubject, Observable, of } from 'rxjs';
import { catchError, finalize } from 'rxjs/operators'; 

import { FinalFourService } from '../final-four.service';


export class SelectionDataSource implements DataSource<Selection> {
    private selectionsSubject = new BehaviorSubject<Selection[]>([]);

    constructor(private finalFourService: FinalFourService) {}

    connect(collectionViewer: CollectionViewer): Observable<Selection[]> {
        return this.selectionsSubject.asObservable();
    }

    disconnect(collectionViewer: CollectionViewer): void {
        this.selectionsSubject.complete();
    }
  
    loadSelections() {
        this.finalFourService.listSelections().pipe(
            catchError(() => of([])),
            finalize(() => console.log("Done"))
        )
        .subscribe(selections => {
            console.log(selections)
            if(selections){
                let formattedSelections: Selection[] = this.formatSelections(selections);
                this.selectionsSubject.next(formattedSelections)
            }
        });
    }  
    
    //TODO move this to a more angular-approps locash
    private formatSelections(selections: any[]): Selection[] {
        let formattedSelections: Selection[] = [];
        selections.forEach(selectionResponse => {
            let selection: Selection = new Selection;
            selection.email = selectionResponse.Email;
            selection.first = this.convertTeamResponse(selectionResponse.Teams[0]);
            selection.second = this.convertTeamResponse(selectionResponse.Teams[1]);
            selection.third = this.convertTeamResponse(selectionResponse.Teams[2]);
            selection.fourth = this.convertTeamResponse(selectionResponse.Teams[3]);
            formattedSelections.push(selection);
        });
        return formattedSelections;

    }

    //TODO move this to a more angular-approps locash
    private convertTeamResponse(teamResponse: any): Team {
        let team: Team = new Team;
        team.id = teamResponse.ID;
        team.imageUrl = teamResponse.ImageURL;
        team.name = teamResponse.Name;
        return team;
    }
}