import { Component, OnChanges, OnDestroy, OnInit, SimpleChanges } from '@angular/core';
import { BackendConnectService, userInfo } from 'src/app/backend-connect.service';
import { trigger, state, style, transition, animate, keyframes, stagger, query } from '@angular/animations';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { UserService } from '../user.service';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { goal } from '../goals/goals.component';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css'],
  animations: [
    trigger('details', [

      transition(':enter', [
        style({ top: '100%' }),
        animate('0.8s ease', keyframes([
          style({ top: '-100%', offset: 0 }),
          style({ top: '0%', offset: 1 })
        ]))
      ]),

    ]),

    trigger('goals', [

      transition(':enter', [
        style({ top: '-100%' }),
        animate('0.8s 0.2s ease', keyframes([
          style({ top: '-100%', offset: 0 }),
          style({ top: '0%', offset: 1 })
        ]))
      ]),

    ]),

    trigger('friends', [
      transition(':enter', [
        style({ opacity: 0, transform: 'translateX(200px)' }),
        animate(500)
      ]),
      transition(':leave', [
        style({ opacity: 1, transform: 'translateX(800px)' }),
        animate(500)
      ])
    ]),
  ]
})
export class ProfileComponent implements OnInit, OnChanges, OnDestroy {

  user: userInfo = { ID: 0, FirstName: "error", LastName: "error", Username: "error", Password: "Not to View", Email: "error", loggedIn: false };

  theUser: boolean = false;
  id: Number = 0;
  userGoals: goal[] = [];
  topUserGoals: goal[] = [];
  pendingFriends: number[] = [];
  friends: number[] = [];
  theCount: number = 0;
  requested: boolean = false;
  added: boolean = false;

  editDescription: boolean = false;
  descriptionForm: FormGroup = this.formBuilder.group({
    Description: new FormControl(''),
  });

  constructor(private backend: BackendConnectService, private formBuilder: FormBuilder, private userService: UserService, private activatedRoute: ActivatedRoute) { }


  ngOnChanges(changes: SimpleChanges): void {

  }
  ngOnDestroy(): void {


  }

  ngOnInit() {

    this.activatedRoute.params.subscribe((url) => {
      //console.log(url["id"]);
      this.id = url["id"];
    });

    // Get user info for this profile
    this.backend.getInfo(this.id).subscribe((data) => {
      // console.log(data)
      this.user.ID = data.ThisUser.ID;
      this.user.FirstName = data.ThisUser.FirstName;
      this.user.LastName = data.ThisUser.LastName;
      this.user.Email = data.ThisUser.Email;
      this.user.Username = data.ThisUser.Username;
      this.user.Description = data.ThisUser.Description;
      this.user.XP = data.ThisUser.XP;
      this.theUser = data.ThisUser.ID == this.userService.getUserData().ID;
    });

    // Get user goals for this profile
    this.backend.getGoals(this.id).subscribe((data) => {

      if (!data.Successful || data.ErrorExist || data == null) {
        console.log("Error getting goals (getGoals)");
      }
      else if (data.Goals.length > 0) {

        this.userGoals = [];
        data.Goals.forEach((item: any) => {

          // Push uncompleted goals to userGoals
          if (item.Completed === false) {
            this.userGoals.push({ Title: item.Title, Description: item.Description, goalID: item.ID, Completed: item.Completed });
          }

        });
      } else {
        this.userGoals = [{ Title: "No goals yet", Description: "It looks like this journey is just beginning!", goalID: -1, Completed: false }];
      }

      // Slice top 3 goals
      if (this.userGoals.length >= 3) {
        this.topUserGoals = this.userGoals.slice(0, 3);
      } else {
        this.topUserGoals = this.userGoals;
      }
      this.theCount = this.topUserGoals.length;
    })

    // Get pending friend requests
    this.backend.getOutgoingRequests(this.userService.getUserData().ID).subscribe((data) => {
      if (data.ErrorExist || data == null) {
        console.log("Error getting outgoing requests (getOutgoingRequests)");
      }
      else if (data.IDs != null && data.IDs.length > 0) {
        console.log("Outgoing requests found.");
        data.IDs.forEach((item: any) => {
          this.pendingFriends.push(item);
        });

        // Check if friend request has been sent already if not the user
        if (!this.theUser && this.pendingFriends.includes(this.user.ID)) {
          this.requested = true;
        }

      }
      else {
        console.log("No outgoing requests.");
      }

    });

    // Get friends
    this.backend.getFriends(this.userService.getUserData().ID).subscribe((data) => {
      if (data.ErrorExist || data == null) {
        console.log("Error getting friends (getFriends)");
      }
      else if (data.IDs != null && data.IDs.length > 0) {
        data.IDs.forEach((item: any) => {
          this.friends.push(item);
        });

        // Check if user profile is a friend
        if (!this.theUser && this.friends.includes(this.user.ID)) {
          this.added = true;
        }

      }
      else {
        console.log("No friends.");
      }


    });



  }

  updateDescription() {

    this.user.Description = this.descriptionForm.value.Description;
    // Backend call to update description
    this.backend.updateDescription(this.user.ID, this.descriptionForm.value.Description).subscribe((data) => {
      if (!data.Successful || data.ErrorExist || data == null) {
        console.log("Error updating description (updateDescription)");
      }
      else {
        console.log("Description updated.");
        this.editDescription = false;
      }
    });

  }

  FriendRequest(): void {

    this.requested = true;
    this.backend.sendFriendRequest(this.userService.getUserData().ID, this.user.ID).subscribe((data) => {
      if (!data.Successful || data.ErrorExist || data == null) {
        console.log("Error sending friend request (sendFriendRequest)");
      }
      else {
        console.log("Friend request sent.");
      }
    });


  }

  more(): void {
    if (this.userGoals.length - this.theCount > 3) {
      this.topUserGoals = this.topUserGoals.concat(this.userGoals.slice(this.theCount, this.theCount + 3))
    } else {
      this.topUserGoals = this.userGoals;
    }
    this.theCount = this.topUserGoals.length;
  }


}
