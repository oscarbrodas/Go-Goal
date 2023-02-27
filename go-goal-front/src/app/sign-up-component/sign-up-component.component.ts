import { Component, Input, OnInit } from '@angular/core';
import {NgForm} from '@angular/forms'
import { BackendConnectService } from '../backend-connect.service';
import { HttpClient } from '@angular/common/http';

import { userInfo } from '../login-page/login.service';
@Component({
  selector: 'app-sign-up-component',
  templateUrl: './sign-up-component.component.html',
  styleUrls: ['./sign-up-component.component.css'],
  //providers: [BackendConnectService] Not ctually needed
})
export class SignUpComponentComponent{
  constructor(private backend: BackendConnectService){}
  userData: userInfo = {loggedIn : false, FirstName : '', LastName : '', Email : '', Username : '', Password : ''};
  signUpMessage?: string;
  submitted: boolean = false;
  users: userInfo[] = [];
  Submit(userData: userInfo): void{
    this.submitted = true
    if(!userData.Email.includes('@')){
      this.signUpMessage = 'Not a valid email address'
      return;
    }
    else if(userData.Password.length < 8){
      this.signUpMessage = 'This account needs a more secure password'
      return;
    }
    else{
      this.backend.signThemUp(userData).subscribe(user => {this.users.push(user)})
      this.signUpMessage = 'Account Created!'
    }
    //To Add: Checks for valid email not used before, open username, strong password
  }

}