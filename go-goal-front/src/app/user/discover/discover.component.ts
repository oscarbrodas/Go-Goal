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
        style({ width: '0%', content: '', border: 'none', top: '-100%' }),
        group([
          style({ width: '0%', content: '', border: 'none', top: '-50%' }),
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
        group([
          animate('0.5s 0.9s ease', keyframes([
            style({ height: '0', background: '*' }),
            style({ height: '*', border: '*' }),

          ])),

          query('button', [
            style({ opacity: 0 }),
            animate('0.6s 1.3s ease', keyframes([
              style({ opacity: 0 }),
              style({ opacity: 1 }),
            ])),

          ]),


        ])

      ]),



    ]),


  ]
})
export class DiscoverComponent implements OnInit, OnChanges {

  friendRequested: boolean = false;
  removeRequested: boolean = false;

  user?: userInfo;
  userFriendsIDs: number[] = [];
  userFriends: Map<number, string> = new Map([
    [3, 'Friend 1']]);
  userSearches: Map<number, string> = new Map([[3, 'Search 1']]);
  outgoingFriendRequests: number[] = [];
  incomingFriendRequests: number[] = [];

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
        this.userFriendsIDs = data.IDs;
        console.log('Friends IDs Loaded.');
      }

    });

    // For each friend's ID, grab their username
    const fr = async (): Promise<boolean> => {
      return this.getFriendsUsernames();
    };

    fr().then((res) => {

      if (res) {
        console.log('Friends Loaded.');


      } else {
        console.log('Error: Could not get friends usernames.');
      }

    });

    // Get outgoing and incoming friend requests
    this.getRequests();

  }

  ngOnChanges(): void {


  }

  // Search for a user among database
  search(): void {
    // Validate input
    let s: string = this.searchForm.value.search?.toString()!;
    if (Number.isNaN(Number(s)) || Number(s) <= 0 || Number(s) == this.user?.ID) {
      alert('Please enter a valid ID to search for.')
      return;
    }
    else if (this.userFriendsIDs.includes(Number(s))) {
      alert('This user is already your friend.')
      return;
    }

    // Backend call
    this.http.get<any>(`http://localhost:9000/api/users?id=${this.searchForm.value.search}`).subscribe((data) => {
      if (data.ErrorExist) {
        this.friendForm.setValue({
          friendUsername: 'User Not Found',
          friend: `User Not Found`
        });
        console.log('Error: Could not get friend data.');
      } else { // User found, set form values
        this.friendForm.setValue({
          friendUsername: `Username: ${data.ThisUser.Username}`,
          friend: `Name: ${data.ThisUser.FirstName + ' ' + data.ThisUser.LastName}`
        });
        this.message = "Add Friend";
        this.IDSearch = data.ThisUser.ID;

        // Check if friend request has already been sent
        if (this.outgoingFriendRequests.includes(this.IDSearch)) {
          this.message = "Friend Request Sent";
          this.friendRequested = true;
        }

      }

    });

    this.removeRequested = false;


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

      // Backend call to add friend
      this.http.post<any>(`http://localhost:9000/api/friends/sendFriendRequest/${this.user?.ID}/${this.IDSearch}`, {}).subscribe((data) => {
        if (data.ErrorExist || data.Successful === false) {
          console.log('Error: Could not add friend.');
        } else {
          console.log('Friend added.');
        }

      });


    } else { // Remove friend section from profile section
      console.log('Removing friend...');
      this.message = "Removed Friend";
      this.friendRequested = false;
      this.removeRequested = true;

      // Remove friend from map and backend
      this.userFriends.delete(this.IDSearch);

    }

  }

  // Get friend usernames from IDs
  getFriendsUsernames(): boolean {

    let res = true;
    this.userFriendsIDs.map((friend) => {

      this.http.get<any>(`http://localhost:9000/api/users?id=${friend}`).subscribe((data) => {
        if (data.ErrorExist) res = false;
        console.log(data);

        this.userFriends.set(friend, data.ThisUser.Username);

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
    this.http.get<any>(`http://localhost:9000/api/friends/getIncomingFriendRequests/${this.user?.ID}`).subscribe((data) => {
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
