import { Component, Inject } from '@angular/core';
import { BreakpointObserver, BreakpointState } from '@angular/cdk/layout';
import { UserService } from 'src/app/user/user.service';
import { LoginService } from 'src/app/login-page/login.service';
import { MatDialogRef, MatDialog, MAT_DIALOG_DATA } from '@angular/material/dialog';

@Component({
  selector: 'app-navbar-top',
  templateUrl: './navbar-top.component.html',
  styleUrls: ['./navbar-top.component.css']
})
export class NavbarTopComponent {

  screenThin: boolean = false;
  verified: boolean = false;

  constructor(
    private breakpointObserver: BreakpointObserver, private userService: UserService, private loginService: LoginService, public dialog: MatDialog) {

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
    this.verified = this.userService.isLoggedIn();
  }

  logout() {
    this.openDialog('500ms', '50ms');


  }

  openDialog(enterAnimationDuration: string, exitAnimationDuration: string): void {
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
    @Inject(MAT_DIALOG_DATA) public data: DialogData
  ) { }

  logout(): void {
    console.log("Logging out...");
    this.loginService.logout();
    this.dialogRef.close();
  }

  onNoClick(): void {
    this.dialogRef.close();
  }
}

export interface DialogData {
  message: string;
}
