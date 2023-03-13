import { Component, Input, OnInit, Output } from '@angular/core';
import { Subscription } from 'cypress/types/net-stubbing';
import { userInfo } from 'src/app/backend-connect.service';
import { LoginService } from 'src/app/login-page/login.service';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {

  logggedIn: boolean = false;
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

  constructor(private loginService: LoginService) {
  }

  ngOnInit() {
    console.log('User page loaded.');
    this.logggedIn = this.loginService.loggedIn;
    this.user = this.loginService.user;
    console.log(this.user);


  }



}

