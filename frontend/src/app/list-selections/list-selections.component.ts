import { Component, OnInit } from '@angular/core';
import { FinalFourSelection } from '../final-four-selection';
import { MatTable, MatTableDataSource, MatTableModule, MatHeaderRow, MatHeaderRowDef, MatHeaderCell, MatHeaderCellDef } from '@angular/material';
import { FinalFourService } from '../final-four.service';
import { SelectionDataSource } from './list-selections-data-source';
import { Team } from '../team';
import { Selection } from './selection';

@Component({
  selector: 'app-list-selections',
  templateUrl: './list-selections.component.html',
  styleUrls: ['./list-selections.component.css']
})

export class ListSelectionsComponent implements OnInit {
  dataSource: SelectionDataSource;
  displayedColumns: string[] = ['email', 'first', 'second', 'third', 'fourth'];

  constructor(private finalFourService: FinalFourService) { }

  ngOnInit() {
    this.dataSource = new SelectionDataSource(this.finalFourService);
    this.dataSource.loadSelections();
  }
}
