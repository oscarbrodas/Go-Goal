import { Component, Input, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, NgForm, FormControl } from '@angular/forms'
import { BackendConnectService, loginInfo } from '../backend-connect.service';
import { HttpClient } from '@angular/common/http';
import { userInfo } from '../backend-connect.service';
import { LoginService } from '../login-page/login.service';

@Component({
  selector: 'app-sign-up-component',
  templateUrl: './sign-up-component.component.html',
  styleUrls: ['./sign-up-component.component.css'],
  //providers: [BackendConnectService] Not ctually needed
})
export class SignUpComponentComponent {

  constructor(
    private backend: BackendConnectService,
    private http: HttpClient,
    private formBuilder: FormBuilder,
    private loginService: LoginService
  ) { }

  userForm = this.formBuilder.group({
    loggedIn: new FormControl(false),
    ID: new FormControl(0),
    FirstName: new FormControl(''),
    LastName: new FormControl(''),
    Email: new FormControl(''),
    Username: new FormControl(''),
    Password: new FormControl('')
  });
  loginForm = this.formBuilder.group({
    Email: new FormControl(""),
    Password: new FormControl(""),
  });
  userData: userInfo = {
    loggedIn: false,
    ID: 0,
    FirstName: '',
    LastName: '',
    Email: '',
    Username: '',
    Password: ''
  };

  signUpMessage?: string;
  submitted: boolean = false;
  users: userInfo[] = [];

  Submit(uf: FormGroup): void {
    this.submitted = true

    // Map Form Data to User Data
    this.userData.loggedIn = uf.value.loggedIn;
    this.userData.ID = uf.value.ID;
    this.userData.FirstName = uf.value.FirstName;
    this.userData.LastName = uf.value.LastName;
    this.userData.Email = uf.value.Email;
    this.userData.Username = uf.value.Username;
    this.userData.Password = uf.value.Password;

    // Attempt to Sign Up
    if (!this.userData.Email.includes('@') && !this.userData.Email.includes('.')) {
      console.log(this.userData);

      this.signUpMessage = 'Not a valid email address'
      return;
    }
    else if (this.userData.Password.length < 8) {
      this.signUpMessage = 'This account needs a more secure password'
      return;
    }
    else {

      this.backend.signThemUp(this.userData).subscribe((data) => {
        console.log('User Sign up request sent');

        // Login Procedure after successful sign up
        this.loginForm.controls['Email'].setValue(this.userData.Email);
        this.loginForm.controls['Password'].setValue(this.userData.Password);
        this.loginService.login(this.loginForm);
        // DO: Reroute to User Page
        this.signUpMessage = 'Account Created! Redirecting...'
      });
    }
  }
}




