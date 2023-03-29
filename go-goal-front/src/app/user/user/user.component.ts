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

  ngOnInit() {

  }



}