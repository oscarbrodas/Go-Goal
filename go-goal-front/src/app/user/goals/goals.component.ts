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
      transition('* <=> *', [
        query(':enter', [
          style({ top: '-2000px' }),
          stagger(200, [
            animate('1250ms ease-out', keyframes([
              style({ offset: 0, position: 'relative', top: '-2000px' }),
              style({ transform: 'translateY(2000px)', offset: 0.6 }),
              style({ transform: 'translateY(1990px)', offset: 0.7 }),
              style({ transform: 'translateY(2000px)', offset: 0.8 }),
              style({ transform: 'translateY(1999px)', offset: 0.9 }),
              style({ transform: 'translateY(2000px)', offset: 1, opacity: 1 }),
            ]))
          ])
        ], { optional: true }),
        query(':leave', [
          animate('0.5s', style({ opacity: 0 }))
        ], { optional: true }),
      ]),
    ],
    ),
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
    goalID: new FormControl(0)
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
          this.userGoals.push({ Title: item.Title, Description: item.Description, goalID: item.ID, Completed: false });
          this.listLoaded = true;
        });
      }
      else {
        this.userGoals = [];
      }

    });

    this.listLoaded = true;
  }

  addGoal() {
    if (this.newGoal.value.Title == "") {
      alert("Please enter a title");
      return;
    }
    else if (this.newGoal.value.Description == "null") {
      alert("Please enter a description");
      return;
    }
    else {
      // Send goals to backend
      this.backend.createGoal({ Title: this.newGoal.value.Title, Description: this.newGoal.value.Description, Completed: false }, this.userService.getUserData().ID).subscribe((data) => {
        console.log(data);

      });

      // Push to list, delay to allow backend to update
      this.userGoals.push({ Title: this.newGoal.value.Title!, Description: this.newGoal.value.Description!, goalID: 0, Completed: false });

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




  }

  editGoal() {
    if (this.newGoal.value.Title == "" || this.newGoal.value.Description == "") {
      alert("Please enter a valid title and description");
      return;
    }

    this.userGoals.forEach((item) => {
      if (item.goalID == this.newGoal.value.goalID) {
        // Edit List
        item.Description = this.newGoal.value.Description!;
        item.Title = this.newGoal.value.Title!;

        // Edit goal in backend
        this.backend.updateGoal({ Title: item.Title, Description: item.Description }, item.goalID).subscribe((data) => { });
      }
    });

    // Clear form
    this.newGoal.reset();

    // Log
    console.log("Goal edited");

  }

  normalize() {
    this.norm = true
    this.deleteTime = false;
    this.editTime = false;
    this.completeGoalTime = false;
    this.newGoal.reset();
  }
  deleteMode() {
    this.norm = false;
    this.deleteTime = !this.deleteTime;
    this.editTime = false;
    this.completeGoalTime = false;

    if (!this.deleteTime) this.normalize();
  }
  editMode() {
    this.norm = false;
    this.deleteTime = false;
    this.editTime = !this.editTime;
    this.completeGoalTime = false;

    if (!this.editTime) this.normalize();
  }
  completeMode() {
    this.norm = false;
    this.deleteTime = false;
    this.editTime = false;
    this.completeGoalTime = !this.completeGoalTime;

    if (!this.completeGoalTime) this.normalize();


  }

  goalButton(goal: goal) {

    if (this.norm) {
      this.userGoals.forEach((item) => {
        if (item === goal) {
          item.Completed = true;
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
      this.userGoals.forEach((item, index) => {
        if (item === goal) {
          // Set form values
          this.newGoal.setValue({ Title: item.Title, Description: item.Description, goalID: item.goalID });

        }
      });

    }
    else {

    }

  }

}

export interface goal {
  Title: string;
  Description: string;
  goalID: number;
  Completed: boolean;
}

