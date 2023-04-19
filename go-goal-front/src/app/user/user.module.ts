import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { UserRoutingModule } from './user-routing.module';
import { UserComponent } from './user/user.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { GoalsComponent } from './goals/goals.component';
import { MatIcon, MatIconModule } from '@angular/material/icon';
import { FormGroup, ReactiveFormsModule } from '@angular/forms';
import { SettingsComponent } from './settings/settings.component';
import { ProfileComponent } from './profile/profile.component';
import { DiscoverComponent } from './discover/discover.component';
import { ImageCropperModule } from 'ngx-image-cropper';


@NgModule({
  declarations: [
    UserComponent,
    GoalsComponent,
    SettingsComponent,
    ProfileComponent,
    DiscoverComponent
  ],
  imports: [
    CommonModule,
    UserRoutingModule,
    MatIconModule,
    ReactiveFormsModule,
    ImageCropperModule,
  ]
})
export class UserModule { }
