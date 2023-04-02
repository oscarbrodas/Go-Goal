import { Component, OnInit, OnChanges, SimpleChanges, Input } from '@angular/core';
import { UserService } from '../user/user.service';

@Component({
  selector: 'app-linkbar',
  templateUrl: './linkbar.component.html',
  styleUrls: ['./linkbar.component.css']
})
export class LinkbarComponent implements OnInit, OnChanges {

  @Input() loggedIn: boolean = false;

  constructor(private userService: UserService) { }

  ngOnInit(): void {
    this.loggedIn = this.userService.isLoggedIn();

  }

  ngOnChanges(changes: SimpleChanges): void {
  }

}
