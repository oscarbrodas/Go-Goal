import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { MainComponent } from './main/main.component';
import {LoginPageComponent} from './login-page/login-page.component'

const routes: Routes = [
  {path: 'main', component: MainComponent},
  {path: '', redirectTo: '/main', pathMatch: 'full'},
  {path: 'login', component: LoginPageComponent},

  //Must leave 404 redirect as last route on page
  {path: '**', component: LoginPageComponent}, // IMPLEMENT: pagenotfound-component, Change null url to said component
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
