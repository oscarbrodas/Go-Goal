import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { GoalsComponent } from './goals/goals.component';
import { UserComponent } from './user/user.component';

// CREATE ALL USER LOGGED IN COMPONENTS HERE

@NgModule({
  declarations: [
    GoalsComponent,
    UserComponent
  ],
  imports: [
    CommonModule
  ]
})
export class UserModule {


}
