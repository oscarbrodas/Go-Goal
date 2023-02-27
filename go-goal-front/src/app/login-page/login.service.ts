import { Injectable } from '@angular/core';
import { BackendConnectService } from '../backend-connect.service';
import { LoginPageComponent } from './login-page.component';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor(private backend: BackendConnectService) { } // INJECT: BACKEND SERVICE
  user: userInfo = {
    loggedIn: false,
    Username: '',
    FirstName: '',
    LastName: '',
    Email: '',
    Password: ''

  }
  users: userInfo[] = [];
  loginFailed: boolean = false;

  getUser(): userInfo { return this.user }

  loggedInStatus(): boolean { return this.user.loggedIn }

  clearUser(): void {
    this.user = {
      loggedIn: false,
      Username: '',
      FirstName: '',
      LastName: '',
      Email: '',
      Password: ''

    }
  }


  login(): void {
    // ADD: Get and submit loginForm to backend for verification from loginComponent
    
    // ADD: Get Data using http, update current user data and loggedin status in login service
    if (this.user.loggedIn) {
      this.backend.getLoginInfo().subscribe(users => this.users = users); 
      console.log("Successfully logged in.");
    }
    else {
      console.log('ERROR: Login in status failed to update');
      this.loginFailed = true;
    }

    this.verifyLogin(this.user);
    this.users = []
  }

  verifyLogin(user: userInfo): void {

    if (user.loggedIn) { } // CHANGE & ADD: Reroute to User Page
    else {
      this.loginFailed = true;
    }
  }

  logout(): void {
    // ADD: Logout functionality as needed

    this.clearUser();
    this.loginFailed = false;

  }




}

export interface userInfo { // ADD: User data as necessary 
  loggedIn: boolean;
  Username: string;
  FirstName: string;
  LastName: string;
  Email: string;
  Password: string;


}
