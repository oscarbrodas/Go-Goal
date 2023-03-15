import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { NotFoundComponent } from '../not-found/not-found.component';
import { UserComponent } from './user/user.component';

const routes: Routes = [
  { path: 'home', component: UserComponent },
  { path: 'goals', component: NotFoundComponent },
  { path: 'discover', component: NotFoundComponent },
  { path: 'settings', component: NotFoundComponent },
  { path: '**', redirectTo: 'home' }

];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UserRoutingModule { }
