import { Component } from '@angular/core';
import { LoginService } from './login.service';
import { BackendConnectService } from '../backend-connect.service';
import { FormBuilder } from '@angular/forms';
import { loginInfo } from '../backend-connect.service';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css'],

})
export class LoginPageComponent {


  constructor(
    private formBuilder: FormBuilder,
    private loginService: LoginService
  ) { } // INJECT: BACKEND SERVICE as needed

  loginForm: loginInfo = {
    Email: '',
    Password: ''
  }

  Submit(loginForm: loginInfo): void {
    console.log(loginForm);
  }

  checkFailedLogin(): boolean {
    return this.loginService.loginFailed
  }


}


