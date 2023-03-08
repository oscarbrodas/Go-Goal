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

  public getLoginInfo(loginInfo: loginInfo): Observable<loginInfo> {
    return this.http.post<loginInfo>(`${this.backendURL}login`, loginInfo, this.httpOptions);
  };

  public signThemUp(userData: userInfo): Observable<userInfo> {
    return this.http.post<userInfo>(`${this.backendURL}users`, userData, this.httpOptions);
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