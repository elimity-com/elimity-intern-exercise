import {Component, OnInit, ViewChild} from '@angular/core';
import {MatPaginator, MatTableDataSource} from '@angular/material';
import { USERS } from './MOCK_DATA';

export interface User {
  name: string;
  address: string;
  phoneNumber: string;
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent implements OnInit {
  columnsToDisplay = ['name', 'address', 'phone'];
  dataSource = new MatTableDataSource<User>(USERS);

  @ViewChild(MatPaginator) paginator: MatPaginator;

  ngOnInit() {
    this.dataSource.paginator = this.paginator;
  }
}
