import { Component, OnChanges, OnInit } from '@angular/core';
import { BackendConnectService, userInfo } from '../../backend-connect.service';
import { LoginService } from 'src/app/login-page/login.service';
import { UserService } from '../user.service';
import { HttpClient } from '@angular/common/http';
import { FormControl, FormGroup } from '@angular/forms';
import { KeyValue } from '@angular/common';

@Component({
  selector: 'app-discover',
  templateUrl: './discover.component.html',
  styleUrls: ['./discover.component.css']
})
export class DiscoverComponent implements OnInit, OnChanges {

  friends: boolean = true;
  users: boolean = false;

  user?: userInfo;
  userFriendsIDs: number[] = [7];
  userFriends: Map<number, string> = new Map([[7, 'Friend 1']]);
  userSearches: Map<number, string> = new Map([[8, 'Search 1']]);
  searchForm = new FormGroup({
    search: new FormControl(''),
  });
  friendForm = new FormGroup({
    friendUsername: new FormControl(''),
    friendID: new FormControl(-1),
  });

  constructor(private loginService: LoginService, private userService: UserService, private http: HttpClient) {


  }

  ngOnInit(): void {

    // Grab user's info
    this.user = this.userService.getUserData();

    // Grab friend's IDs
    this.userFriendsIDs = this.loginService.friends;
    console.log(this.userFriendsIDs);

    // For each friend's ID, grab their username
    const fr = async (): Promise<boolean> => {
      return this.getFriendsUsernames();
    };

    fr().then((res) => {

      if (res) {
        console.log('Friends Loaded.');
        console.log(this.userFriends);


      } else {
        console.log('Error: Could not get friends usernames.');
      }

    });




  }

  ngOnChanges(): void {


  }

  search(): void {
    console.log(this.searchForm.value.search);


  }

  setFriendProfile(friend: KeyValue<number, string>) {
    this.friendForm.setValue({
      friendUsername: friend.value,
      friendID: friend.key,
    });
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
