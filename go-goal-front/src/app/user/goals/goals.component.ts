import { Component, Input, OnChanges, OnDestroy, OnInit, SimpleChanges } from '@angular/core';
import { parse } from '@fortawesome/fontawesome-svg-core';
import { BackendConnectService, userInfo } from 'src/app/backend-connect.service';
import { trigger, state, style, transition, animate, keyframes, stagger, query } from '@angular/animations';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { UserService } from '../user.service';
import { HttpClient } from '@angular/common/http';



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
        // query(':leave', [
        //   animate('0.5s', style({ opacity: 0 }))
        // ], { optional: true }),
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


    ]),

    trigger('xp', [
      transition(':enter', [
        style({ width: '0' }),
        animate('1200ms 10s ease-out', keyframes([
          style({ offset: 0, width: '0' }),
          style({ width: '*', offset: 1 }),
        ]))


      ]),

    ])

  ]
})
export class GoalsComponent implements OnInit, OnChanges {

  initList: boolean = true;
  addToList: boolean = false;
  listLoaded: boolean = false;

  @Input() goalsUncompleted: goal[] = [];
  @Input() goalsCompleted: goal[] = [];
  norm: boolean = true; deleteTime: boolean = false; editTime: boolean = false; completeGoalTime: boolean = false;
  newGoal = this.formBuilder.group({
    Title: new FormControl(""),
    Description: new FormControl(""),
    goalID: new FormControl(-1)
  });

  XP: number = 0;
  uLevel: number = 0;
  progWidth: number = 0;
  uLevelName: string = "Newbie";
  levelNames: string[] = ['Newbie', 'Goal Keeper', 'Goal Getter', 'Goal Master', 'Overachiever', 'Dream Chaser', 'Visionary', 'Legend in the Making', 'Idol', 'Ascendant', 'God of Goals'];


  constructor(private backend: BackendConnectService, private formBuilder: FormBuilder, private userService: UserService, private http: HttpClient) {
  }

  ngOnInit(): void {
    // Backend call to get goals
    this.getGoals();
    console.log('Goals loaded');

    // Get XP
    this.backend.getInfo(this.userService.getUserData().ID).subscribe((data) => {
      this.XP = data.ThisUser.XP;
      this.uLevel = Math.floor(this.XP / 500);
      if (this.uLevel >= 10) {
        this.uLevelName = this.levelNames[10];
      }
      else if (this.uLevel < 0) {
        this.uLevelName = this.levelNames[0];
      }
      else {
        this.uLevelName = this.levelNames[this.uLevel];
      }

    });
    console.log('XP loaded');




  }

  ngOnChanges(changes: SimpleChanges): void {

  }




  getGoals() {
    this.backend.getGoals(this.userService.getUserData().ID).subscribe((data) => {
      if (!data.Successful || data.ErrorExist || data == null) {
        console.log("Error getting goals (getGoals)");
      }
      else if (data.Goals.length > 0) {

        // Fill lists
        this.goalsUncompleted = [];
        data.Goals.forEach((item: any) => {
          if (item.Completed == false) {
            this.goalsUncompleted.push({ Title: item.Title, Description: item.Description, goalID: item.ID, Completed: item.Completed });
          }
          else {
            this.goalsCompleted.push({ Title: item.Title, Description: item.Description, goalID: item.ID, Completed: item.Completed });
          }

        });


      }
      else {
        this.goalsUncompleted = [];
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
      this.backend.createGoal({ Title: this.newGoal.value.Title, Description: this.newGoal.value.Description, Completed: false }, this.userService.getUserData().ID).subscribe((data) => { });

      // Push to list, delay to allow backend to update
      this.goalsUncompleted.push({ Title: this.newGoal.value.Title!, Description: this.newGoal.value.Description!, goalID: 0, Completed: false });

      // Clear form
      this.newGoal.reset();
      this.normalize();

      // Update goalID
      this.backend.getGoals(this.userService.getUserData().ID).subscribe((data) => {
        this.goalsUncompleted[this.goalsUncompleted.length - 1].goalID = data.Goals[data.Goals.length - 1].ID;
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
    else if (this.newGoal.value.goalID === -1) {
      alert("Please select a goal to edit");
      return;
    }
    console.log(this.newGoal.value.goalID);


    this.goalsUncompleted.forEach((item) => {
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
      this.goalsUncompleted.forEach((item) => {
        if (item === goal) {
          console.log("Completed goal: " + item.Title);

          // Update goal in backend
          this.backend.updateGoal({ Title: item.Title, Description: item.Description, Completed: true }, item.goalID).subscribe((data) => { });

          // Remove from uncompleted list
          this.goalsUncompleted.splice(this.goalsUncompleted.indexOf(item), 1);

          // Add to completed list
          this.goalsCompleted.push(item);

          // Update XP
          this.XP += 100;
          this.updateXP();

        }
      });

      // ADD: Complete goal in backend
    }
    else if (this.deleteTime) {
      this.goalsUncompleted.forEach((item, index) => {
        if (item === goal) {

          // Delete goal in backend
          this.backend.deleteGoals(item.goalID).subscribe((data) => { });
          this.goalsUncompleted.splice(index, 1);
          console.log("Deleted goal: " + item.Title);

        }
      });
    }
    else if (this.editTime) {
      this.goalsUncompleted.forEach((item, index) => {
        if (item === goal) {
          // Set form values
          this.newGoal.setValue({ Title: item.Title, Description: item.Description, goalID: item.goalID });

        }
      });

    }
    else { // Completed goals
      this.goalsCompleted.forEach((item) => {
        if (item === goal) {
          console.log("Deleted goal: " + item.Title);

          // Delete goal in backend
          this.backend.deleteGoals(item.goalID).subscribe((data) => { });

          // Remove from completed list
          this.goalsCompleted.splice(this.goalsCompleted.indexOf(item), 1);

        }

      });

    }

  }

  updateXP() {
    this.http.put<JSON>(`http://localhost:9000/api/users/${this.userService.getUserData().ID}/xp`, { NewXP: 100 }).subscribe((data) => {
    });
  }

}

export interface goal {
  Title: string;
  Description: string;
  goalID: number;
  Completed: boolean;
}

