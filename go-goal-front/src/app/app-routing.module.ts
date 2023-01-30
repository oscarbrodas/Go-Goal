import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { MainComponent } from './main/main.component';

const routes: Routes = [
  {path: 'main', component: MainComponent},
  {path: '**', component: MainComponent}, // IMPLEMENT: pagenotfound-component, Change null url to said component
  {path: '', redirectTo: '/main', pathMatch: 'full'}

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
