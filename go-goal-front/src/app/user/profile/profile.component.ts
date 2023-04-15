import { Component, OnChanges, OnDestroy, OnInit, SimpleChanges } from '@angular/core';
import { BackendConnectService, userInfo } from 'src/app/backend-connect.service';
import { trigger, state, style, transition, animate, keyframes, stagger, query } from '@angular/animations';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { UserService } from '../user.service';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { goal } from '../goals/goals.component';
import { HttpClient } from '@angular/common/http';

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
        style({ right: '-100%' }),
        animate('0.8s 0.35s ease', keyframes([
          style({ right: '-100%', offset: 0 }),
          style({ right: '0%', offset: 1 })
        ]))
      ]),


    ]),
  ]
})
export class ProfileComponent implements OnInit, OnChanges, OnDestroy {

  user: userInfo = { ID: 0, FirstName: "error", LastName: "error", Username: "error", Password: "Not to View", Email: "error", loggedIn: false };

  theUser: boolean = false;
  id: Number = 0;
  userGoals: goal[] = [];
  topUserGoals: goal[] = [];
  pendingIDs: number[] = [];
  outgoingIDs: number[] = [];
  viewProfileIDs: number[] = [];
  viewProfileUsernames: string[] = [];
  mutualFriends: Map<number, string> = new Map();
  pendingFriends: Map<number, string> = new Map();
  friends: number[] = [];
  theCount: number = 0;
  requested: boolean = false;
  added: boolean = false;

  editDescription: boolean = false;
  descriptionForm: FormGroup = this.formBuilder.group({
    Description: new FormControl(''),
  });

  constructor(private backend: BackendConnectService, private formBuilder: FormBuilder, private userService: UserService, private activatedRoute: ActivatedRoute, private http: HttpClient) { }


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

    // Get outgoing friend requests for logged in user
    this.backend.getOutgoingRequests(this.userService.getUserData().ID).subscribe((data) => {
      if (data.ErrorExist || data == null) {
        console.log("Error getting outgoing requests (getOutgoingRequests)");
      }
      else if (data.IDs != null && data.IDs.length > 0) {
        console.log("Outgoing requests found.");
        data.IDs.forEach((item: any) => {
          this.outgoingIDs.push(item);
        });
      }
      else {
        console.log("No outgoing requests.");
      }

      // Check if friend request has been sent already if not the user
      if (!this.theUser && this.outgoingIDs.includes(this.user.ID)) {
        this.requested = true;
      }

    });

    // Get incoming friend requests for logged in user
    this.backend.getIngoingRequests(this.userService.getUserData().ID).subscribe((data) => {
      if (data.ErrorExist || data == null) {
        console.log("Error getting ingoing requests (getIngoingRequests)");
      }
      else if (data.IDs != null && data.IDs.length > 0) {
        console.log("Incoming requests found.");
        data.IDs.forEach((item: any) => {
          this.pendingIDs.push(item);
        });
        this.getFriendsUsernames();

      }
      else {
        console.log("No incoming requests.");
      }

    });

    // Get friends of logged in user
    this.backend.getFriends(this.userService.getUserData().ID).subscribe((data) => {
      if (data.ErrorExist || data == null) {
        console.log("Error getting friends (getFriends)");
      }
      else if (data.IDs != null && data.IDs.length > 0) {
        console.log("Friends found.");
        data.IDs.forEach((item: any) => {
          this.friends.push(item);
        });

        // Check if user profile is a friend
        if (!this.theUser && this.friends.includes(this.user.ID)) {
          this.requested = false;
          this.added = true;
        }

      }
      else {
        console.log("Friends found. (Zero friends)");
      }


    });

    // Get friends of user profile and check mutual friends
    this.backend.getFriends(this.id).subscribe((data) => {
      if (data.ErrorExist || data == null) {
        console.log("Error getting friends (getFriends)");
      }
      else if (data.IDs != null && data.IDs.length > 0) {
        console.log("Friends found.");
        data.IDs.forEach((item: any) => {
          this.viewProfileIDs.push(item);
        });
        this.getMutualFriends();

      }
      else {
        console.log("Friends found. (Zero friends of user profile)");
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

  acceptRequest(id: number): void {

    this.http.put<any>(`http://localhost:9000/api/friends/acceptFriendRequest/${id}/${this.userService.getUserData().ID}`, {}).subscribe((data) => {

      if (data.ErrorExist) console.log("Error accepting friend request (acceptRequest)");
      else {
        this.pendingIDs = this.pendingIDs.filter((item) => item !== id);
        this.pendingFriends.delete(id);
        this.friends.push(id);
        console.log("Friend request accepted.");

      }


    });

  }

  declineRequest(id: number): void {

    this.http.delete<any>(`http://localhost:9000/api/friends/declineFriendRequest/${id}/${this.userService.getUserData().ID}`).subscribe((data) => {

      if (data.ErrorExist) console.log("Error declining friend request (declineRequest)");
      else {
        this.pendingIDs = this.pendingIDs.filter((item) => item !== id);
        this.pendingFriends.delete(id);
        console.log("Friend request declined.");

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

  // Get friend usernames from IDs
  getFriendsUsernames() {

    this.pendingIDs.map((pending) => {

      this.http.get<any>(`http://localhost:9000/api/users?id=${pending}`).subscribe((data) => {
        if (data.ErrorExist) console.log("Error getting friend username (getFriendsUsernames)");

        this.pendingFriends.set(pending, data.ThisUser.Username);

      });

    });

  }

  // Get usernames of mutual friends
  getMutualFriends() {

    this.viewProfileIDs.map((friend) => {

      this.http.get<any>(`http://localhost:9000/api/users?id=${friend}`).subscribe((data) => {
        if (data.ErrorExist) console.log("Error getting friend username (getMutualUsernames)");

        if (this.friends.includes(friend)) this.mutualFriends.set(friend, data.ThisUser.Username);


      });

    });

  }


}
