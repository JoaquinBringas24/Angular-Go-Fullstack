import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { NavComponent } from './nav/nav.component';
import { CarrouselComponent } from './carrousel/carrousel.component';
import { ItemComponent } from './item/item.component';
import { Item } from './item';
import { DataService } from './data.service';
@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, NavComponent, CarrouselComponent, ItemComponent, ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  title = 'tutorial';

  items: Item[] = [];

  constructor(private data_service: DataService) {}
  ngOnInit() {
    this.data_service.getHomePosts().subscribe({
      next: (items) => {
        this.items = items;
        console.log(items[0]);
      }
    })
  }
}
