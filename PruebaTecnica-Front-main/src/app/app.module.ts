import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';
import {MatFormFieldModule} from '@angular/material/form-field';

//Modulos
import { AppRoutingModule } from './app-routing.module';

//Componentes
import { AppComponent } from './app.component';
import { NavbaarComponent } from './components/navbar/navbaar.component';
import { AgendaComponent } from './components/agenda/agenda.component';
import { CreatecontactComponent } from './components/createcontact/createcontact.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ModalComponent } from './components/modal/modal.component';




@NgModule({
  declarations: [
    AppComponent,
    NavbaarComponent,
    AgendaComponent,
    CreatecontactComponent,
    ModalComponent,

  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ReactiveFormsModule,
    BrowserAnimationsModule,
    MatFormFieldModule
    
    ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
