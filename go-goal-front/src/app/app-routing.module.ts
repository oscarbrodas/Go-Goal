import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { MainComponent } from './main/main.component';
import { LoginPageComponent } from './login-page/login-page.component'
import { SignUpComponentComponent } from './sign-up-component/sign-up-component.component'
import { NotFoundComponent } from './not-found/not-found.component';
import { UserModule } from './user/user.module';

const routes: Routes = [
  { path: 'main', component: MainComponent },
  { path: '', redirectTo: '/main', pathMatch: 'full' },
  { path: 'login', component: LoginPageComponent },
  { path: 'sign-up', component: SignUpComponentComponent },
  { path: 'aboutus', component: NotFoundComponent },
  { path: 'help', component: NotFoundComponent },
  { path: 'user/:id', loadChildren: () => import('./user/user.module').then(m => m.UserModule) },
  //Must leave 404 redirect as last route on page
  { path: '**', component: NotFoundComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
