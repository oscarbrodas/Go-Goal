import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { GoalsComponent } from './goals/goals.component';

// CREATE ALL USER LOGGED IN COMPONENTS HERE

@NgModule({
  declarations: [
    GoalsComponent
  ],
  imports: [
    CommonModule
  ]
})
export class UserModule { }
