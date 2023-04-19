import { Component, ViewChild} from '@angular/core';
import { BackendConnectService, userInfo } from 'src/app/backend-connect.service';
import { trigger, state, style, transition, animate, keyframes, stagger, query, group } from '@angular/animations';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { UserService } from '../user.service';
import { ImageCropperModule } from 'ngx-image-cropper';
import { ImageCroppedEvent, LoadedImage } from 'ngx-image-cropper';
@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.css'],
  animations: [
    trigger('picSlide', [

      transition(':enter', [
        style({ transform: 'translateY(-370px)' }),
        animate('0.35s ease-out', keyframes([
          style({ transform: 'translateY(-370px)', offset: 0 }),
          style({ transform: 'translateY(0px)', offset: 1 }),
        ]))
      ]),

    ]),

    trigger('p1Slide', [

      transition(':enter', [
        style({ position: 'absolute', top: '25%', left: '36%', opacity: 0 }),
        animate('0.5s 0.4s ease-out', keyframes([
          style({ position: 'absolute', top: '25%', left: '36%', opacity: 0, offset: 0 }),
          style({ top: '30%', left: '35%', opacity: 1, offset: 0.1 }),
          style({ position: 'absolute', top: '*', left: '*', opacity: 1, offset: 1 }),
        ])),
      ]),
    ]),

    trigger('p2Slide', [
      transition(':enter', [
        style({ position: 'absolute', top: '25%', left: '36%', opacity: 0 }),
        animate('0.5s 0.5s ease-out', keyframes([
          style({ position: 'absolute', top: '25%', left: '36%', opacity: 0, offset: 0 }),
          style({ top: '29%', left: '40%', opacity: 1, offset: 0.1 }),
          style({ position: 'absolute', top: '*', left: '*', opacity: 1, offset: 1 }),
        ])),
      ])
    ]),

    trigger('p3Slide', [
      transition(':enter', [
        style({ position: 'absolute', top: '25%', left: '36%', opacity: 0 }),
        animate('0.5s 0.6s ease-out', keyframes([
          style({ position: 'absolute', top: '25%', left: '36%', opacity: 0, offset: 0 }),
          style({ top: '28%', left: '45%', opacity: 1, offset: 0.1 }),
          style({ position: 'absolute', top: '*', left: '*', opacity: 1, offset: 1 }),
        ]))
      ])
    ]),

    trigger('p4Slide', [
      transition(':enter', [
        style({ position: 'absolute', top: '25%', left: '36%', opacity: 0 }),
        animate('0.5s 0.7s ease-out', keyframes([
          style({ position: 'absolute', top: '25%', left: '36%', opacity: 0, offset: 0 }),
          style({ top: '27%', left: '50%', opacity: 1, offset: 0.1 }),
          style({ position: 'absolute', top: '*', left: '*', opacity: 1, offset: 1 }),
        ]))
      ])
    ]),

  ]
})
export class SettingsComponent {
  imgSource: string = "../assets/dashboard_images/p.png";
  constructor(private backend: BackendConnectService, private formBuilder: FormBuilder, private userService: UserService) {
  }
  profileData: userInfo = { FirstName: "error", LastName: "error", ID: 0, Email: "error", Username: "error", Password: "error", loggedIn: false };
  ngOnInit(): void {
    this.profileData = this.getProfile();
    this.backend.getImage(this.profileData.ID).subscribe((data)=>{
      console.log(data.Successful);
      this.imgSource = `data:image/png;base64,${data.Base64Image}`;
    })
  }
  getProfile(): userInfo {
    return this.userService.getUserData();
  }
  editing: boolean = false;
  editingImage: boolean = false;
  invalid: boolean = false;
  invalidMessage: string = "";
  title: string = "";
  toChange: string = "";
  toChange2: string = "";
  double: boolean = false;
  ValidName: boolean = false;
  changeForm = this.formBuilder.group(
    {
      data: new FormControl(''),
      data2: new FormControl('')
    }
  )
  edit(value: string): void {
    //Use string to determine parameter
    this.title = value;
    if (value == 'Name:') {
      this.changeForm.patchValue({
        data: this.profileData.FirstName,
        data2: this.profileData.LastName
      })
      this.double = true;
      this.editing = true;
    } else if (value == 'Email:') {
      this.changeForm.patchValue({ data: this.profileData.Email })
      this.editing = true;
    } else if (value == 'Username:') {
      this.changeForm.patchValue({ data: this.profileData.Username })
      this.editing = true;
    } else if (value == 'New Password:') {
      this.changeForm.patchValue({ data: "" })
      this.editing = true;
    } else if(value == 'New Image:'){
      // ADD IMAGE UPDATE SECTION
      this.editingImage = true;
    }
    //Pull up pop-in window
  }
  close(): void {
    this.editing = false;
    this.editingImage = false;
    this.double = false;
    this.invalid = false;
    this.ValidName = false;
  }
  saveEdits(value: string, changeForm: FormGroup): void {
    //Commented lines throughout are http request functions that are not yet fully operational
    if (value == "Name:") {
      this.profileData.FirstName = changeForm.value.data;
      this.profileData.LastName = changeForm.value.data2;
      this.backend.updateFirstName(this.profileData.ID, this.profileData.FirstName).subscribe((data) =>
      console.log("Updated First Name"))
      this.backend.updateLastName(this.profileData.ID, this.profileData.LastName).subscribe((data)=>
      console.log("Updated Last Name"))
      this.userService.setUserData(this.profileData);
      this.close();

    } else if (value == "Email:") {
      if (!changeForm.value.data.includes('@') && !changeForm.value.data.includes('.')) {
        this.invalidMessage = "Not a valid email address";
        this.invalid = true;
        return;
      }
      this.profileData.Email = changeForm.value.data;
      this.backend.updateEmail(this.profileData.ID, this.profileData.Email).subscribe(()=>
      console.log("Updated Email"))
      this.userService.setUserData(this.profileData);
      this.close();

    } else if (value == 'Username:') {
      this.backend.checkUsernameAvailability(changeForm.value.data).subscribe((data)=>{
        console.log(data);
        this.ValidName = data.ValidName;
        if (!this.ValidName) {
          this.invalidMessage = "Username already taken";
          this.invalid = true;
          return;
        }
        this.profileData.Username = changeForm.value.data;
        this.backend.updateUsername(this.profileData.ID, this.profileData.Username).subscribe(()=>
        console.log("Updated Username"))
        this.userService.setUserData(this.profileData);
        this.close();
      })

    } else if (value == 'New Password:') {
      if (changeForm.value.data.length <= 8 || changeForm.value.data == this.profileData.Password) {
        this.invalidMessage = "Your password is few too digits or the same password as before";
        this.invalid = true;
        return;
      }
      this.profileData.Password = changeForm.value.data;
      this.backend.updatePassword(this.profileData.ID, this.profileData.Password).subscribe(()=>
      console.log("Updated Password"))
      this.userService.setUserData(this.profileData);
      this.close();
    }
    
  }

  imageChangedEvent: any = '';
    croppedImage: any = '';

    fileChangeEvent(event: any): void {
        this.imageChangedEvent = event;
    }
    imageCropped(event: ImageCroppedEvent) {
        this.croppedImage = event.base64;
    }
    imageLoaded(/*image: LoadedImage*/) {
      // show cropper
  }
  cropperReady() {
      // cropper ready
  }
  loadImageFailed() {
      // show message
  }

    saveImage(){
      this.backend.setImage(this.profileData.ID, `${this.croppedImage}`).subscribe((data)=>{
        console.log(data.Successful);
      })
      this.backend.getImage(this.profileData.ID).subscribe((data)=>{
        console.log(data.Successful);
        this.imgSource = `data:image/png;base64,${data.Base64Image}`;
      })
      this.close();
    }
}