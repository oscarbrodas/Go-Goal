import { Component } from '@angular/core';
import { BackendConnectService, userInfo } from 'src/app/backend-connect.service';
import { trigger, state, style, transition, animate, keyframes, stagger, query } from '@angular/animations';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { UserService } from '../user.service';

@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.css']
})
export class SettingsComponent {
  constructor(private backend: BackendConnectService, private formBuilder: FormBuilder, private userService: UserService) {
  }
  profileData: userInfo = {FirstName: "error", LastName: "error", ID: 0, Email: "error", Username: "error", Password: "error", loggedIn: false};
  ngOnInit(): void{
    this.profileData = this.getProfile();
  }
  getProfile(): userInfo{
    return this.userService.getUserData();
  }
  editing: boolean = false;
  title: string = "";
  toChange: string = "";
  toChange2: string = "";
  double: boolean = false;
  changeForm = this.formBuilder.group(
    {data: new FormControl(''),
    data2: new FormControl('')
  }
  )
  edit(value: string): void{
    //Use string to determine parameter
    this.title = value;
    if (value == 'Name:'){
      this.changeForm.patchValue({data: this.profileData.FirstName,
        data2: this.profileData.LastName})
      this.double = true;
    }else if (value == 'Email:'){
      this.changeForm.patchValue({data: this.profileData.Email})
    }else if (value == 'Username:'){
      this.changeForm.patchValue({data: this.profileData.Username})
    }else if (value == 'Password:'){
      this.changeForm.patchValue({data: this.profileData.Password})
    }
    //Pull up pop-in window
    this.editing = true;
    //Close window with submission there and reload
  }
  close(): void {
    this.editing = false;
    this.double = false;
  }
  saveEdits(value: string, changeForm: FormGroup): void {
    if (value == "Name:"){
      this.profileData.FirstName = changeForm.value.data;
      this.profileData.LastName = changeForm.value.data2;
      this.backend.updateFirstName(this.profileData).subscribe(() =>
      console.log("Updated First Name"))
      this.backend.updateLastName(this.profileData).subscribe(()=>
      console.log("Updated Last Name"))
    }else if (value == "Email:"){
      this.profileData.Email = changeForm.value.data;
      this.backend.updateEmail(this.profileData)
    }else if (value == 'Username:'){
      this.profileData.Username = changeForm.value.data;
      this.backend.updateUsername(this.profileData)
    }else if (value == 'Password:'){
      this.profileData.Password = changeForm.value.data;
      this.backend.updatePassword(this.profileData)
    }
    this.userService.setUserData(this.profileData);
    this.close();
  }
}
