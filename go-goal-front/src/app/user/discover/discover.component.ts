import { Component, OnChanges, OnInit } from '@angular/core';
import { BackendConnectService, userInfo } from '../../backend-connect.service';
import { LoginService } from 'src/app/login-page/login.service';
import { UserService } from '../user.service';
import { HttpClient } from '@angular/common/http';
import { FormControl, FormGroup } from '@angular/forms';
import { KeyValue } from '@angular/common';
import { Router } from '@angular/router';
import { trigger, state, style, transition, animate, keyframes, stagger, query, group } from '@angular/animations';


@Component({
  selector: 'app-discover',
  templateUrl: './discover.component.html',
  styleUrls: ['./discover.component.css'],
  animations: [

    trigger('tools', [
      transition(':enter', [
        style({ left: '-50%' }),
        animate('0.6s ease', style({ left: '0%' }))
      ]),
    ]),


    trigger('banner', [
      transition(':enter', [
        style({ width: '0%', content: '', border: 'none' }),
        group([
          style({ width: '0%', content: '', border: 'none' }),
          animate('0.4s 0.5s ease', keyframes([
            style({ width: '0%', border: '*' }),
            style({ width: '100%', }),
          ])),


          query('p', [
            style({ opacity: 0 }),
            animate('0.6s 0.7s ease', style({ opacity: 1 }))
          ]),

        ])

      ]),

    ]),

    trigger('friends', [
      transition(':enter', [
        style({ height: '0', content: '', background: 'none', border: 'none' }),
        animate('0.5s 0.9s ease', keyframes([
          style({ height: '0', background: '*' }),
          style({ height: '*', border: '*' }),

        ])),
      ]),


    ]),

    trigger('button', [
      transition(':enter', [
        style({ opacity: 0 }),
        animate('0.5s 1.2s ease', style({ opacity: 1 }))


      ]),

    ])


  ]
})
export class DiscoverComponent implements OnInit, OnChanges {

  friendRequested: boolean = false;
  removeRequested: boolean = false;
  index: number = 0;
  user?: userInfo;
  userFriendsIDs: number[] = [];
  userSearched: userInfo[] = [];
  userFriends: Map<number, string> = new Map([]);
  outgoingFriendRequests: number[] = [];
  incomingFriendRequests: number[] = [];
  standard: string = "Click on a Profile or search for an ID to display their information."

  searchForm = new FormGroup({
    search: new FormControl(''),
  });
  friendForm = new FormGroup({
    friendUsername: new FormControl('[ USER ]'),
    friend: new FormControl('-1'),
  });

  IDSearch: number = -1;
  message: string = "";

  constructor(private loginService: LoginService, private userService: UserService, private backend: BackendConnectService, private http: HttpClient, private router: Router) {


  }

  ngOnInit(): void {

    // Grab user's info
    this.user = this.userService.getUserData();

    // Grab friend's IDs
    this.backend.getFriends(this.user?.ID!).subscribe((data) => {

      if (data.ErrorExist) {
        console.log('Error: Could not get friends IDs.');
      } else if (data.IDs === null) {
        this.userFriendsIDs = [];
        console.log('No friends found.');
      } else {
        console.log('Friends IDs Loaded.');
        this.userFriendsIDs = data.IDs;

        // For each friend's ID, grab their username
        this.getFriendsUsernames();

      }

    });


    // Get outgoing and incoming friend requests
    this.getRequests();

  }

  ngOnChanges(): void {


  }

