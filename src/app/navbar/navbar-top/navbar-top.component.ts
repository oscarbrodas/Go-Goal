import { Component } from '@angular/core';
import { BreakpointObserver, BreakpointState } from '@angular/cdk/layout';

@Component({
  selector: 'app-navbar-top',
  templateUrl: './navbar-top.component.html',
  styleUrls: ['./navbar-top.component.css']
})
export class NavbarTopComponent {

  screenThin: boolean = false;

  constructor(
    private breakpointObserver: BreakpointObserver,
  ) {
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

}
