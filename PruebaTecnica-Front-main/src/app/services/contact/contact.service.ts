import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs/internal/Observable';
import { catchError } from 'rxjs';
import { Contact } from 'src/app/components/agenda/agenda.component';


@Injectable({
  providedIn: 'root'
})
export class SeriesService {


  constructor(private http: HttpClient,) { }


  httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json',
    })
  }

  GetAllContacts(): Observable<any> {
    return this.http.get<any>('http://localhost:8090/v1/contacto', this.httpOptions)
      .pipe(catchError((error: HttpErrorResponse) => {
        console.log(error);
        throw new Error('Error in Get');
      }));
  }

  CreateAContact( NombreCompleto: string, NumeroDocumento: string, Direccion: string,): Observable<any[]> {
    
    const data = {
      NombreCompleto : NombreCompleto,
      NumeroDocumento : NumeroDocumento,
      Direccion :  Direccion,
    }
    return this.http.post<Contact[]>("http://localhost:8080/v1/contactos", data, this.httpOptions)
      .pipe(catchError((error: HttpErrorResponse) => {
        console.log(error);
        throw new Error('Error in Post');
      }));
  }

  UpdateAContact(ID: string, fullName:string, document:string, direction:string, PhoneNumber:string, Email:string): Observable<Contact[]> {
    console.log(ID)
    const url = `http://localhost:3000/contact/${ID}`;
    const data = {
      fullName,
      document,
      direction,
      PhoneNumber,
      Email
    }
    return this.http.put<Contact[]>(url, data, this.httpOptions)
      .pipe(catchError((error: HttpErrorResponse) => {
        console.log(error);
        throw new Error('Error in Put');
      }));
  }

  DeleteAContact(ID: number): Observable<Contact[]> {
    const url = `http://localhost:8080/v1/contactos/${ID}`;
    return this.http.delete<Contact[]>(url, this.httpOptions)
      .pipe(catchError((error: HttpErrorResponse) => {
        console.log(error);
        throw new Error('Error in Delete');
      }));
  }
}
