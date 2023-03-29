import { Component } from '@angular/core';
import { parse } from '@fortawesome/fontawesome-svg-core';
import { BackendConnectService, userInfo } from 'src/app/backend-connect.service';
import { trigger, state, style, transition, animate, keyframes, stagger, query } from '@angular/animations';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { UserService } from '../user.service';
import {Router, ActivatedRoute, Params} from '@angular/router';
import {goal} from '../goals/goals.component';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent {
  constructor(private backend: BackendConnectService, private formBuilder: FormBuilder, private userService: UserService, private activatedRoute: ActivatedRoute) {
  }
  user: userInfo = {ID: 0, FirstName: "error", LastName: "error", Username: "error", Password: "Not to View", Email: "error", loggedIn: false};
  theUser: boolean = false;
  id: number = 0;
  userGoals: goal[] = [];
  topUserGoals: goal[] = [];
  theCount: number = 0;
  requested: boolean = false;
  added: boolean = false; //These booleans will be implemented when friends added into system
  ngOnInit(){
    this.activatedRoute.params.subscribe((url) => {
      console.log(url["id"]);
      this.id = url["id"];
    });
    this.backend.getInfo(this.id).subscribe((data) => {
      console.log(data)
      this.user.FirstName = data.ThisUser.FirstName;
      this.user.LastName = data.ThisUser.LastName;
      this.user.Email = data.ThisUser.Email;
      this.user.Username = data.ThisUser.Username;
      this.theUser = data.ThisUser.ID == this.userService.getUserData().ID;
    });
    this.backend.getGoals(this.id).subscribe((data) => {
      if (!data.Successful || data.ErrorExist || data == null) {
        console.log("Error getting goals (getGoals)");
      }
      else if (data.Goals.length > 0) {
        this.userGoals = [];
        data.Goals.forEach((item: any) => {
          this.userGoals.push({ Title: item.Title, Description: item.Description, goalID: item.ID, completed: false });
        });
      }
      else {
        this.userGoals = [{Title: "No goals yet", Description: "It looks like this journey is just beginning!", goalID: -1, completed: false}];
      }
      if(this.userGoals.length >= 3){
        this.topUserGoals = this.userGoals.slice(0,3);
      }else{
        this.topUserGoals = this.userGoals;
      }
      this.theCount = this.topUserGoals.length;
    })

  }
  FriendRequest(): void{
    //to add, command to send friend request when button clicked
    this.requested = true;
  }
  more(): void{
      if(this.userGoals.length-this.theCount > 3){
        this.topUserGoals = this.topUserGoals.concat(this.userGoals.slice(this.theCount,this.theCount+3))
      }else{
        this.topUserGoals = this.userGoals;
      }
      this.theCount = this.topUserGoals.length;
  }
}
