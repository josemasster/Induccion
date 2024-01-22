import { AfterViewInit, Component, ViewChild } from '@angular/core';
import { Form, FormControl, FormGroup, Validators } from '@angular/forms';

import { SeriesService } from 'src/app/services/contact/contact.service';
import { ModalComponent } from '../modal/modal.component';


export interface ContactItem {
  Id :number  ,  
	NombreCompleto :string ,
	NumeroDocumento :string ,
	Direccion: string ,
	FechaCreacion:string ,
}
export interface Contact {
   Data: ContactItem[],
   Message: string, 
   Status: string, 
   Sucess: boolean
}



@Component({
  selector: 'app-agenda',
  templateUrl: './agenda.component.html',
  styleUrls: ['./agenda.component.css']
})
export class AgendaComponent implements AfterViewInit {
  @ViewChild(ModalComponent) funcsModal!: ModalComponent;
  data!: Contact[];

  autoForm!: FormGroup ;
  dataactualized! : ContactItem[];
  



  constructor(private dataSeire: SeriesService) { }
  ngAfterViewInit(): void {
  }
  
  showModal(option: number){
    console.log(option);
    this.funcsModal.showModal(option)
  }

  closeModal() {
    const myModal = document.getElementById('myModal')
    if (myModal!= null) {
      myModal.style.display = 'none';
    }
  }






  DeleteContactData(ID: number) {
    if (confirm('Seguro de eliminar')) {
      this.dataSeire.DeleteAContact(ID).subscribe((response) => {
        console.log(response)
        alert('Contato Eliminado, refresca para visualizar el cambio')
        location.reload();
      })

    }
  }

  ngOnInit(): void {


    this.dataSeire.GetAllContacts().subscribe((response) => {
      console.log(response)
      this.dataactualized = response.Data.DatosContactos;
    });
  }
}
