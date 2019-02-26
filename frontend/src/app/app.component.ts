import { Component } from '@angular/core';
import { AppService } from './app.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'angular-finalfour';
  greeting = '';

  constructor(private appService: AppService){}

  ngOnInit(){
  }

  
}
