import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';
import { goal } from './user/goals/goals.component';
import { UserService } from './user/user.service';

@Injectable({ providedIn: 'root' })
export class BackendConnectService {
  body: any;
  backendURL = "http://localhost:9000/api/"
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };
  constructor(private http: HttpClient, private userService: UserService) {

  }

  public getLoginInfo(loginInfo: loginInfo): Observable<any> {
    return this.http.get<JSON>(`${this.backendURL}login/${loginInfo.Email}/${loginInfo.Password}`);
  };

  public verifyLogin(loginInfo: loginInfo): Observable<any> {
    return this.http.get<JSON>(`${this.backendURL}login/${loginInfo.Email}/${loginInfo.Password}`);
  };

  public signThemUp(userData: userInfo): Observable<any> {
    return this.http.post<JSON>(`${this.backendURL}users`, { ID: userData.ID, Username: userData.Username, FirstName: userData.FirstName, LastName: userData.LastName, Email: userData.Email, Password: userData.Password }, this.httpOptions);
  }
  public getInfo(userID: Number): Observable<any> {
    return this.http.get<JSON>(`${this.backendURL}users?id=${userID}`);
  }
  public updateFirstName(userID: number, newName: string): Observable<any> {
    return this.http.put(`${this.backendURL}users/${userID}/firstname`, { "new": newName }, this.httpOptions)
  }
  public updateLastName(userID: number, newName: string): Observable<any> {
    return this.http.put(`${this.backendURL}users/${userID}/lastname`, newName, this.httpOptions)
  }
  public updateEmail(userID: number, newName: string): Observable<any> {
    return this.http.put<JSON>(`${this.backendURL}users/${userID}/email`, { newName }, this.httpOptions)
  }
  public updateUsername(userID: number, newName: string): Observable<any> {
    return this.http.put<JSON>(`${this.backendURL}users/${userID}/username`, { newName }, this.httpOptions)
  }
  public updatePassword(userID: number, newName: string): Observable<any> {
    return this.http.put<JSON>(`${this.backendURL}users/${userID}/password`, { newName }, this.httpOptions)
  }
  public createGoal(goalData: any, ID: Number): Observable<any> {
    return this.http.post<JSON>(`${this.backendURL}goals/${ID}`, goalData, this.httpOptions);
  }

  public updateGoal(goalData: any, ID: Number): Observable<any> {
    return this.http.put<JSON>(`${this.backendURL}goals/${ID}`, goalData, this.httpOptions);
  }

  public deleteGoals(gID: Number): Observable<any> {
    console.log(gID);

    return this.http.delete<JSON>(`${this.backendURL}goals/${gID}`, this.httpOptions);
  }

  public getGoals(ID: Number): Observable<any> {
    return this.http.get<JSON>(`${this.backendURL}goals/${ID}`, this.httpOptions);
  }

  public getFriends(ID: Number): Observable<any> {
    return this.http.get<JSON>(`${this.backendURL}friends/${ID}`, this.httpOptions);
  }


}

export interface userInfo { // ADD: User data as necessary 
  ID: number;
  loggedIn: boolean;
  Username: string;
  FirstName: string;
  LastName: string;
  Email: string;
  Password: string;
  XP?: number;
  Description?: string;
}

export interface loginInfo {
  Email: string;
  Password: string;
}