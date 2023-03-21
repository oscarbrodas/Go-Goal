import { Component } from '@angular/core';
import { parse } from '@fortawesome/fontawesome-svg-core';
import { BackendConnectService, userInfo } from 'src/app/backend-connect.service';
import { trigger, state, style, transition, animate, keyframes } from '@angular/animations';


@Component({
  selector: 'app-goals',
  templateUrl: './goals.component.html',
  styleUrls: ['./goals.component.css'],
  animations: [
    trigger('goals', [
      transition('void => *', [
        animate(400, keyframes([
          style({ offset: 0, position: 'relative', top: '-500px' }),
          style({ transform: 'translateY(500px)', offset: 0.6 }),
          style({ transform: 'translateY(480px)', offset: 0.7 }),
          style({ transform: 'translateY(480px)', offset: 0.8 }),
          style({ transform: 'translateY(500px)', offset: 1 }),
        ]))
      ]),
    ])
  ]
})
export class GoalsComponent {

  userGoals: goal[];

  constructor(private backend: BackendConnectService) {
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

}

export interface goal {
  Title: string;
  Description: string;
  UserID: number;
  User: userInfo;
}

