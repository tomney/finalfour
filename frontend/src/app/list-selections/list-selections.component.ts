import { Component, OnInit } from '@angular/core';
import { FinalFourSelection } from '../final-four-selection';
import { FinalFourService } from '../final-four.service';

@Component({
  selector: 'app-list-selections',
  templateUrl: './list-selections.component.html',
  styleUrls: ['./list-selections.component.css']
})
export class ListSelectionsComponent implements OnInit {
  selections: FinalFourSelection[]

  constructor(private finalFourService: FinalFourService) { }

  ngOnInit() {
    this.listSelections()
  }

  listSelections(): void {
    this.finalFourService.listSelections().subscribe(
      data => console.log(data)
    )
  }


}
