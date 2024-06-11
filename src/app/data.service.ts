import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Item } from './item';

@Injectable({
  providedIn: 'root'
})
export class DataService {

  apiURL = "http://localhost:1000/start"

  constructor(private http: HttpClient) { }

  getHomePosts(): Observable<Item[]> {
    return this.http.get<Item[]>(this.apiURL)
  }

}
