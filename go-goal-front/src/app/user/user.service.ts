import { Injectable } from '@angular/core';
import { userInfo } from '../backend-connect.service';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private storageName: string = 'user';

  public loggedIn: boolean = false;
  logIn: Promise<boolean> = new Promise((resolve, reject) => { });


  constructor() { }

  setUserData(data: userInfo) {
    localStorage.setItem(this.storageName, JSON.stringify(data));
    this.loggedIn = true;
    this.logIn = Promise.resolve(true);
  }

  getUserData(): userInfo {
    let data = localStorage.getItem(this.storageName);
    return JSON.parse(data!);
  }

  clearUserData() {
    localStorage.removeItem(this.storageName);

  }

  cleanStorage() {
    localStorage.clear();
  }

  isLoggedIn(): boolean {


    let data = localStorage.getItem(this.storageName);
    if (data == null) {
      return false;
    }
    else {
      return JSON.parse(data!).loggedIn;
    }




  }




}
