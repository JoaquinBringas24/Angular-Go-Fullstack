import { CommonModule } from '@angular/common';
import { HttpClient, provideHttpClient } from '@angular/common/http';
import { Component, Input, OnInit, inject } from '@angular/core';

@Component({
  selector: 'app-item',
  standalone: true,
  imports: [CommonModule,],
  templateUrl: './item.component.html',
  styleUrl: './item.component.scss',
})
export class ItemComponent { 
    @Input() name: string
    @Input() description: string
    @Input() terms: string
    @Input() sku: string
    @Input() price: string
    @Input() images: string
    @Input() section: string
}
