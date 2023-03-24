import { Component } from '@angular/core';
import { parse } from '@fortawesome/fontawesome-svg-core';
import { BackendConnectService, userInfo } from 'src/app/backend-connect.service';
import { trigger, state, style, transition, animate, keyframes, stagger, query } from '@angular/animations';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';


@Component({
  selector: 'app-goals',
  templateUrl: './goals.component.html',
  styleUrls: ['./goals.component.css'],
  animations: [
    trigger('goals', [
      transition('void => *', [
        query(':enter', [
          style({ opacity: 1, transform: 'translateY(-500px)' }),
          stagger(200, [
            animate(650, keyframes([
              style({ offset: 0, position: 'relative', top: '-500px' }),
              style({ transform: 'translateY(500px)', offset: 0.6 }),
              style({ transform: 'translateY(490px)', offset: 0.7 }),
              style({ transform: 'translateY(490px)', offset: 0.9 }),
              style({ transform: 'translateY(500px)', offset: 1, opacity: 1 }),
            ]))
          ])
        ])
      ]),
      transition(':leave', [
        query(':leave', [
          stagger(250, [
            animate(250, keyframes([
              style({ offset: 0 }),
              style({ transform: 'translateX(2000px)', offset: 1 }),
            ]))
          ])
        ])
      ])
    ],
    ), // end of goals trigger
    trigger('sidebar', [

      transition(':enter', [
        animate('1200ms ease-out', keyframes([
          style({ offset: 0, left: '-500px' }),
          style({ left: '-400px', offset: 0.45 }),
          style({ left: '*', offset: 1 }),
        ]))
      ]),
      transition(':leave', [
        animate('1200ms ease-out', keyframes([
          style({ offset: 0, left: '*' }),
          style({ left: '-400px', offset: 0.45 }),
          style({ left: '-500px', offset: 1 }),
        ]))
      ])
    ])
  ]
})
export class GoalsComponent {

  userGoals: goal[];
  newGoalTime: boolean = false;
  completeGoalTime: boolean = false;
  newGoal = this.formBuilder.group({
    Title: new FormControl(""),
    Description: new FormControl(""),
  });


  constructor(private backend: BackendConnectService, private formBuilder: FormBuilder) {
    this.userGoals = [
      { Title: "Goal 1", Description: "Fix Bike", UserID: 2, User: { FirstName: "Test", LastName: "Test", Email: "Test", Password: "Test", Username: "test", loggedIn: true, ID: 2 } },
      { Title: "Goal 2", Description: "Ride Bike", UserID: 2, User: { FirstName: "Test", LastName: "Test", Email: "Test", Password: "Test", Username: "test", loggedIn: true, ID: 2 } },
      { Title: "Goal 3", Description: "Ride Wife", UserID: 2, User: { FirstName: "Test", LastName: "Test", Email: "Test", Password: "Test", Username: "test", loggedIn: true, ID: 2 } },

    ];
  }

  ngOnInit(): void {
    this.userGoals = [
      { Title: "Goal 1", Description: "Fix Bike", UserID: 2, User: { FirstName: "Test", LastName: "Test", Email: "Test", Password: "Test", Username: "test", loggedIn: true, ID: 2 } },
      { Title: "Goal 2", Description: "Ride Bike", UserID: 2, User: { FirstName: "Test", LastName: "Test", Email: "Test", Password: "Test", Username: "test", loggedIn: true, ID: 2 } },
      { Title: "Goal 3", Description: "Ride Wife", UserID: 2, User: { FirstName: "Test", LastName: "Test", Email: "Test", Password: "Test", Username: "test", loggedIn: true, ID: 2 } },

    ];
  }

  // getGoals(): JSON {

  //   return JSON.parse(this.backend.getGoals());
  // }

  addGoal() {


    this.newGoalTime = false;
  }

}

export interface goal {
  Title: string;
  Description: string;
  UserID: number;
  User: userInfo;
}

