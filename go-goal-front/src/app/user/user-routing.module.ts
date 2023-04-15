import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { NotFoundComponent } from '../not-found/not-found.component';
import { UserComponent } from './user/user.component';
import { GoalsComponent } from './goals/goals.component';
import { SettingsComponent } from './settings/settings.component';
import { ProfileComponent } from './profile/profile.component';
import { DiscoverComponent } from './discover/discover.component';

const routes: Routes = [
  { path: 'profile', component: ProfileComponent },
  { path: 'goals', component: GoalsComponent },
  { path: 'discover', component: DiscoverComponent },
  { path: 'settings', component: SettingsComponent },
  { path: '', redirectTo: 'profile', pathMatch: 'full' },
  { path: '**', component: NotFoundComponent }

];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UserRoutingModule { }
