import { Component } from '@angular/core';
import { parse } from '@fortawesome/fontawesome-svg-core';
import { BackendConnectService, userInfo } from 'src/app/backend-connect.service';
import { trigger, state, style, transition, animate, keyframes, stagger, query } from '@angular/animations';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { UserService } from '../user.service';



@Component({
  selector: 'app-goals',
  templateUrl: './goals.component.html',
  styleUrls: ['./goals.component.css'],
  animations: [
    trigger('goals', [
      transition('void => *', [
        query(':enter', [
          style({ opacity: 1, transform: 'translateY(-2500px)' }),
          stagger(200, [
            animate(650, keyframes([
              style({ offset: 0, position: 'relative', top: '-2500px' }),
              style({ transform: 'translateY(2500px)', offset: 0.6 }),
              style({ transform: 'translateY(2490px)', offset: 0.7 }),
              style({ transform: 'translateY(2490px)', offset: 0.9 }),
              style({ transform: 'translateY(2500px)', offset: 1, opacity: 1 }),
            ]))
          ])
        ], { optional: true })
      ]),
      transition(':leave', [
        query(':leave', [
          animate(250, keyframes([
            style({ offset: 0 }),
            style({ opacity: '0.5', offset: 0.5 }),
            style({ offset: 1, opacity: 0 }),
          ]))
        ], { optional: true }),
      ])
    ],
    ),
    trigger('goal', [
      transition(':enter', [
        style({ opacity: 1, transform: 'translateY(-2500px)' }),
        animate(650, keyframes([
          style({ offset: 0, position: 'relative', top: '-2500px' }),
          style({ transform: 'translateY(2500px)', offset: 0.6 }),
          style({ transform: 'translateY(2490px)', offset: 0.7 }),
          style({ transform: 'translateY(2490px)', offset: 0.9 }),
          style({ transform: 'translateY(2500px)', offset: 1, opacity: 1 }),
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
  listLoaded: boolean = false;

  userGoals: goal[] = [];
  norm: boolean = true; deleteTime: boolean = false; editTime: boolean = false; completeGoalTime: boolean = false;
  newGoal = this.formBuilder.group({
    Title: new FormControl(""),
    Description: new FormControl(""),
  });


  constructor(private backend: BackendConnectService, private formBuilder: FormBuilder, private userService: UserService) {
  }

  ngOnInit(): void {
    // Backend call to get goals
    this.getGoals();
  }

  getGoals() {
    this.backend.getGoals(this.userService.getUserData().ID).subscribe((data) => {
      if (!data.Successful || data.ErrorExist || data == null) {
        console.log("Error getting goals (getGoals)");
      }
      else if (data.Goals.length > 0) {
        this.userGoals = [];
        data.Goals.forEach((item: any) => {
          this.userGoals.push({ Title: item.Title, Description: item.Description, goalID: item.ID, completed: false });
          this.listLoaded = true;
        });
      }
      else {
        this.userGoals = [];
      }

    });


  }

  addGoal() {
    if (this.newGoal.value.Title == "") {
      alert("Please enter a title");
      return;
    }
    if (this.newGoal.value.Description == "") {
      alert("Please enter a description");
      return;
    }

    // Send goals to backend
    this.backend.createGoal({ Title: this.newGoal.value.Title, Description: this.newGoal.value.Description }, this.userService.getUserData().ID).subscribe((data) => { });

    // Push to list, delay to allow backend to update
    this.userGoals.push({ Title: this.newGoal.value.Title!, Description: this.newGoal.value.Description!, goalID: 0, completed: false });

    // Clear form
    this.newGoal.reset();
    this.normalize();

    // Update goalID
    this.backend.getGoals(this.userService.getUserData().ID).subscribe((data) => {
      this.userGoals[this.userGoals.length - 1].goalID = data.Goals[data.Goals.length - 1].ID;
    });


    // Log
    console.log('New goal added');
  }
  normalize() {
    this.norm = true
    this.deleteTime = false;
    this.editTime = false;
    this.completeGoalTime = false;
  }
  deleteMode() {
    this.norm = false;
    this.deleteTime = true;
    this.editTime = false;
    this.completeGoalTime = false;
  }
  editMode() {
    this.norm = false;
    this.deleteTime = false;
    this.editTime = true;
    this.completeGoalTime = false;
  }
  completeMode() {
    this.norm = false;
    this.deleteTime = false;
    this.editTime = false;
    this.completeGoalTime = true;
  }

  goalButton(goal: goal) {

    if (this.norm) {
      this.userGoals.forEach((item) => {
        if (item === goal) {
          item.completed = true;
        }
      });

      // ADD: Complete goal in backend
    }
    else if (this.deleteTime) {
      this.userGoals.forEach((item, index) => {
        if (item === goal) {
          this.backend.deleteGoals(item.goalID).subscribe((data) => { });
          this.userGoals.splice(index, 1);
          console.log("Deleted goal: " + item.Title);

        }
      });
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
  goalID: number;
  completed?: boolean;
}

