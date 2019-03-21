import { Team } from './team'

export class Validity {
    valid: boolean;
    err: string;
}

//TODO make this a top level import
export class FinalFourSelection {
    email: string;
    teams: Team[];

    validate(): Validity {
        if(!this.hasFourTeams()){
            return {valid: false, err: "Please select four teams"};
        }
        if(!this.hasValidEmail()){
            return {valid: false, err: "Please enter your name"};
        }
        return {valid: true, err: ""};
    }

    private hasValidEmail(): boolean {
        return typeof this.email != undefined && this.email != "";
    }

    private hasFourTeams(): boolean {
        return this.teams.length === 4;
    }
} 