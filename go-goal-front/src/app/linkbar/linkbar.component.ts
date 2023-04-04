import { Component, OnInit, OnChanges, SimpleChanges, Input } from '@angular/core';
import { UserService } from '../user/user.service';
import { LoginService } from '../login-page/login.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-linkbar',
  templateUrl: './linkbar.component.html',
  styleUrls: ['./linkbar.component.css']
})
export class LinkbarComponent implements OnInit, OnChanges {

  @Input() loggedIn: boolean = false;
  @Input() userID: number = 0;

  constructor(private userService: UserService, private loginService: LoginService, private router: Router) { }

  ngOnInit(): void {
    this.loggedIn = this.userService.isLoggedIn();
    this.userID = this.userService.getUserData().ID;
  }

  ngOnChanges(changes: SimpleChanges): void {
  }

  top() {
    window.scroll(0, 0);
  }
  logout() { this.loginService.logout() }
  profile() {
    this.router.navigate([`user/${this.userID}/profile`]);
  }

}
