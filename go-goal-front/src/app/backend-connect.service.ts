import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';

@Injectable({ providedIn: 'root' })
export class BackendConnectService {
  backendURL = "http://localhost:9000/api/"
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };
  constructor(private http: HttpClient) {

  }

  public getLoginInfo(loginInfo: loginInfo): Observable<any> {
    return this.http.get<JSON>(`${this.backendURL}login/${loginInfo.Email}/${loginInfo.Password}`);
  };

  public signThemUp(userData: userInfo): Observable<any> {
    return this.http.post<JSON>(`${this.backendURL}users`, userData, this.httpOptions);
  }

  public getGoals(): Observable<any> {
    return this.http.get<JSON>(`${this.backendURL}goals`, this.httpOptions);
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