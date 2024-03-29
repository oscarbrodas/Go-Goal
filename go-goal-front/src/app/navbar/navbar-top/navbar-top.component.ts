import { Component, Inject, Input, SimpleChanges } from '@angular/core';
import { BreakpointObserver, BreakpointState } from '@angular/cdk/layout';
import { UserService } from 'src/app/user/user.service';
import { LoginService } from 'src/app/login-page/login.service';
import { MatDialogRef, MatDialog, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { trigger, state, style, transition, animate, keyframes } from '@angular/animations';
import { BackendConnectService } from 'src/app/backend-connect.service';

@Component({
  selector: 'app-navbar-top',
  templateUrl: './navbar-top.component.html',
  styleUrls: ['./navbar-top.component.css'],
  animations: [
    trigger('menuTrigger', [

      state('true', style({ position: '*', right: '-5px' })),
      state('false', style({ position: '*', right: '-400px' })),

      transition('false => true', [
        animate(450, keyframes([
          style({ offset: 0 }),
          style({ transform: 'translateX(-395px)', offset: 0.45 }),
          style({ transform: 'translateX(-380px)', offset: 0.5 }),
          style({ transform: 'translateX(-395px)', offset: 0.6 }),
          style({ transform: 'translateX(-392px)', offset: 0.65 }),
          style({ transform: 'translateX(-395px)', offset: 0.75 }),
          style({ transform: 'translateX(-394px)', offset: 0.85 }),
          style({ transform: 'translateX(-395px)', offset: 1 }),
        ]))
      ]),

      transition('true => false', [
        animate(200, keyframes([
          style({ transform: 'translateX(0px)', offset: 0 }),
          style({ transform: 'translateX(400px)', offset: 1 }),
        ]))
      ])
    ])
  ]
})
export class NavbarTopComponent {

  screenThin: boolean = false;
  @Input() verified: boolean = this.userService.isLoggedIn();
  subMenu: boolean = false;
  username: string = ''

  constructor(
    private breakpointObserver: BreakpointObserver, private userService: UserService, private loginService: LoginService, public dialog: MatDialog, private router: Router, private backend: BackendConnectService) {

    // Gets rid of Title if screen is too thin
    this.breakpointObserver.observe([
      "(max-width: 550px)"
    ]).subscribe((result: BreakpointState) => {
      if (result.matches) {
        this.screenThin = true;
      } else {
        this.screenThin = false;
      }
    });
  }

  ngOnInit() {

    if (this.userService.getUserData() == null) {
      console.log('No user data');
      this.verified = false;
    }
    else if (this.userService.isLoggedIn()) {
      this.verified = true;
      this.username = this.userService.getUserData().Username;
    }
    else {
      this.verified = false;
    }


  }

  ngOnChanges(changes: SimpleChanges) {
    console.log(changes);


  }

  ngOnDestroy() {
    console.log('destroyed');
  }


  toggleMenu(): void {
    this.subMenu = !this.subMenu;
  }
  profilePage(): void {
    this.router.navigate([`user/${this.userService.getUserData().ID}/profile`]);
  }
  goalsPage(): void {
    this.router.navigate([`user/${this.userService.getUserData().ID}/goals`]);
  }
  discoverPage(): void {
    this.router.navigate([`user/${this.userService.getUserData().ID}/discover`]);
  }
  settingsPage(): void {
    this.router.navigate([`user/${this.userService.getUserData().ID}/settings`]);

  }
  logout() {
    this.openDialog('500ms', '50ms');
    this.verified = false;
  }

  openDialog(enterAnimationDuration: string, exitAnimationDuration: string): void {
    this.verified = true;
    const dialogRef = this.dialog.open(logoutDialog, {
      width: '250px',
      enterAnimationDuration,
      exitAnimationDuration,
    });
  }

}

// Dialog component for logout
@Component({
  selector: 'logout-dialog',
  templateUrl: './logout-dialog.html',
  styleUrls: ['./navbar-top.component.css']
})
export class logoutDialog {
  constructor(
    public dialogRef: MatDialogRef<logoutDialog>, private loginService: LoginService,
    @Inject(MAT_DIALOG_DATA) public data: DialogData, private router: Router
  ) { }

  logout(): void {
    console.log("Logging out...");
    this.loginService.logout();
    this.dialogRef.close();
  }

  onNoClick(): void {
    this.dialogRef.close();
    this.router.navigate([`${this.router.url}`]).then(() => {
      window.location.reload();
    });
  }
}

export interface DialogData {
  message: string;
}