  // Search for a user among database
  search(): void {
    // Backend call
    this.http.get<any>(`http://localhost:9000/api/friends/search/${this.searchForm.value.search}`, this.backend.httpOptions).subscribe((data) => {
    console.log(data)
    this.userSearched= []
    if (data.ErrorExist || data.Users === null) {
        this.friendForm.setValue({
          friendUsername: 'No Users Found',
          friend: `-1`
        });
        this.standard = "No users with that username were found, please try again"
        console.log('Error: Could not get friend data.');
      } else { // User found, set form values
        for(let entry=0;entry<data.Users.length;entry++){
          this.userSearched.push({"ID": data.Users[entry].ID,
            "loggedIn": false,
            "Username": data.Users[entry].Username,
            "FirstName": data.Users[entry].FirstName,
            "LastName": data.Users[entry].LastName,
            "Email": data.Users[entry].Email,
            "Password": data.Users[entry].Password})
        }
        console.log(this.userSearched)
        this.friendForm.setValue({
          friendUsername: `Username: ${data.Users[0].Username}`,
          friend: `Name: ${data.Users[0].FirstName + ' ' + data.Users[0].LastName}`
        });
        this.message = "Add Friend";
        this.IDSearch = data.Users[0].ID;
        this.index = 0;
        // Check if friend request has already been sent
        if(this.userFriendsIDs.includes(this.userSearched[this.index].ID)){
          this.message = "Remove Friend"
        }
        if (this.outgoingFriendRequests.includes(this.IDSearch)) {
          this.message = "Friend Request Sent";
          this.friendRequested = true;
        } else {
          this.friendRequested = false;
        }
        console.log(this.userSearched)
      }

    });
    this.friendRequested = false;
    this.removeRequested = false;


  }
  prev(){
    this.index = (this.index+this.userSearched.length-1)%this.userSearched.length
    this.friendForm.setValue({
      friendUsername: `Username: ${this.userSearched[this.index].Username}`,
      friend: `Name: ${this.userSearched[this.index].FirstName + ' ' + this.userSearched[this.index].LastName}`
    });
    this.friendRequested = false;
    this.removeRequested = false;
    this.message = "Add Friend";
    this.IDSearch = this.userSearched[this.index].ID;
    // Check if friend request has already been sent
    if(this.userFriendsIDs.includes(this.userSearched[this.index].ID)){
      this.message = "Remove Friend"
    }
    if (this.outgoingFriendRequests.includes(this.IDSearch)) {
      this.message = "Friend Request Sent";
      this.friendRequested = true;
    } else {
      this.friendRequested = false;
    }
  }
  next(){
    this.index = (this.index+1)%this.userSearched.length
    this.friendForm.setValue({
      friendUsername: `Username: ${this.userSearched[this.index].Username}`,
      friend: `Name: ${this.userSearched[this.index].FirstName + ' ' + this.userSearched[this.index].LastName}`
    });
    this.friendRequested = false;
    this.removeRequested = false;
    this.message = "Add Friend";
    this.IDSearch = this.userSearched[this.index].ID;
    // Check if friend request has already been sent
    if(this.userFriendsIDs.includes(this.userSearched[this.index].ID)){
      this.message = "Remove Friend"
    }
    if (this.outgoingFriendRequests.includes(this.IDSearch)) {
      this.message = "Friend Request Sent";
      this.friendRequested = true;
    } else {
      this.friendRequested = false;
    }
  }
  // Navigate to profile of user selected
  viewProfile(): void {
    this.router.navigate([`/user/${this.IDSearch}/profile`]);
  }

  // Set friend profile based friend selected
  setFriendProfile(friend: KeyValue<number, string>) {
    this.friendForm.setValue({
      friendUsername: `Username: ${friend.value}`,
      friend: `Friend ID: ${friend.key.toString()}`,
    });
    this.IDSearch = friend.key;
    this.message = "Remove Friend";
    this.friendRequested = false;
    this.removeRequested = false;
  }


  // Add or remove friend based on if profile is already a friend or not
  addOrRemove(): void {
    if (this.message == "Add Friend") { // Add friend section from search
      console.log('Adding friend...');
      this.friendRequested = true;
      this.removeRequested = false;
      this.message = "Friend Request Sent";

      // Backend call to send friend request
      this.http.post<any>(`http://localhost:9000/api/friends/sendFriendRequest/${this.user?.ID}/${this.IDSearch}`, {}).subscribe((data) => {
        if (data.ErrorExist || data.Successful === false) {
          console.log('Error: Could not add friend.');
        } else {
          console.log('Friend Request Sent.');
        }

      });


    } else if (this.message == "Remove Friend"){ // Remove friend section from profile section
      console.log('Removing friend...');
      this.message = "Removed Friend";
      this.friendRequested = false;
      this.removeRequested = true;

      // Remove friend from map and backend
      this.userFriends.delete(this.IDSearch);
      this.userFriendsIDs = this.userFriendsIDs.filter((ele)=> ele !== this.IDSearch);
      console.log(this.IDSearch)
      this.http.delete<any>(`http://localhost:9000/api/friends/removeFriend/${this.user?.ID}/${this.IDSearch}`, {}).subscribe((data) => {

        if (data.ErrorExist) {
          console.log('Error: Could not remove friend.');
        } else {
          console.log('Friend removed.');
        }

      });

    }

  }

  // Get friend usernames from IDs
  getFriendsUsernames(): boolean {

    let res: boolean = true;
    this.userFriendsIDs.map((friend) => {

      this.http.get<any>(`http://localhost:9000/api/users?id=${friend}`).subscribe((data) => {
        if (data.ErrorExist) {
          res = false;
          console.log('Error: Could not get friend data.');

        }
        else {
          this.userFriends.set(friend, data.ThisUser.Username);
        }


      });

    });
    return res;
  }


  // Get friend requests in and out
  getRequests(): void {


    // Get outgoing friend requests
    this.http.get<any>(`http://localhost:9000/api/friends/getOutgoingFriendRequests/${this.user?.ID}`).subscribe((data) => {
      if (data.ErrorExist) {
        console.log('Error: Could not get outgoing friend requests.');
      } else if (data.IDs === null) {
        console.log('Got outgoing friend requests.');
        this.outgoingFriendRequests = [];
      } else {
        console.log('Got outgoing friend requests.');
        this.outgoingFriendRequests = data.IDs;
      }

    });

    // Get incoming friend requests
    this.http.get<any>(`http://localhost:9000/api/friends/getIngoingFriendRequests/${this.user?.ID}`).subscribe((data) => {
      if (data.ErrorExist) {
        console.log('Error: Could not get incoming friend requests.');
      } else if (data.IDs === null) {
        console.log('Got incoming friend requests.');
        this.incomingFriendRequests = [];
      } else {
        console.log('Got incoming friend requests.');
        this.incomingFriendRequests = data.IDs;
      }


    });

  }

}
