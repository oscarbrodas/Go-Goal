import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { UserComponent } from './user/user.component';
import { userInfo } from '../backend-connect.service';

// CREATE ALL USER LOGGED IN COMPONENTS HERE

@NgModule({
  declarations: [
    UserComponent
  ],
  imports: [
    CommonModule
  ]
})
export class UserModule {

  user: userInfo = {
    loggedIn: false,
    ID: 0,
    Username: '',
    FirstName: '',
    LastName: '',
    Email: '',
    Password: ''
  }

  constructor() {
    console.log("User Module Loaded");


  }






}
