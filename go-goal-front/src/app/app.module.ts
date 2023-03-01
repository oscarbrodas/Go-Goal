import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { MatToolbarModule } from '@angular/material/toolbar';
import { FormsModule } from '@angular/forms'
import { AppRoutingModule } from './app-routing.module';
import { CarouselModule, CardComponent, CardModule } from '@coreui/angular';

import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NavbarTopComponent } from './navbar/navbar-top/navbar-top.component';
import { MainComponent } from './main/main.component';
import { LoginPageComponent } from './login-page/login-page.component';
import { SignUpComponentComponent } from './sign-up-component/sign-up-component.component'
import { LinkbarComponent } from './linkbar/linkbar.component';

import { HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';
import { BackendConnectService } from './backend-connect.service'
import { LoginService } from './login-page/login.service';


@NgModule({
  declarations: [
    AppComponent,
    NavbarTopComponent,
    LoginPageComponent,
    SignUpComponentComponent,
    MainComponent,
    LinkbarComponent
  ],
  imports: [
    BrowserModule,
    ReactiveFormsModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatToolbarModule,
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
