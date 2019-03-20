import { Component, OnInit } from '@angular/core';
import {FinalFourSelection} from '../final-four-selection';

@Component({
  selector: 'app-list-selections',
  templateUrl: './list-selections.component.html',
  styleUrls: ['./list-selections.component.css']
})
export class ListSelectionsComponent implements OnInit {
  selections: FinalFourSelection[]

  constructor() { }

  ngOnInit() {
    this.listSelections()
  }

  listSelections(): void {
    //TODO write a list selections function
    console.log("ooh we are listin those selections")
  }


}
