import { ChangeDetectionStrategy, Component, HostListener, Input, OnChanges, OnDestroy, OnInit, Renderer2, SimpleChanges } from '@angular/core';
import { fromEvent, Subscription } from 'rxjs';
import { trigger, state, style, transition, animate, keyframes } from '@angular/animations';
import { Router } from '@angular/router';


@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css'],
  animations: [



  ]
})
export class MainComponent implements OnInit {

  screenHeight?: number;
  orginalScreenHeight?: number;
  screenWidth?: number;

  section2active = false;
  section3active = false;

  @HostListener('window:resize', ['$event']) onResize(event?: any) {
    this.screenHeight = window.innerHeight;
    this.screenWidth = window.innerWidth;
    console.log('screenHeight', this.screenHeight);


  }

  @HostListener('window:scroll', ['$event']) onWindowScroll(e: any) {
    if (window.pageYOffset > window.innerHeight) {

      console.log('here');

    }

  }

  constructor(private renderer: Renderer2, private router: Router) {
    this.onResize();
  }


  ngOnInit(): void {
    this.orginalScreenHeight = window.innerHeight;
  }






  join() {
    this.router.navigate(['/sign-up']);
  }



}


