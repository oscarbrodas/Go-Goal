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

  public getLoginInfo(li: FormGroup): Observable<any> {
    let loginParams = new HttpParams()
    loginParams = loginParams.append("Email", li.getRawValue().Email);
    loginParams = loginParams.append("Password", li.getRawValue().Password);
    return this.http.get<FormGroup>(`${this.backendURL}login`, { params: loginParams, headers: new HttpHeaders({ 'Content-Type': 'application/json' }) })
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