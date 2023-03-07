import { Injectable } from '@angular/core';
import { BackendConnectService, userInfo } from '../backend-connect.service';
import { LoginPageComponent } from './login-page.component';
import { HttpClient } from '@angular/common/http';
import { ActivationStart, Router, ActivatedRoute } from '@angular/router';
import { loginInfo } from '../backend-connect.service';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  loginFailed: boolean = false;
  loggedIn: boolean = false;

  constructor(private backend: BackendConnectService, private route: ActivatedRoute, private router: Router) {

    // Checks if user leaves login page, if so, resets loginFailed to false
    this.router.events.subscribe((event) => {
      if (event instanceof ActivationStart) {
        if (this.route.component != LoginPageComponent) {
          this.loginFailed = false;
        }
      }
    });

  }

  user: userInfo = {
    loggedIn: false,
    Username: '',
    FirstName: '',
    LastName: '',
    Email: '',
    Password: ''

  }
  loginInfo: loginInfo = {
    Email: '',
    Password: ''
  }
  users: userInfo[] = [];


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


  login(): void { // DOES NOT WORK YET
    // ADD: Get and submit loginForm to backend for verification from loginComponent

    // ADD: Get Data using http, update current user data and loggedin status in login service
    if (this.user.loggedIn) {
      // this.backend.getLoginInfo().subscribe(() => { }); // ADD: Get user data from backend ONCE BACKEND IS CONNECTED
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


