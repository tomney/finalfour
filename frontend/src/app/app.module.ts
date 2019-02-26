import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { MatCardModule, MatButtonModule, MatCheckboxModule, MatSnackBarModule} from '@angular/material'
import { AppComponent } from './app.component';
import { AppService } from './app.service';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';
import { FourSelectorComponent } from './four-selector/four-selector.component';
import { FormsModule } from '@angular/forms';

@NgModule({
  declarations: [
    AppComponent,
    FourSelectorComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    FormsModule,
    HttpClientModule,
    MatCardModule,
    MatButtonModule,
    MatCheckboxModule,
    MatSnackBarModule
  ],
  providers: [AppService],
  bootstrap: [AppComponent]
})
export class AppModule { }
