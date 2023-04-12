import { ChangeDetectionStrategy, Component, HostListener, Input, OnChanges, OnDestroy, OnInit, Renderer2, SimpleChanges } from '@angular/core';
import { fromEvent, Subscription } from 'rxjs';
import { trigger, state, style, transition, animate, keyframes, group, query } from '@angular/animations';
import { Router } from '@angular/router';


@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css'],
  animations: [

    trigger('titlePanel', [

      transition(':enter', [
        style({ position: 'relative', top: '-100%' }),
        animate('0.8s 0.5s ease', keyframes([
          style({ position: 'relative', top: '-100%' }),
          style({ position: 'relative', top: '*' })
        ]))
      ]),

    ]),

    trigger('joinButton', [

      transition(':enter', [
        style({ position: 'relative', top: '-20%', opacity: 0 }),
        animate('0.5s 1.2s ease', keyframes([
          style({ position: 'relative', top: '-20%', opacity: 1, offset: 0 }),
          style({ position: 'relative', top: '*', opacity: 1, offset: 1 })
        ]))
      ]),

    ]),

    trigger('icons', [

      transition(':enter', [
        style({ top: '48%', left: '48%', opacity: 0 }),
        group([
          animate('0.5s 1.5s ease', keyframes([
            style({ top: '48%', left: '48%', opacity: 1, offset: 0 }),
            style({ top: '*', left: '*', opacity: 1, offset: 1 })
          ])),
        ])
      ]),

    ]),



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


  }

  @HostListener('window:scroll', ['$event']) onWindowScroll(e: any) {
    if (window.pageYOffset > window.innerHeight) {



    }

  }

  constructor(private renderer: Renderer2, private router: Router) {
    this.onResize();
  }


  ngOnInit(): void {
    this.orginalScreenHeight = window.innerHeight;
    window.scrollTo(0, 0);
  }






  join() {
    this.router.navigate(['/sign-up']);
  }



}


