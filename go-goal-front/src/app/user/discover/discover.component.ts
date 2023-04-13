import { Component, OnChanges, OnInit } from '@angular/core';
import { BackendConnectService, userInfo } from '../../backend-connect.service';
import { LoginService } from 'src/app/login-page/login.service';
import { UserService } from '../user.service';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-discover',
  templateUrl: './discover.component.html',
  styleUrls: ['./discover.component.css']
})
export class DiscoverComponent implements OnInit, OnChanges {

  user?: userInfo;
  userFriendsIDs: Number[] = [];
  userFriendsUsernames: String[] = [];

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
        console.log(this.userFriendsUsernames);


      } else {
        console.log('Error: Could not get friends usernames.');
      }

    });




  }

  ngOnChanges(): void {


  }

  getFriendsUsernames(): boolean {

    let res = true;
    this.userFriendsIDs.map((friend) => {

      this.http.get<any>(`http://localhost:9000/api/users?id=${friend}`).subscribe((data) => {
        if (data.ErrorExist) res = false;
        console.log(data);

        this.userFriendsUsernames.push(data.ThisUser.Username);

      });

    });
    return res;
  }

}
