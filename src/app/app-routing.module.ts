import { NgModule } from '@angular/core';
import { Routes, RouterModule, PreloadAllModules } from '@angular/router';

// Add more routes here if preloading is used
/*const routes: Routes = [
  {
    path: '',
    loadChildren: ''
  }
];*/

@NgModule({
  imports: [RouterModule.forRoot([], {
    // preload all modules; optionally we could
    // implement a custom preloading strategy for just some
    // of the modules (PRs welcome 😉)
    preloadingStrategy: PreloadAllModules
  })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
