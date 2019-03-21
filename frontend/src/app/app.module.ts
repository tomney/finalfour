import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { MatCardModule, MatButtonModule, MatCheckboxModule, MatSnackBarModule,  MatFormFieldModule, MatInputModule, MatTable, MatTableDataSource, MatTableModule, MatToolbarModule} from '@angular/material'
import { AppComponent } from './app.component';
import { AppRoutingModule } from './app-routing.module';
import { AppService } from './app.service';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FourSelectorComponent } from './four-selector/four-selector.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { ListSelectionsComponent } from './list-selections/list-selections.component';

@NgModule({
  declarations: [
    AppComponent,
    FourSelectorComponent,
    ListSelectionsComponent
  ],
  imports: [
    AppRoutingModule,
    BrowserModule,
    BrowserAnimationsModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    MatCardModule,
    MatButtonModule,
    MatCheckboxModule,
    MatSnackBarModule,
    MatTableModule,
    MatToolbarModule,
    MatFormFieldModule,
    MatInputModule,
  ],
  providers: [AppService],
  bootstrap: [AppComponent]
})
export class AppModule { }
