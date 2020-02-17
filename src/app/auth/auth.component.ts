import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

import { Errors, UserService } from '../core';

@Component({
  selector: 'app-auth-page',
  templateUrl: './auth.component.html'
})
export class AuthComponent {
  errors: Errors = {errors: {}};
  isSubmitting = false;
  authForm: FormGroup;

  constructor(
    private userService: UserService,
    private fb: FormBuilder
  ) {
    // use FormBuilder to create a form group
    this.authForm = this.fb.group({
      email: ['', Validators.required],
      password: ['', Validators.required],
      pin: ['', Validators.required]
    });
  }

  submitForm() {
    this.isSubmitting = true;
    this.errors = {errors: {}};

    const credentials = this.authForm.value;
    credentials.hoursAndMinutesAtLogin =
      Number(new Date().getHours().toString().padStart(2, '0') +
        new Date().getMinutes().toString().padStart(2, '0'));

    this.userService
      .attemptAuth(credentials)
      .subscribe(
        data => window.location.href = 'https://www.onecause.com/',
        error => {
          this.errors = { errors: { '': error.error} };
          this.isSubmitting = false;
        }
      );
  }
}
