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
  invalid: boolean = false;
  invalidMessage: string = "";
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
    }else if (value == 'New Password:'){
      this.changeForm.patchValue({data: ""})
    }
    //Pull up pop-in window
    this.editing = true;
    //Close window with submission there and reload
  }
  close(): void {
    this.editing = false;
    this.double = false;
    this.invalid = false;
  }
  saveEdits(value: string, changeForm: FormGroup): void {
    //COmmented lines throughout are http request functions that are not yet fully operational
    if (value == "Name:"){
      this.profileData.FirstName = changeForm.value.data;
      this.profileData.LastName = changeForm.value.data2;
      //this.backend.updateFirstName(this.profileData.ID, this.profileData.FirstName).subscribe((data) =>
      //console.log("Updated First Name"))
      //this.backend.updateLastName(this.profileData.ID, this.profileData.LastName).subscribe((data)=>
      //console.log("Updated Last Name"))
    }else if (value == "Email:"){
      if(!changeForm.value.data.includes('@') && !changeForm.value.data.includes('.')){
        this.invalidMessage = "Not a valid email address";
        this.invalid = true;
        return;
      }
      this.profileData.Email = changeForm.value.data;
      //this.backend.updateEmail(this.profileData.ID, this.profileData.Email).subscribe(()=>
      //console.log("Updated Email"))
    }else if (value == 'Username:'){
      this.profileData.Username = changeForm.value.data;
      //this.backend.updateUsername(this.profileData.ID, this.profileData.Username).subscribe(()=>
      //console.log("Updated Username"))
    }else if (value == 'New Password:'){
      if(changeForm.value.data.length <= 8 || changeForm.value.data == this.profileData.Password){
        this.invalidMessage = "Your password is few too digits or the same password as before";
        this.invalid = true;
        return;
      }
      this.profileData.Password = changeForm.value.data;
      //this.backend.updatePassword(this.profileData.ID, this.profileData.Password).subscribe(()=>
      //console.log("Updated Password"))
    }
    this.userService.setUserData(this.profileData);
    this.close();
  }
}
