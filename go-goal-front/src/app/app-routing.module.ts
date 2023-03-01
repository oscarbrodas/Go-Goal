import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { MainComponent } from './main/main.component';
import {LoginPageComponent} from './login-page/login-page.component'
import {SignUpComponentComponent} from './sign-up-component/sign-up-component.component'

const routes: Routes = [
  {path: 'main', component: MainComponent},
  {path: '', redirectTo: '/main', pathMatch: 'full'},
  {path: 'login', component: LoginPageComponent},
  {path: 'sign-up',component: SignUpComponentComponent},
  //Must leave 404 redirect as last route on page
  {path: '**', component: MainComponent}, // IMPLEMENT: pagenotfound-component, Change null url to said component
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
