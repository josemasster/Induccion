import { Component, EventEmitter, Input, Output } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { SeriesService } from 'src/app/services/contact/contact.service';

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
  selector: 'app-modal',
  templateUrl: './modal.component.html',
  styleUrls: ['./modal.component.css']
})
export class ModalComponent {
  @Input() option!: string ;
  autoForm!: FormGroup
  dataactualized! : ContactItem[];
  data!: Contact[];

  constructor(private dataSeire: SeriesService){}

  showModal(option: number){
    console.log(option)

    console.log(option);
    if (option == 1){
      this.option= "Crear ";
       
    }
    if (option == 2){
      this.option= "Editar";
     
    }
    const myModal = document.getElementById('myModal')
    if (myModal != null) {
      myModal.style.display = 'block';
    }
  }
  closeModal() {
    const myModal = document.getElementById('myModal')
    if (myModal!= null) {
      myModal.style.display = 'none';
    }
  }

 

  CreateContactData(fullName: string, document: string, direction: string) {
    console.log(fullName +  document + direction)
    console.log('Ejecuta')
      this.dataSeire.CreateAContact(fullName, document, direction).subscribe((response: Contact[])=>{
        console.log(response)
        location.reload();

     }
    )}

    DeleteContactData(Id: number) {
      console.log(Id)
      console.log('Ejecuta')
        this.dataSeire.DeleteAContact(Id).subscribe((response: Contact[])=>{
          console.log(response)
         
  
  
       }
      )}


  ngOnInit(): void {





    this.autoForm = new FormGroup({ 
      nombre: new FormControl(
        {
          value: "",
          disabled: false

        },
        { validators: [Validators.required, Validators.maxLength(50)] }
      ),
      documento: new FormControl( 
        {
          value: "",
          disabled: false

        },
        { validators: [Validators.required, Validators.maxLength(10), Validators.minLength(10)],}
      ),
      direccion: new FormControl(
        {
          value: "",
          disabled: false

        },
        { validators: [Validators.required]}
      ),
      numero: new FormControl(
        {
          value: "",
          disabled: false

        },
        { validators: [Validators.required]}
      ),
      email: new FormControl(
        {
          value: "",
          disabled: false

        },
        { validators: [Validators.required, Validators.email] }
      ),
    })
  }
}




