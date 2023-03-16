import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatDialogModule } from '@angular/material/dialog';
import { MatButtonModule } from '@angular/material/button';
import { FormsModule } from '@angular/forms'
import { AppRoutingModule } from './app-routing.module';
import { CarouselModule, CardComponent, CardModule } from '@coreui/angular';
import { MatIconModule } from '@angular/material/icon';

import { AppComponent } from './app.component';
import { BrowserAnimationsModule, NoopAnimationsModule } from '@angular/platform-browser/animations';
import { NavbarTopComponent } from './navbar/navbar-top/navbar-top.component';
import { MainComponent } from './main/main.component';
import { LoginPageComponent } from './login-page/login-page.component';
import { SignUpComponentComponent } from './sign-up-component/sign-up-component.component'
import { LinkbarComponent } from './linkbar/linkbar.component';

import { HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';
import { BackendConnectService } from './backend-connect.service'
import { LoginService } from './login-page/login.service';
import { NotFoundComponent } from './not-found/not-found.component';


@NgModule({
  declarations: [
    AppComponent,
    NavbarTopComponent,
    LoginPageComponent,
    SignUpComponentComponent,
    MainComponent,
    LinkbarComponent,
    NotFoundComponent
  ],
  imports: [
    BrowserModule,
    ReactiveFormsModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatToolbarModule,
    MatDialogModule,
    MatButtonModule,
    MatIconModule,
    HttpClientModule,
    FormsModule,
    CarouselModule,
    CardModule
  ],
  exports: [],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
