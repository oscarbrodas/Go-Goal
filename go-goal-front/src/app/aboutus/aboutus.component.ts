import { Component } from '@angular/core';
import { trigger, state, style, transition, animate, keyframes, stagger, query } from '@angular/animations';

@Component({
  selector: 'app-aboutus',
  templateUrl: './aboutus.component.html',
  styleUrls: ['./aboutus.component.css'],
  animations: [
    trigger('tags', [

      transition(':enter', [
        query('div', [
          style({ position: 'relative', top: '-200%' }),
          stagger(100, [
            animate('0.5s ease-in-out', style({ position: 'relative', top: '0%' }))
          ])

        ]),

      ]),

    ]),
  ]
})
export class AboutusComponent {

  constructor() { }



}
