import { Component, OnChanges, OnInit } from '@angular/core';
import { BackendConnectService, userInfo } from '../../backend-connect.service';
import { LoginService } from 'src/app/login-page/login.service';
import { UserService } from '../user.service';
import { HttpClient } from '@angular/common/http';
import { FormControl, FormGroup } from '@angular/forms';
import { KeyValue } from '@angular/common';
import { Router } from '@angular/router';

@Component({
  selector: 'app-discover',
  templateUrl: './discover.component.html',
  styleUrls: ['./discover.component.css']
})
export class DiscoverComponent implements OnInit, OnChanges {

  friends: boolean = true;
  users: boolean = false;

  user?: userInfo;
  userFriendsIDs: number[] = [3];
  userFriends: Map<number, string> = new Map([[3, 'Friend 1']]);
  userSearches: Map<number, string> = new Map([[3, 'Search 1']]);
  searchForm = new FormGroup({
    search: new FormControl(''),
  });
  friendForm = new FormGroup({
    friendUsername: new FormControl('[ USER ]'),
    friend: new FormControl('-1'),
  });
  IDSearch: number = -1;
  message: string = "";

  constructor(private loginService: LoginService, private userService: UserService, private http: HttpClient, private router: Router) {


  }

  ngOnInit(): void {

    // Grab user's info
    this.user = this.userService.getUserData();

    // Grab friend's IDs
    this.userFriendsIDs = this.loginService.friends;

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




  }

  ngOnChanges(): void {


  }

  search(): void {
    let s: string = this.searchForm.value.search?.toString()!;
    if (Number.isNaN(Number(s)) || Number(s) <= 0) {
      alert('Please enter a valid ID to search for.')
      return;
    }


    this.http.get<any>(`http://localhost:9000/api/users?id=${this.searchForm.value.search}`).subscribe((data) => {
      if (data.ErrorExist) {
        this.friendForm.setValue({
          friendUsername: 'User Not Found',
          friend: `User Not Found`
        });
        console.log('Error: Could not get user.');
      } else {
        this.friendForm.setValue({
          friendUsername: `Username: ${data.ThisUser.Username}`,
          friend: `Name: ${data.ThisUser.FirstName + ' ' + data.ThisUser.LastName}`
        });
        this.message = "Add Friend";
        this.IDSearch = data.ThisUser.ID;
      }

    });



  }

  viewProfile(): void {
    this.router.navigate([`/user/${this.IDSearch}/profile`]);
  }

  setFriendProfile(friend: KeyValue<number, string>) {
    this.friendForm.setValue({
      friendUsername: `Username: ${friend.value}`,
      friend: `Friend ID: ${friend.key.toString()}`,
    });
    this.IDSearch = friend.key;
    this.message = "Remove Friend";


  }

  addFriend(): void {

  }

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

}
