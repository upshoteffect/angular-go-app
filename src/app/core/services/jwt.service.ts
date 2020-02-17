import { Injectable } from '@angular/core';

const localStorageProp = 'jwtToken';

@Injectable()
export class JwtService {

  getToken(): string {
    return window.localStorage[localStorageProp];
  }

  saveToken(token: string) {
    window.localStorage[localStorageProp] = token;
  }

  destroyToken() {
    window.localStorage.removeItem(localStorageProp);
  }

}
