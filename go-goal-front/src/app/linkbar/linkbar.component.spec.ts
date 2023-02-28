import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LinkbarComponent } from './linkbar.component';

describe('LinkbarComponent', () => {
  let component: LinkbarComponent;
  let fixture: ComponentFixture<LinkbarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LinkbarComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LinkbarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
