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
          animate(250, keyframes([
            style({ offset: 0 }),
            style({ opacity: '0.5', offset: 0.5 }),
            style({ offset: 1, opacity: 0 }),
          ]))
        ])
      ])
    ],
    ),
    trigger('goal', [
      transition(':enter', [
        style({ opacity: 1, transform: 'translateY(-500px)' }),
        animate(650, keyframes([
          style({ offset: 0, position: 'relative', top: '-500px' }),
          style({ transform: 'translateY(500px)', offset: 0.6 }),
          style({ transform: 'translateY(490px)', offset: 0.7 }),
          style({ transform: 'translateY(490px)', offset: 0.9 }),
          style({ transform: 'translateY(500px)', offset: 1, opacity: 1 }),
        ]))
      ]),
    ]),
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

  initList: boolean = true;
  addToList: boolean = false;

  userGoals: goal[];
  norm: boolean = true; deleteTime: boolean = false; editTime: boolean = false; completeGoalTime: boolean = false;
  newGoal = this.formBuilder.group({
    Title: new FormControl(""),
    Description: new FormControl(""),
  });


  constructor(private backend: BackendConnectService, private formBuilder: FormBuilder) {
    this.userGoals = [
      { Title: "Goal 1", Description: "Fix Bike", UserID: 2, },
      { Title: "Goal 2", Description: "Ride Bike", UserID: 2, },
      { Title: "Goal 3", Description: "Ride Wife", UserID: 2, },

    ];
  }

  ngOnInit(): void {
    this.userGoals = [
      { Title: "Goal 1", Description: "Fix Bike", UserID: 2, },
      { Title: "Goal 2", Description: "Ride Bike", UserID: 2, },
      { Title: "Goal 3", Description: "Ride Wife", UserID: 2 },

    ];
  }


  // getGoals(): JSON {

  //   return JSON.parse(this.backend.getGoals());
  // }

  addGoal() {
    if (this.newGoal.value.Title == "") {
      alert("Please enter a title");
      return;
    }
    if (this.newGoal.value.Description == "") {
      alert("Please enter a description");
      return;
    }

    this.userGoals.push({ Title: this.newGoal.value.Title!, Description: this.newGoal.value.Description!, UserID: 2, completed: false });

    // add goal to backend

  }
  normalize() {
    this.norm = true
    this.deleteTime = false;
    this.editTime = false;
    this.completeGoalTime = false;
  }
  deleteGoal() {
    this.norm = false;
    this.deleteTime = true;
    this.editTime = false;
    this.completeGoalTime = false;
  }
  editGoal() {
    this.norm = false;
    this.deleteTime = false;
    this.editTime = true;
    this.completeGoalTime = false;
  }
  completeGoal(goal: goal) {
    this.norm = false;
    this.deleteTime = false;
    this.editTime = false;
    this.completeGoalTime = true;
  }

  goalButton(goal: goal) {

    if (this.norm) {
      this.userGoals.forEach((item, index) => {
        if (item === goal) item.completed = true;
      });
    }
    else if (this.deleteTime) {
    }
    else if (this.editTime) {
    }
    else {

    }

  }

}

export interface goal {
  Title: string;
  Description: string;
  UserID: number;
  completed?: boolean;
}

