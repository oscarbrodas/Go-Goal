import { Component } from '@angular/core';
import { trigger, state, style, transition, animate, keyframes, stagger, query } from '@angular/animations';

@Component({
  selector: 'app-aboutus',
  templateUrl: './aboutus.component.html',
  styleUrls: ['./aboutus.component.css'],
  animations: [
    trigger('tags', [

      transition(':enter', [
        query('.tag', [
          style({ position: 'relative', transform: 'translateY(-500px)', }),
          stagger(100, [
            animate('0.5s 0.2s ease-in-out', keyframes([
              style({ transform: 'translateY(-500px)', offset: 0 }),
              style({ transform: 'translateY(0)', offset: 1 }),
            ]))
          ])

        ]),

      ]),

    ]),

    trigger('cards1', [

      transition(':enter', [
        query('.infocard', [
          style({ position: 'relative', transform: 'translateX(-1000px)', }),
          stagger(400, [
            animate('0.5s 0.8s ease-in-out', keyframes([
              style({ transform: 'translateX(-1000px)', offset: 0 }),
              style({ transform: 'translateX(0)', offset: 1 }),
            ]))
          ])
        ]),
      ]),
    ]),

    trigger('cards2', [

      transition(':enter', [
        query('.infocard', [
          style({ position: 'relative', transform: 'translateX(1000px)', }),
          stagger(400, [
            animate('0.5s 0.8s ease-in-out', keyframes([
              style({ transform: 'translateX(1000px)', offset: 0 }),
              style({ transform: 'translateX(0)', offset: 1 }),
            ]))

          ])
        ])
      ]),
    ]),

  ]
})
export class AboutusComponent {

  constructor() { }



}
