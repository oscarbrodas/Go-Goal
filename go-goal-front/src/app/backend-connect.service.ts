import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';

@Injectable()
export class BackendConnectService {
  backendURL = "/api"

  constructor(private http: HttpClient) {

  }

  public getLoginInfo(): Observable<string> {
    return this.http.get<string>(this.backendURL);
  }

}
