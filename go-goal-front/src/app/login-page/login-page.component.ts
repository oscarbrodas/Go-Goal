import { Component } from '@angular/core';
import { LoginService } from './login.service';
import { BackendConnectService } from '../backend-connect.service';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { loginInfo } from '../backend-connect.service';
import { ReactiveFormsModule } from '@angular/forms';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css'],

})
export class LoginPageComponent {

  loginForm = this.formBuilder.group({
    Email: new FormControl(""),
    Password: new FormControl(""),
  });


  constructor(
    private formBuilder: FormBuilder,
    private loginService: LoginService
  ) { } // INJECT: BACKEND SERVICE as needed


  Submit(loginForm: FormGroup): void {
    this.loginService.login(loginForm);
  }

  checkFailedLogin(): boolean {
    return this.loginService.loginFailed
  }

  checkSuccessLogin(): boolean {
    return this.loginService.loginSuccess
  }


}


