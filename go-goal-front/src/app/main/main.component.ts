import { Component, OnInit } from '@angular/core';


@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css'],
})
export class MainComponent implements OnInit {

  slides1: any[] = new Array(6).fill({ id: -1, src: '', title: '', subtitle: '' });

  constructor() { }

  ngOnInit(): void {
    this.slides1[0] = {
      id: 1,
      src: '../assets/dashboard_images/p.png',
      title: '',
      subtitle: ''
    }

    this.slides1[1] = {
      id: 2,
      src: '../assets/dashboard_images/p.png',
      title: '',
      subtitle: ''
    }

    this.slides1[2] = {
      id: 3,
      src: '../assets/dashboard_images/p.png',
      title: '',
      subtitle: ''
    }

    this.slides1[3] = {
      id: 4,
      src: '../assets/dashboard_images/p.png',
      title: '',
      subtitle: ''
    }

    this.slides1[4] = {
      id: 5,
      src: '../assets/dashboard_images/p.png'
    }

    this.slides1[5] = {
      id: 6,
      src: '../assets/dashboard_images/nate_in_utah.jpg',
      title: 'Goal Setters',
      subtitle: 'We always shoot for the moon and aim no lower. That\'s why at GoGoal, we are committed to    helping you reach your goals.'
    };

  }

}


