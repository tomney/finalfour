import { Component } from '@angular/core';
import { AppService } from './app.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'angular-bolg';
  greeting = '';

  constructor(private appService: AppService){}

  ngOnInit(){
    this.showGreeting();
  }

  showGreeting(): void {
    this.appService.getGreeting().subscribe(
      data => {
        console.log("The return from our endpoint: ", data)
        this.greeting = data["greeting"]
      },
      err => console.log("The endpoint returned an error: ", err)
    )
  };
}
