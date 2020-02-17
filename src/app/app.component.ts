import { Component, OnInit } from '@angular/core';

import { UserService } from './core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html'
})
export class AppComponent implements OnInit {
  constructor(
    private userService: UserService,
    private router: Router
  ) { }

  ngOnInit() {
    this.userService.populate();
    // Redirect to the log in page since that is the only page currently
    this.router.navigate(['/login']);
  }
}
