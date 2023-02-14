import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';
import { userInfo } from './login-page/login.service';

@Injectable({providedIn: 'root'})
export class BackendConnectService {
  backendURL = "localhost:9000/api/"
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };
  constructor(private http: HttpClient) {

  }

  public getLoginInfo(): Observable<userInfo[]> {
    return this.http.get<userInfo[]>(`${this.backendURL}login`);
  }
  public signThemUp(userData: userInfo): Observable<userInfo>{
    return this.http.post<userInfo>(`${this.backendURL}sign-up`, userData, this.httpOptions)
  }

}