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

  public signThemUp(userData: userInfo): Observable<any> {
    return this.http.post<JSON>(`${this.backendURL}users`, userData, this.httpOptions);
  }
  public updateFirstName(userData: userInfo): Observable<any>{
    return this.http.put<JSON>(`${this.backendURL}users/${userData.ID}/firstname`, {ID: userData.ID, FirstName: userData.FirstName}, this.httpOptions)
  }
  public updateLastName(userData: userInfo): Observable<any>{
    return this.http.put<JSON>(`${this.backendURL}users/${userData.ID}/lastname`, {ID: userData.ID, LastName: userData.LastName}, this.httpOptions)
  }
  public updateEmail(userData: userInfo): Observable<any>{
    return this.http.put<JSON>(`${this.backendURL}users/${userData.ID}/email`, userData, this.httpOptions)
  }
  public updateUsername(userData: userInfo): Observable<any>{
    return this.http.put<JSON>(`${this.backendURL}users/${userData.ID}/username`, userData, this.httpOptions)
  }
  public updatePassword(userData: userInfo): Observable<any>{
    return this.http.put<JSON>(`${this.backendURL}users/${userData.ID}/password`, userData, this.httpOptions)
  }
  public createGoal(goalData: any, ID: Number): Observable<any> {
    return this.http.post<JSON>(`${this.backendURL}goals/${ID}`, goalData, this.httpOptions);
  }

  public deleteGoals(gID: Number): Observable<any> {
    console.log(gID);

    return this.http.delete<JSON>(`${this.backendURL}goals/${gID}`, this.httpOptions);
  }

  public getGoals(ID: number): Observable<any> {
    return this.http.get<JSON>(`${this.backendURL}goals/${ID}`, this.httpOptions);
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
}

export interface loginInfo {
  Email: string;
  Password: string;
}