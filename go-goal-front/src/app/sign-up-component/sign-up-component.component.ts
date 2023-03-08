import { Component, Input, OnInit } from '@angular/core';
import { FormGroup, NgForm } from '@angular/forms'
import { BackendConnectService, loginInfo } from '../backend-connect.service';
import { HttpClient } from '@angular/common/http';
import { userInfo } from '../backend-connect.service';

@Component({
  selector: 'app-sign-up-component',
  templateUrl: './sign-up-component.component.html',
  styleUrls: ['./sign-up-component.component.css'],
  //providers: [BackendConnectService] Not ctually needed
})
export class SignUpComponentComponent {

  constructor(
    private backend: BackendConnectService,
    private http: HttpClient

  ) {
  }
  userData: userInfo = { loggedIn: false, ID: 0, FirstName: '', LastName: '', Email: '', Username: '', Password: '' };
  signUpMessage?: string;
  submitted: boolean = false;
  users: userInfo[] = [];

  Submit(userData: userInfo): void {
    this.submitted = true
    if (!userData.Email.includes('@')) {
      this.signUpMessage = 'Not a valid email address'
      return;
    }
    else if (userData.Password.length < 8) {
      this.signUpMessage = 'This account needs a more secure password'
      return;
    }
    else {
      this.backend.signThemUp(userData).subscribe(() => { console.log('User Sign up request sent') })
      this.signUpMessage = 'Account Created!';

      // TO DO: CHECK IF THE BACKEND STORED THE USER -- WORK WITH BACKEND TO DO THIS


    }
    //To Add: Checks for valid email not used before, open username, strong password
  }

}