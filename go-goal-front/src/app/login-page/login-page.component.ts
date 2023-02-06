import { Component } from '@angular/core';
import { BackendConnectService } from '../backend-connect.service';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css']
})
export class LoginPageComponent {
  /*
  data: string = "";
  Login(): void{
    this.backend.getLoginInfo().subscribe(data => this.data = data);
  }
  constructor(public backend: BackendConnectService){

  }

For reasons unclear to me at this time, this code makes login page no longer work
*/
}
