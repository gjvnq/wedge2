import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'wedge²-webui';
  hero: Hero = {
  	id: 2,
  	name: "Olá"
  };
}

export class Hero {
  id: number;
  name: string;
}
