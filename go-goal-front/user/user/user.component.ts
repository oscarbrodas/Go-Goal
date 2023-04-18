import { Component, OnInit } from '@angular/core';
import { userInfo } from 'src/app/backend-connect.service';
import { LoginService } from 'src/app/login-page/login.service';
import { UserService } from '../user.service';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {

  user: userInfo = {
    loggedIn: false,
    ID: 0,
    Username: '',
    FirstName: '',
    LastName: '',
    Email: '',
    Password: ''
  }
  friends: userInfo[] = [];

  constructor(private loginService: LoginService, private userService: UserService) {

  }

  // Checks if user is logs in and saves data for refresh or redirect
  ngOnInit() {
    // Debugging Purposes
    // console.log('User page loaded.');

    if (this.userService.isLoggedIn() && !this.user.loggedIn) {
      this.user = this.userService.getUserData();
      //console.log('Returing user data from user service.');
    }
    else {
      this.user = this.loginService.user;
      this.userService.setUserData(this.user);
      // console.log('saving user data to user service.');
    }

    // Debugging Purposes
    if (this.user.loggedIn) {
      console.log('' + this.user.FirstName + ' ' + this.user.LastName + ' is logged in.');
    }
  }



}