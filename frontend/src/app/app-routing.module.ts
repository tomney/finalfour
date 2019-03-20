import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';
import { FourSelectorComponent } from './four-selector/four-selector.component';
import { ListSelectionsComponent }   from './list-selections/list-selections.component';


const routes: Routes = [
  { path: '', redirectTo: '/selector', pathMatch: 'full' },
  { path: 'selector', component: FourSelectorComponent },
  { path: 'selections', component: ListSelectionsComponent },
]

@NgModule({
  exports: [
    RouterModule
  ], 
  imports: [ 
    RouterModule.forRoot(routes) 
  ],
})
export class AppRoutingModule { }