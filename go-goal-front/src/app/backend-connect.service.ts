import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';

@Injectable({ providedIn: 'root' })
export class BackendConnectService {
  backendURL = "http://localhost:9000/"
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };
  constructor(private http: HttpClient) {

  }

  public getLoginInfo(li: loginInfo): Observable<userInfo> {
    return this.http.get<userInfo>(`${this.backendURL}login`) // SOMEHOW SEND THE LOGIN INFO TO THE BACKEND FOR THIS TO WORK
  };

  public signThemUp(userData: userInfo): Observable<userInfo> {
    return this.http.post<userInfo>(`${this.backendURL}users`, userData, this.httpOptions)
  }

}

export interface userInfo { // ADD: User data as necessary 
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