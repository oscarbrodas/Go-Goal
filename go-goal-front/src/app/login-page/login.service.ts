import { Injectable, SimpleChanges } from '@angular/core';
import { BackendConnectService, userInfo } from '../backend-connect.service';
import { LoginPageComponent } from './login-page.component';
import { HttpClient } from '@angular/common/http';
import { ActivationStart, Router, ActivatedRoute } from '@angular/router';
import { loginInfo } from '../backend-connect.service';
import { FormBuilder, FormGroup } from '@angular/forms';
import { UserService } from '../user/user.service';


@Injectable({
  providedIn: 'root'
})
export class LoginService {

  loginFailed: boolean = false;
  loginSuccess: boolean = false;
  loggedIn: boolean = false;

  constructor(private backend: BackendConnectService, private route: ActivatedRoute, private router: Router, private userService: UserService) {

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
    ID: 0,
    Username: '',
    FirstName: '',
    LastName: '',
    Email: '',
    Password: ''

  }

  public friends: number[] = [];

  loginInfo: loginInfo = {
    Email: '',
    Password: ''
  }
  users: userInfo[] = [];



  getUser(): userInfo { return this.user }
  loggedInStatus(): boolean { return this.user.loggedIn }

  // Clears User
  clearUser(): void {
    this.user = {
      loggedIn: false,
      ID: 0,
      Username: '',
      FirstName: '',
      LastName: '',
      Email: '',
      Password: ''

    }
  }

  // Logs in user
  login(li: FormGroup): void {
    this.loginInfo.Email = li.value.Email;
    this.loginInfo.Password = li.value.Password;

    // GET REQUEST FOR USER INFORMATION
    console.log("Attempting to log in...");
    this.backend.getLoginInfo(this.loginInfo).subscribe((data) => {

      // Checks if user exists in database and sets user data accordingly
      if (data.FindEmail == true && data.FindPassword == true) {
        console.log("Successfully logged in.");
        this.user.loggedIn = true;
        this.user.ID = data.ThisUser.ID;
        this.user.Username = data.ThisUser.Username;
        this.user.FirstName = data.ThisUser.FirstName;
        this.user.LastName = data.ThisUser.LastName;
        this.user.Email = data.ThisUser.Email;
        this.user.Password = data.ThisUser.Password;
        this.loginFailed = false;
        this.loginSuccess = true;
        this.loggedIn = true;

        this.backend.getFriends(this.user.ID).subscribe((data) => {
          if (!data.ErrorExist) {
            this.friends = data.Friends;
            console.log(this.friends);
          }

        });

        this.userService.setUserData(this.user);

        this.verifyLogin(this.user);

      }
      else {
        console.log('ERROR: Login in status failed to update');
        this.loginFailed = true;
        this.loginSuccess = false;
      }

    });


  }

  verifyLogin(user: userInfo): void {

    if (user.loggedIn) {
      console.log('Redirecting to user page...');
      this.router.navigate(['/user/' + this.user.ID + '/profile']).then(() => {
        window.location.reload();
      });

    }
    else {
      console.log("Unable to verify login.");
      this.loginFailed = true;
      this.loginSuccess = false;
    }
  }

  logout(): void {
    // ADD: Logout functionality as needed
    this.clearUser();
    this.loginFailed = false;
    this.loggedIn = false;
    this.loginSuccess = false;

    this.userService.clearUserData();
    this.userService.cleanStorage();
    this.userService.loggedIn = false;

    this.router.navigate(['/main']).then(() => {
      window.location.reload();
    });
  }




}


